package comptypes

import (
	"regexp"
	"strings"
)

// GetFieldNamesFromArg will return the columns that match the provided arg, or nil on failure
func GetMatchedPartsFromArgs(ct TypeGroup, arg string, matchedFields []interface{}) (string, []interface{}) {
	var (
		argRemaining = arg
	)

	if !initalised {
		return argRemaining, []interface{}{}
	}

	for _, cond := range ct.GroupsRulesKeys {
		var matchedFieldsRet []interface{}
		if cond.Group != nil {
			argRemaining, matchedFieldsRet = GetMatchedPartsFromArgs(*cond.Group, argRemaining, matchedFieldsRet)
		} else if cond.Rule != nil {
			rxString := cond.Rule.Rx
			if !strings.HasPrefix(rxString, "^") {
				rxString = "^" + rxString
			}

			rx, err := regexp.Compile(rxString)
			if err != nil {
				return argRemaining, matchedFields
			}

			matches := rx.FindStringSubmatch(argRemaining)
			if matches == nil {
				matchedFieldsRet = []interface{}{nil}
			} else {
				argRemaining = strings.TrimPrefix(argRemaining, matches[0])
				matchedFieldsRet = []interface{}{matches[0]}
			}
		} else {
			if cond.TypeKey == nil {
				return argRemaining, matchedFields
			}

			ct := GetComptypeRules(*cond.TypeKey)
			if ct == nil {
				return argRemaining, matchedFields
			}

			argRemaining, matchedFieldsRet = GetMatchedPartsFromArgs(*ct, argRemaining, matchedFieldsRet)
		}

		if len(matchedFieldsRet) > 0 {
			matchedFields = append(matchedFields, matchedFieldsRet...)
		}

		if ct.Op == And && len(argRemaining) == len(arg) {
			break
		}
	}

	return argRemaining, matchedFields
}
