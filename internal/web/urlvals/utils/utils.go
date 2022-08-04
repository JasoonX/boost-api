package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

func GetStructTag(f reflect.StructField, tagName string) string {
	return f.Tag.Get(tagName)
}

func ToString(value interface{}) string {
	if v, ok := value.(reflect.Value); ok {
		value = v.Interface()
	}

	var strValue []string
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		s := reflect.ValueOf(value)

		for i := 0; i < s.Len(); i++ {
			if s.Index(i).CanInterface() {
				strValue = append(strValue, cast.ToString(s.Index(i).Interface()))
			}
		}

		value = strValue
	}

	if v, ok := value.([]string); ok {
		value = strings.Join(v, ",")
	}

	// some magic to convert values of custom aliased types to their,
	// underlying type, because cast fails to do this:
	if value = convertToUnderlyingType(value); value == nil {
		return ""
	}

	return cast.ToString(value)
}

func convertToUnderlyingType(value interface{}) interface{} {
	var rType reflect.Type
	var rValue = reflect.ValueOf(value)
	var rKind = rValue.Kind()

	if rKind == reflect.Pointer {
		if !rValue.Elem().IsValid() {
			return nil
		}

		rValue = rValue.Elem()
		rKind = rValue.Kind()
	}

	switch rKind {
	case reflect.String:
		rType = reflect.TypeOf("")
	case reflect.Bool:
		rType = reflect.TypeOf(false)
	case reflect.Int:
		rType = reflect.TypeOf(0)
	case reflect.Int8:
		rType = reflect.TypeOf(int8(0))
	case reflect.Int16:
		rType = reflect.TypeOf(int16(0))
	case reflect.Int32:
		rType = reflect.TypeOf(int32(0))
	case reflect.Int64:
		rType = reflect.TypeOf(int64(0))
	case reflect.Uint:
		rType = reflect.TypeOf(uint(0))
	case reflect.Uint8:
		rType = reflect.TypeOf(uint8(0))
	case reflect.Uint16:
		rType = reflect.TypeOf(uint16(0))
	case reflect.Uint32:
		rType = reflect.TypeOf(uint32(0))
	case reflect.Uint64:
		rType = reflect.TypeOf(uint64(0))
	case reflect.Float32:
		rType = reflect.TypeOf(uint32(0))
	case reflect.Float64:
		rType = reflect.TypeOf(uint64(0))
	case reflect.Slice:
		rType = rValue.Type()
	default:
		panic(fmt.Sprintf("unknown rKind: %v", rKind))
	}

	return rValue.Convert(rType).Interface()
}

func GetKey(tagName, tagKey string) string {
	return strings.TrimRight(
		strings.TrimLeft(
			strings.TrimPrefix(tagName, tagKey),
			"["),
		"]")
}

func UnwrapTag(tag string) (string, []string) {
	keys := strings.Split(tag, ",")
	return keys[0], keys[1:]
}
