package comptypes

func GetComptypeRules(name string) *TypeGroup {
	if ctg, ok := comptypeLookup[name]; ok {
		return &ctg
	}
	return nil
}

func GetComptypeNames() []string {
	names := make([]string, len(data.Comptypes))
	for i, ct := range data.Comptypes {
		names[i] = ct.Name
	}
	return names
}
