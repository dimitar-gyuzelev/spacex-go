package v4

import (
	"fmt"
	"reflect"
	"strings"
)

func stringFn[T any](t T) string {
	var sb strings.Builder

	fields := reflect.ValueOf(&t).Elem()
	typeOfT := fields.Type()

	sb.WriteString("{\n")
	for i := 0; i < fields.NumField(); i++ {
		sb.WriteString(fmt.Sprintf("  %s:%+v\n", typeOfT.Field(i).Name, fields.Field(i).Interface()))
	}
	sb.WriteString("}\n")

	return sb.String()
}
