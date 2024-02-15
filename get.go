package comptypes

func GetComptypeRules(name string) *TypeGroup {
	if !initalised {
		return nil
	}

	if ctg, ok := (comptypeLookup)[name]; ok {
		return &ctg
	}
	return nil
}

func GetComptypeNames() []string {
	var names []string
	if !initalised {
		return names
	}

	names = make([]string, len(data.Comptypes))
	for i, ct := range data.Comptypes {
		names[i] = ct.Name
	}
	return names
}
