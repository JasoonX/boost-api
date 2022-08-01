package urlvals

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

func Encode(src interface{}) string {
	values := url.Values{}

	for _, tagName := range supportedTagsList {
		populateValues(values, tagName, reflect.ValueOf(src))
	}

	return values.Encode()
}

func populateValues(values url.Values, tagName string, v reflect.Value) {
	if !v.IsValid() {
		return
	}

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsValid() {
			continue
		}

		f := v.Type().Field(i)
		if !f.IsExported() {
			continue
		}
		// recurring in order to get embedded structs
		populateValues(values, tagName, v.Field(i))

		tag := getStructTag(f, tagName)
		if tag == "" {
			continue
		}

		value := v.Field(i).Interface()

		tagKey := fmt.Sprintf("%s[%s]", tagName, tag)
		setUrlValues(values, tagKey, value)
	}
}

func getStructTag(f reflect.StructField, tagName string) string {
	return f.Tag.Get(tagName)
}

func setUrlValues(values url.Values, key string, value interface{}) {
	if strVal := toString(value); strVal != "" {
		values.Set(key, strVal)
	}
}

func toString(value interface{}) string {
	if v, ok := value.(reflect.Value); ok {
		value = v.Interface()
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

	if rKind == reflect.Ptr {
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
