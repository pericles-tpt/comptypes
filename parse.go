package comptypes

import (
	"fmt"
	"strconv"
	"time"
)

const (
	iso8601 = "2006-01-02T15:04:05.000Z"
)

// TODO: Do we want users to specify type conversion functions in the future?
var (
	funcMap = map[string]func(inp string) interface{}{
		"float": func(inp string) interface{} {
			ret, err := strconv.ParseFloat(inp, 64)
			if err != nil {
				return ret
			}
			return nil
		},
		"int": func(inp string) interface{} {
			ret, err := strconv.ParseInt(inp, 10, 64)
			if err != nil {
				return ret
			}
			return nil
		},
		"time": func(inp string) interface{} {
			ret, err := time.Parse(iso8601, inp)
			if err != nil {
				return ret
			}
			return nil
		},
	}
)

func ExecuteParseFunction(funcKey, value string) (interface{}, error) {
	if fn, ok := funcMap[funcKey]; ok {
		return fn(value), nil
	}
	return nil, fmt.Errorf("failed to find function matching key: %s", funcKey)
}
