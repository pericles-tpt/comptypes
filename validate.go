package comptypes

import (
	"errors"
	"fmt"

	"github.com/pericles-tpt/rterror"
)

func (g *TypeGroup) Valid(rootRuleKey string) error {
	var err error

	if !initalised {
		return errors.New("error `LoadCompValidtypes` not called before `Valid`")
	}

	// TODO:
	// Some conditions below return vague errors that don't indicate what point in the recursive structure the error occurred
	// SOLUTION: Generate unique ids for each `GroupRuleKey` for error logging (not necessary for conditions 2 & 4 below)
	for _, gr := range g.GroupsRulesKeys {
		var (
			onlyGroup   = gr.Group != nil && gr.Rule == nil && gr.TypeKey == nil
			onlyRule    = gr.Group == nil && gr.Rule != nil && gr.TypeKey == nil
			onlyTypeKey = gr.Group == nil && gr.Rule == nil && gr.TypeKey != nil
		)

		if !(onlyGroup || onlyRule || onlyTypeKey) {
			return rterror.PrependErrorWithRuntimeInfo(nil, "in the set of `GroupRuleKey` properties (`Group`, `Rule`, `TypeKey`), you MUST specify ONE non-nil property")
		} else if onlyRule {
			err = gr.Rule.Valid()
			if err != nil {
				return rterror.PrependErrorWithRuntimeInfo(err, "non-nil `Rule` '%s' invalid", gr.Rule.Name)
			}
		} else if onlyGroup {
			err = gr.Group.Valid(rootRuleKey)
			if err != nil {
				return rterror.PrependErrorWithRuntimeInfo(err, "non-nil `Group` invalid")
			}
		} else if onlyTypeKey {
			// To avoid infinite recursion
			// Basic: A CustomRule CAN have a `TypeKey` property iff:
			//	1. The other `TypeKey` != `rootRuleKey`
			//  2. The CustomRule exists
			//	3. The CustomRule at that `TypeKey`, DOESN'T contain another `TypeKey`
			// TODO: Advanced (enabled in setting): No limits on nesting of `TypeKey`s in CustomRule(s), except:
			// 	Should be advanced, since allowing nesting of `TypeKey`, potentially increases complexity for users
			//  having to troubleshoot loops in their rules
			//  1. The CustomRule exists
			// 	2. The path taken by a `TypeKey` cannot contain a loop
			//		i.e. refer to the same `TypeKey` more than once, starting from the `rootRuleKey`
			// TODO: Could cache whether `TypeKey` contains another `TypeKey` in global state, saves work
			// TODO: As above, except for caching whether a `CustomRule` contains a loop through its nested `TypeKey`s
			if *(gr.TypeKey) == rootRuleKey {
				return fmt.Errorf("non-nil `TypeKey` invalid, cannot match the root node's key: '%s'", rootRuleKey)
			}

			if !ComptypeExists(*gr.TypeKey) {
				return fmt.Errorf("non-nil `TypeKey` invalid, cannot find existing rule matching key: '%s'", *(gr.TypeKey))
			}

			if ComptypeContainsATypeKey(*gr.TypeKey) {
				return fmt.Errorf("non-nil `TypeKey` invalid, the CustomRule contains another `TypeKey` (not allowed in BASIC mode): '%s'", *(gr.TypeKey))
			}
		}
	}
	return nil
}

func (r *TypeRule) Valid() error {
	if !initalised {
		return errors.New("error `LoadCompValidtypes` not called before `Valid`")
	}

	if r.Category == Alias && (r.PropsParse != nil || r.PropsEnum != nil) {
		return rterror.PrependErrorWithRuntimeInfo(nil, "'Alias' category specified but one of `PropsParse` or `PropsEnum` is not nil")
	} else if r.PropsParse == nil && r.PropsEnum == nil {
		return rterror.PrependErrorWithRuntimeInfo(nil, "'Parse' or 'Enum' category specified, BUT both `PropsParse` and `PropsEnum` are nil")
	} else if r.PropsParse != nil && r.PropsEnum != nil {
		return rterror.PrependErrorWithRuntimeInfo(nil, "'Parse' or 'Enum' category specified, BUT both `PropsParse` and `PropsEnum` are non-nil")
	}
	return nil
}

func ComptypeContainsATypeKey(typeKey string) bool {
	if !initalised {
		return false
	}

	if rule, ok := comptypeLookup[typeKey]; ok {
		for _, v := range rule.GroupsRulesKeys {
			if v.TypeKey != nil {
				return true
			}
		}
	}
	return false
}

func ComptypeExists(target string) bool {
	if !initalised {
		return false
	}

	_, ok := comptypeLookup[target]
	return ok
}
