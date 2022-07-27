package debug

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const (
	FieldPage = "page"
	FieldID   = "id"
)

func Struct(i interface{}) string {
	typ := reflect.ValueOf(i)

	if typ.Kind() != reflect.Struct {
		typ = typ.Elem()
	}
	res := strings.Builder{}

	for i := 0; i < typ.NumField(); i++ {
		value := typ.Field(i).Interface()
		if typ.Field(i).Kind() == reflect.Pointer && !typ.Field(i).IsNil() {
			value = typ.Field(i).Elem().Interface()
		}
		if v, ok := value.(json.RawMessage); ok {
			if json.Valid(v) {
				value = string(v)
			}
		}
		res.WriteString(fmt.Sprintf("%s => %v; ", typ.Type().Field(i).Name, value))
	}

	return res.String()
}
