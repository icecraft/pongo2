package pongo2

import (
	"fmt"
	"reflect"
	"strings"
)

type NotListError struct {
}

func (v *NotListError) Error() string {
	return "Not List! Please check it when use pongo2 template\n"
}

func filterExpandListInt(in *Value, param *Value) (*Value, *Error) {

	str := make([]string, 0)

	if in.getResolvedValue().Kind() == reflect.Slice {
		items := in.getResolvedValue()
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			str = append(str, fmt.Sprintf("%v", item))
		}
	} else {
		err := &NotListError{}
		return nil, err

	return AsValue(strings.Join(str, ",")), nil
}
