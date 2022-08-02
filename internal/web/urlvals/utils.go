package urlvals

import (
	"reflect"
	"strings"
)

func unwrapTag(tag string) (string, []string) {
	keys := strings.Split(tag, ",")
	return keys[0], keys[1:]
}

func getStructTag(f reflect.StructField, tagName string) string {
	return f.Tag.Get(tagName)
}
