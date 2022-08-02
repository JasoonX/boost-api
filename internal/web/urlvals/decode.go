package urlvals

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// TODO: add sorting, includes to the params

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

	normalizedUrlQuery := normalizeUrlQuery(values)
	normalizedTags := normalizeStructTags(t)

	// merge order matters, since url data has higher priority
	normalizedUrlQuery.Merge(normalizedTags)

	for tagName, tagKeys := range normalizedUrlQuery {
		for tagKey, tagParam := range tagKeys {
			if err := processTag(el, tagName, tagKey, tagParam.Values); err != nil {
				return errors.Wrapf(err, "failed to parse %s tag", tagName)
			}
		}
	}
	return nil
}

func processTag(dest reflect.Value, tagName, tagKey string, values []string) error {
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
			if err := processTag(dest.Field(i), tagName, tagKey, values); err != nil {
				return errors.Wrap(err, "failed to parse tag")
			}

			// tag may also contain modifiers, like "default", thus we need to drop them
			tag, _ := unwrapTag(getStructTag(dest.Type().Field(i), tagName))

			if tag != tagKey {
				continue
			}

			if err := parseIntoValue(dest.Field(i), values...); err != nil {
				return errors.Wrap(err, "failed to set field")
			}
		}
	}

	return nil
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

	if kind == reflect.Pointer {
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
	// to assign type to their aliases
	if reflect.TypeOf(v) != dest.Type() && reflect.TypeOf(v).Kind() == dest.Type().Kind() {
		dest.Set(reflect.ValueOf(v).Convert(dest.Type()))
		return
	}

	dest.Set(reflect.ValueOf(v))
}
