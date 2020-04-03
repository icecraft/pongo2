package pongo2

import (
	"strings"
)

func filterExpandListInt(in *Value, param *Value) (*Value, *Error) {

	listArgs, err := in.([]int)
	listStr := make([]string, 0)

	if err != nil {
		return nil, err
	}
	for i, _ := range listArgs {
		listStr = append(listStr, listArgs[i])
	}
	return AsValue(strings.Join(listStr, ", ")), nil
}
