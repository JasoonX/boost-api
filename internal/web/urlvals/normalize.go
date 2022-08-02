package urlvals

import (
	"net/url"
	"reflect"
	"strings"

	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/provider"
)

func normalizeUrlQuery(urlQuery url.Values) provider.TagProvider {
	normalizedQueries := provider.New()
	for k, v := range urlQuery {
		for _, tagName := range supportedTagsList {
			if strings.HasPrefix(k, tagName) {
				tagKey := getKey(k, tagName)
				normalizedQueries.Set(tagName, tagKey, v)
			}
		}
	}
	return normalizedQueries
}

func normalizeStructTags(dest reflect.Value) provider.TagProvider {
	normalizedQueries := provider.New()
	normalizeStructTagsRec(dest, normalizedQueries)
	return normalizedQueries
}

// this is a recursive function on dfs (first setting the deepest fields)
func normalizeStructTagsRec(dest reflect.Value, normalizedProvider provider.TagProvider) {
	if dest.Kind() == reflect.Pointer {
		dest = dest.Elem()
	}
	if dest.Kind() != reflect.Struct {
		return
	}

	for _, tagName := range supportedTagsList {
		for i := 0; i < dest.NumField(); i++ {
			if !dest.Type().Field(i).IsExported() {
				continue
			}

			// recurring for embedded structs
			normalizeStructTagsRec(dest.Field(i), normalizedProvider)

			// normalizing single fields
			tagKeyModifiers := getStructTag(dest.Type().Field(i), tagName)
			normalizeStructTagsForField(tagName, tagKeyModifiers, normalizedProvider)
		}
	}
}

func normalizeStructTagsForField(tagName, tagKeyModifiers string, normalizedProvider provider.TagProvider) {
	tagKey, modifiers := parseModifiers(tagKeyModifiers)

	if tagKey == "" {
		return
	}

	values := []string{}

	defaultValue, ok := modifiers[defaultModifier]
	if ok {
		values = append(values, defaultValue)
	}

	normalizedProvider.Set(tagName, tagKey, values)
	return
}

func getKey(tagName, tagKey string) string {
	return strings.TrimRight(
		strings.TrimLeft(
			strings.TrimPrefix(tagName, tagKey),
			"["),
		"]")
}

func parseModifiers(fullTag string) (string, map[string]string) {
	modifiers := make(map[string]string)
	tag, keys := unwrapTag(fullTag)
	for _, modifier := range supportedModifiersList {
		for _, key := range keys {
			if strings.HasPrefix(key, modifier) {
				if strings.Contains(key, "=") {
					keyValue := strings.Split(key, "=")
					modifiers[keyValue[0]] = keyValue[1]
				}
			}
		}
	}
	return tag, modifiers
}
