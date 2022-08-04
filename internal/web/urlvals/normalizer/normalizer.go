package normalizer

import (
	"net/url"
	"reflect"
	"strings"

	"github.com/BOOST-2021/boost-app-back/internal/common"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/modifiers"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/provider"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/utils"
)

// NormalizationMode defines which has higher priority, field set value, or modifiers (varies for encoder and decoder)
type NormalizationMode int

const (
	NormalizationModeDecoding NormalizationMode = iota
	NormalizationModeEncoding
)

// Normalizer TODO: refactor, consider using for something beyond urlvals
type Normalizer interface {
	NormalizeUrlQuery(urlQuery url.Values) provider.TagProvider
	NormalizeStructTags(i interface{}) provider.TagProvider
}

type normalizer struct {
	mode        NormalizationMode
	tagProvider provider.TagProvider
	tagsList    []string
}

func New(tagsList []string, mode NormalizationMode) Normalizer {
	return &normalizer{
		tagsList:    tagsList,
		tagProvider: provider.New(),
		mode:        mode,
	}
}

func (n *normalizer) NormalizeUrlQuery(urlQuery url.Values) provider.TagProvider {
	normalizedQueries := provider.New()
	for k, v := range urlQuery {
		for _, tagName := range n.tagsList {
			if strings.HasPrefix(k, tagName) {
				tagKey := utils.GetKey(k, tagName)
				normalizedQueries.Set(tagName, tagKey, v)
			}
		}
	}
	return normalizedQueries
}

func (n *normalizer) NormalizeStructTags(i interface{}) provider.TagProvider {
	n.normalizeStructTagsRec(reflect.ValueOf(i))
	return n.tagProvider
}

func (n *normalizer) normalizeStructTagsRec(dest reflect.Value) {
	dest = reflect.Indirect(dest)

	if dest.Kind() != reflect.Struct {
		return
	}

	for _, tagName := range n.tagsList {
		for i := 0; i < dest.NumField(); i++ {
			if !dest.Type().Field(i).IsExported() {
				continue
			}

			// recurring for embedded structs
			n.normalizeStructTagsRec(dest.Field(i))

			// normalizing single fields
			tagKeyModifiers := utils.GetStructTag(dest.Type().Field(i), tagName)
			n.normalizeStructTagsForField(dest.Field(i), tagName, tagKeyModifiers)
		}
	}
}

func (n *normalizer) normalizeStructTagsForField(field reflect.Value, tagName, tagKeyModifiers string) {
	field = reflect.Indirect(field)

	tagKey, modifiersMap := modifiers.ParseModifiers(tagKeyModifiers)

	if tagKey == "" {
		return
	}

	var values []string

	runFuncs := []func(){
		func() {
			values = modifiers.ApplyModifiers(values, modifiersMap)
		},
		func() {
			if field.IsValid() {
				if len(values) == 0 {
					values = make([]string, 1)
				}
				values[0] = utils.ToString(field.Interface())
			}
		}}

	if n.mode == NormalizationModeEncoding {
		for _, runFunc := range runFuncs {
			runFunc()
		}
	} else if n.mode == NormalizationModeDecoding {
		for _, runFunc := range common.Reverse(runFuncs) {
			runFunc()
		}
	}

	n.tagProvider.Set(tagName, tagKey, values)
	return
}
