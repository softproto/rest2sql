package main

import (
	"strings"
)

var filter map[string]string = map[string]string{
	"(":   " (",
	")":   ") ",
	"eq":  " = ",
	"ne":  " <> ",
	"gt":  " > ",
	"ge":  " >= ",
	"lt":  " < ",
	"le":  " <= ",
	"and": " AND ",
	"or":  " OR ",
	"not": " NOT ",
}

//create the WHERE sentence from the REST "$filter" string
func MakeWhereSentence(filterString string) string {
	if strings.Contains(filterString, "$filter=") != true {
		return ""
	} else {
		var operatorWhereSentence string = "WHERE "

		filterString = strings.TrimPrefix(filterString, "$filter=")

		filterString = strings.Replace(filterString, "(", " ( ", -1)
		filterString = strings.Replace(filterString, ")", " ) ", -1)

		commandList := strings.Fields(filterString)

		for i := range commandList {
			operator, operatorExist := filter[strings.ToLower(commandList[i])]
			if operatorExist {
				operatorWhereSentence = operatorWhereSentence + operator
			} else {
				operatorWhereSentence = operatorWhereSentence + commandList[i]
			}
		}

		return operatorWhereSentence
	}
}
