// rest2sql project main.go
package main

import (
	"fmt"
	"os"
	"rest2sql/sqlWhere"
)

func main() {
	argsList := os.Args[1:]
	var rawFilterString string

	if len(argsList) == 0 {
		rawFilterString = "$filter=(name EQ 'Milk' or name eq 'Eggs') and price Lt 2.55"
	} else {
		var separator string
		for _, field := range argsList {
			rawFilterString = rawFilterString + separator + field
			separator = " "
		}
	}

	fmt.Println("raw $filter string:", rawFilterString)
	whereSentence, err := sqlWhere.MakeWhereSentence(rawFilterString)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("WHERE sentence:", whereSentence)
	}
}
