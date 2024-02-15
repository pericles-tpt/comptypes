package comptypes

import (
	"errors"

	rterror "github.com/pericles-tpt/rterror"
)

/*
Add a user defined `TypeLabel` in `data`
*/
func AddComptype(name string, rules TypeGroup) error {
	if !init {
		return errors.New("error `LoadComptypes` not called before `AddComptype`")
	}

	data.Comptypes = append(data.Comptypes, Comptype{
		Name:  name,
		Rules: rules,
	})

	err := flushUserDataToDisk()
	if err != nil {
		return rterror.PrependErrorWithRuntimeInfo(err, "failed to flush user data")
	}

	return nil
}
