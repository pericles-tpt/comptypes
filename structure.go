package comptypes

type Comptype struct {
	Name  string    `json:"name"`
	Rules TypeGroup `json:"rules"`
}

type TypeGroup struct {
	Op              TypeOp         `json:"op"`
	GroupsRulesKeys []GroupRuleKey `json:"groupsRulesKeys"`
}

// NOTE: This is only valid if TWO are nil
type GroupRuleKey struct {
	Group   *TypeGroup `json:"group"`
	Rule    *TypeRule  `json:"rule"`
	TypeKey *string    `json:"typeKey"`
}

type TypeOp int

const (
	And TypeOp = iota
	Or
)

// ComptypeCategory is an enum used by `ComptypeRule`
type ComptypeCategory int

const (
	Alias ComptypeCategory = iota
	Parse
	Enum
	Unit
)

// ComptypeRule lets user specify rules to evaluate high level "types", a single
// user argument can have multiple rules, which are evaluated left to right, each
// rule corresponds to a column in a table. Rules include a regex to match the type
// on and a category that dictates how to evaluate the type. Categories include:
//   - Alias: simply identifies the matched string with an alias name
//   - Parse: for conversion of the matched string to another type, either
//     directly or through a function (using a function key). A value
//     that's parsed to a numeric type (float64 or int64) can also have
//     a `UnitConversion` defined for it (e.g. g -> lbs)
//   - Enum: are treated as strings by go, but are stored as enum types in the
//     db and identified as enums by the frontend JS
type TypeRule struct {
	Name       string           `json:"name"`
	Rx         string           `json:"regex"`
	Category   ComptypeCategory `json:"category"`
	PropsParse *PropsParse      `json:"propsParse"`
	PropsEnum  *PropsEnum       `json:"propsEnum"`
	PropsUnit  *PropsUnit       `json:"propsUnit"`
}

type PropsParse struct {
	FuncKey    string      `json:"funcKey"`
	OutType    string      `json:"outType"`
	Conversion *Conversion `json:"conversion"`
}

type Conversion struct {
	BaseUnit      string `json:"baseUnit"`
	ThisUnitIndex int    `json:"thisUnitIndex"`
}

type PropsEnum struct {
	CaseSensitive bool     `json:"caseSensitive"`
	ValueSet      []string `json:"valueSet"`
}

type PropsUnit struct {
	BaseUnit   string                    `json:"baseUnit"`
	OtherUnits map[string]UnitConversion `json:"otherUnits"`
}

type UnitConversion struct {
	Rx               string `json:"rx"`
	ConversionFactor string `json:"conversionFactor"`
}
