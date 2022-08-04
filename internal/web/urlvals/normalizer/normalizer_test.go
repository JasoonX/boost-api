package normalizer

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/provider"
)

const (
	searchTag = "search"
	filterTag = "filter"
	pageTag   = "page"
)

func Test_NormalizeUrlQuery(t *testing.T) {
	urlQuery := url.Values{}
	urlQuery.Add("page[limit]", "10")
	urlQuery.Add("page[offset]", "20")
	urlQuery.Add("page[size]", "30")
	urlQuery.Add("search[first_name]", "John")
	urlQuery.Add("search[last_name]", "Doe")
	urlQuery.Add("filter[age]", "30")

	n := New([]string{}, NormalizationModeEncoding)
	act := n.NormalizeUrlQuery(urlQuery)

	exp := provider.TagProvider{
		pageTag: provider.TagKeyProvider{
			"limit": provider.NormalizedTagParam{
				TagName: pageTag,
				TagKey:  "limit",
				Values:  []string{"10"},
			},
			"offset": provider.NormalizedTagParam{
				TagName: pageTag,
				TagKey:  "offset",
				Values:  []string{"20"},
			},
			"size": provider.NormalizedTagParam{
				TagName: pageTag,
				TagKey:  "size",
				Values:  []string{"30"},
			},
		},
		searchTag: provider.TagKeyProvider{
			"first_name": provider.NormalizedTagParam{
				TagName: searchTag,
				TagKey:  "first_name",
				Values:  []string{"John"},
			},
			"last_name": provider.NormalizedTagParam{
				TagName: searchTag,
				TagKey:  "last_name",
				Values:  []string{"Doe"},
			},
		},
		filterTag: provider.TagKeyProvider{
			"age": provider.NormalizedTagParam{
				TagName: filterTag,
				TagKey:  "age",
				Values:  []string{"30"},
			},
		},
	}

	for tagName, v := range exp {
		for tagKey := range v {
			require.Equal(t, exp[tagName][tagKey], act[tagName][tagKey])
		}
	}
}

func Test_NormalizeStructTags(t *testing.T) {
	type Test struct {
		Limit     int    `page:"limit,default=10"`
		Offset    int    `page:"offset,default=20"`
		Size      int    `page:"size,default=30"`
		FirstName string `search:"first_name,default=John"`
		LastName  string `search:"last_name,default=Doe"`
		Age       int    `filter:"age,default=30"`
		City      string `filter:"city"`
	}

	n := New([]string{}, NormalizationModeEncoding)
	act := n.NormalizeStructTags(reflect.ValueOf(Test{}))

	exp := provider.TagProvider{
		pageTag: provider.TagKeyProvider{
			"limit": provider.NormalizedTagParam{
				TagName: pageTag,
				TagKey:  "limit",
				Values:  []string{"10"},
			},
			"offset": provider.NormalizedTagParam{
				TagName: pageTag,
				TagKey:  "offset",
				Values:  []string{"20"},
			},
			"size": provider.NormalizedTagParam{
				TagName: pageTag,
				TagKey:  "size",
				Values:  []string{"30"},
			},
		},
		searchTag: provider.TagKeyProvider{
			"first_name": provider.NormalizedTagParam{
				TagName: searchTag,
				TagKey:  "first_name",
				Values:  []string{"John"},
			},
			"last_name": provider.NormalizedTagParam{
				TagName: searchTag,
				TagKey:  "last_name",
				Values:  []string{"Doe"},
			},
		},
		filterTag: provider.TagKeyProvider{
			"age": provider.NormalizedTagParam{
				TagName: filterTag,
				TagKey:  "age",
				Values:  []string{"30"},
			},
			"city": provider.NormalizedTagParam{
				TagName: filterTag,
				TagKey:  "city",
			},
		},
	}

	for tagName, v := range exp {
		for tagKey := range v {
			require.Equal(t, exp[tagName][tagKey], act[tagName][tagKey])
		}
	}
}
