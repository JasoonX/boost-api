package modifiers

import (
	"strings"

	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/utils"
)

func ParseModifiers(fullTag string) (string, map[string]string) {
	modifiers := make(map[string]string)
	tag, keys := utils.UnwrapTag(fullTag)
	for _, modifier := range SupportedModifiersList {
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

func ApplyModifiers(values []string, modifiers map[string]string) []string {
	defaultModifierValue, ok := modifiers[DefaultModifier]
	if ok {
		if len(values) == 0 {
			return []string{defaultModifierValue}
		}
		values[0] = defaultModifierValue
		return values
	}
	return nil
}
