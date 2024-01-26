package comptypes

import (
	"encoding/json"
	"os"

	rterror "github.com/pericles-tpt/rterror"
)

const filePath = "comptypes.json"

type Data struct {
	Comptypes []Comptype `json:"comptypes"`
}

var (
	data           Data
	comptypeLookup = map[string]TypeGroup{}
)

func LoadComptypes() error {
	// Load `userData`
	tmp, err := os.Stat(filePath)
	if os.IsNotExist(err) || tmp.Size() == 0 {
		f, err := os.Create(filePath)
		if err != nil {
			return rterror.PrependErrorWithRuntimeInfo(err, "failed to create `%s` for storing user data", filePath)
		}

		// Here we can just encode the default initialised `UserData` (see above) into the new file, then return
		je := json.NewEncoder(f)
		err = je.Encode(data)
		if err != nil {
			return rterror.PrependErrorWithRuntimeInfo(err, "failed to write default, initialised `Data` to (new) `%s`", filePath)
		}
		return nil
	} else if err != nil {
		return rterror.PrependErrorWithRuntimeInfo(err, "failed to stat `%s`", filePath)
	}

	f, err := os.Open(filePath)
	if err != nil {
		return rterror.PrependErrorWithRuntimeInfo(err, "failed to open `%s` to read existing user data", filePath)
	}
	defer f.Close()

	jd := json.NewDecoder(f)
	err = jd.Decode(&data)
	if err != nil {
		return rterror.PrependErrorWithRuntimeInfo(err, "failed to decode `%s` as `Data`", filePath)
	}

	populateGlobals()

	return nil
}

func flushUserDataToDisk() error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		return rterror.PrependErrorWithRuntimeInfo(err, "error failed to open `%s` to write updated `UserData`", filePath)
	}
	defer f.Close()

	je := json.NewEncoder(f)
	err = je.Encode(data)
	if err != nil {
		return rterror.PrependErrorWithRuntimeInfo(err, "failed to encode updated `UserData` structure to `%s`", filePath)
	}

	return nil
}

func populateGlobals() {
	for _, ct := range data.Comptypes {
		comptypeLookup[ct.Name] = ct.Rules
	}
}
