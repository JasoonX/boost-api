package urlvals

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// TODO: refactor this module and add sorting, includes to the params

var ErrSingleValueExpected = errors.New("expected single value, got slice")

func Decode(values url.Values, dest interface{}) error {
	t := reflect.ValueOf(dest)
	if t.Kind() != reflect.Pointer {
		return errors.New(fmt.Sprintf("dest type must be pointer, got: %v", t.Kind()))
	}

	el := t.Elem()
	if el.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("dest type must be pointer to a struct, got: %v", el.Kind()))
	}

	for _, tagName := range supportedTagsList {
		for k, v := range values {
			if err := parseTag(el, tagName, k, v); err != nil {
				return errors.Wrapf(err, "failed to parse %s tag", tagName)
			}
		}
	}
	return nil
}

func parseTag(dest reflect.Value, tagName, tagKey string, values []string) error {
	if !dest.IsValid() {
		return nil
	}

	if dest.Kind() == reflect.Pointer {
		dest = dest.Elem()
	}

	if dest.Kind() == reflect.Struct {
		// recurring to struct fields for embedded structs
		for i := 0; i < dest.NumField(); i++ {
			if !dest.Type().Field(i).IsExported() {
				continue
			}
			if err := parseTag(dest.Field(i), tagName, tagKey, values); err != nil {
				return errors.Wrap(err, "failed to parse tag")
			}
		}

		// trying to find field in our dest type to set the value
		fieldName, modifiers := findField(dest.Type(), tagName, getKey(tagName, tagKey))
		if fieldName == "" {
			return nil
		}

		f := dest.FieldByName(fieldName)
		if err := setField(f, values, modifiers); err != nil {
			return errors.Wrap(err, "failed to set field")
		}
	}

	return nil
}

func findField(t reflect.Type, tagName, key string) (string, map[string]string) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if !field.IsExported() {
			continue
		}

		// Get the field tag value
		tag := field.Tag.Get(tagName)

		modifiers := parseModifiers(tag)

		if tag == key {
			return field.Name, modifiers
		}
	}
	return "", nil
}

func parseModifiers(tag string) map[string]string {
	modifiers := make(map[string]string)
	for _, modifier := range modifiers {
		if strings.Contains(tag, modifier) {
			if strings.Contains(tag, "=") {
				keyValue := strings.Split(tag, "=")
				modifiers[keyValue[0]] = keyValue[1]
			}
		}
	}
	return modifiers
}

func setField(dest reflect.Value, values []string, modifiers map[string]string) error {
	return parseIntoValue(dest, values...)
}

// parseSlice takes a slice of strings and converts it's items to the "destElemType".
// Returns error, if types are not convertible to each other.
func parseSlice(destElemType reflect.Type, sourceSlice []string) (interface{}, error) {
	destSliceType := reflect.SliceOf(destElemType)
	destSlice := reflect.MakeSlice(destSliceType, len(sourceSlice), len(sourceSlice))

	if len(sourceSlice) == 0 {
		return destSlice.Interface(), nil
	}

	for j := 0; j < len(sourceSlice); j++ {
		err := parseIntoValue(destSlice.Index(j), sourceSlice[j])
		if err != nil {
			return nil, err
		}
	}

	return destSlice.Interface(), nil
}

func parseIntoValue(dest reflect.Value, input ...string) error {
	if len(input) == 0 || !dest.CanSet() {
		return nil
	}

	t := dest.Type()
	kind := t.Kind()

	if kind == reflect.Ptr {
		v := dest
		if dest.IsNil() {
			v = reflect.New(t.Elem())
		}

		err := parseIntoValue(v.Elem(), input...)
		if err != nil {
			return err
		}
		dest.Set(v)
	}

	if len(input) > 1 && dest.Kind() != reflect.Slice {
		return ErrSingleValueExpected
	}
	var convertedValue interface{}
	var err error
	switch kind {
	case reflect.String:
		convertedValue = input[0]
	case reflect.Bool:
		convertedValue, err = cast.ToBoolE(input[0])
	case reflect.Int:
		convertedValue, err = cast.ToIntE(input[0])
	case reflect.Int8:
		convertedValue, err = cast.ToInt8E(input[0])
	case reflect.Int16:
		convertedValue, err = cast.ToInt16E(input[0])
	case reflect.Int32:
		convertedValue, err = cast.ToInt32E(input[0])
	case reflect.Int64:
		convertedValue, err = cast.ToInt64E(input[0])
	case reflect.Uint:
		convertedValue, err = cast.ToUintE(input[0])
	case reflect.Uint8:
		convertedValue, err = cast.ToUint8E(input[0])
	case reflect.Uint16:
		convertedValue, err = cast.ToUint16E(input[0])
	case reflect.Uint32:
		convertedValue, err = cast.ToUint32E(input[0])
	case reflect.Uint64:
		convertedValue, err = cast.ToUint64E(input[0])
	case reflect.Float32:
		convertedValue, err = cast.ToFloat32E(input[0])
	case reflect.Float64:
		convertedValue, err = cast.ToFloat64E(input[0])
	case reflect.Slice:
		convertedValue, err = parseSlice(dest.Type().Elem(), input)
	case reflect.Complex64, reflect.Complex128, reflect.Chan,
		reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr,
		reflect.Array, reflect.Struct, reflect.UnsafePointer:
	default:
		panic(fmt.Sprintf("unknown field kind: %v", kind))
	}
	if err != nil {
		return errors.Wrapf(err, "expected value to be %s, got %s", t, input)
	}

	if convertedValue == nil {
		return nil
	}
	setValue(dest, convertedValue)
	return nil
}

func setValue(dest reflect.Value, v interface{}) {
	// to assign type to their aliases:
	if reflect.TypeOf(v) != dest.Type() && reflect.TypeOf(v).Kind() == dest.Type().Kind() {
		dest.Set(reflect.ValueOf(v).Convert(dest.Type()))
		return
	}

	dest.Set(reflect.ValueOf(v))
}

func getKey(tagName, tagKey string) string {
	return strings.TrimRight(
		strings.TrimLeft(
			strings.TrimPrefix(tagKey, tagName),
			"["),
		"]")
}
