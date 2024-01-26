package main

import (
	"fmt"

	"github.com/pericles-tpt/comptypes/comptypes"
)

func main() {
	err := comptypes.LoadComptypes()
	if err != nil {
		panic(err)
	}

	var names = comptypes.GetComptypeNames()
	fmt.Println("== Comptypes ==")
	for _, n := range names {
		fmt.Printf("%s: %v\n", n, comptypes.GetComptypeRules(n))
	}
}
