package urlvals

import (
	"fmt"
	"net/url"

	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/normalizer"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/utils"
)

func Encode(src interface{}) string {
	values := url.Values{}

	encodeNormalizer := normalizer.New(supportedTagsList, normalizer.NormalizationModeEncoding)

	normalizedTags := encodeNormalizer.NormalizeStructTags(src)
	for tagName, tagKeys := range normalizedTags {
		for tagKey, tagParam := range tagKeys {
			tag := fmt.Sprintf("%s[%s]", tagName, tagKey)
			if len(tagParam.Values) == 1 {
				setUrlValues(values, tag, tagParam.Values[0])
			} else if len(tagParam.Values) > 1 {
				setUrlValues(values, tag, tagParam.Values)
			}
		}
	}

	return values.Encode()
}

func setUrlValues(values url.Values, key string, value interface{}) {
	if strVal := utils.ToString(value); strVal != "" {
		values.Set(key, strVal)
	}
}
