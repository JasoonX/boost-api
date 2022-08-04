package responses

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/params"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func Test_Links(t *testing.T) {
	address := &url.URL{
		Path: "/api/v1",
	}

	testCases := []struct {
		name     string
		expLinks *resources.Links
		input    params.PageParams
	}{
		{
			name: "OK default",
			expLinks: &resources.Links{
				Self: fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"10"},
				}.Encode()),
				Next: fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"20"},
				}.Encode()),
				Prev: fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"0"},
				}.Encode()),
				First: convert.ToPtr(fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"0"},
				}.Encode())),
				Last: convert.ToPtr(fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"10"},
				}.Encode())),
			},
			input: params.PageParams{
				Limit:  convert.ToPtr[int32](10),
				Offset: convert.ToPtr[int32](10),
				Total:  convert.ToPtr[int32](20),
			},
		},
		{
			name: "OK with next equals self",
			expLinks: &resources.Links{
				Self: fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"10"},
				}.Encode()),
				Next: fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"10"},
				}.Encode()),
				Prev: fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"0"},
				}.Encode()),
				First: nil,
				Last:  nil,
			},
			input: params.PageParams{
				Limit:  convert.ToPtr[int32](10),
				Offset: convert.ToPtr[int32](10),
				Total:  convert.ToPtr[int32](10),
			},
		},
		{
			name: "OK with prev equals self",
			expLinks: &resources.Links{
				Self: fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"0"},
				}.Encode()),
				Next: fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"10"},
				}.Encode()),
				Prev: fmt.Sprintf("%s?%s", address.String(), url.Values{
					"page[limit]":  []string{"10"},
					"page[offset]": []string{"0"},
				}.Encode()),
				First: nil,
				Last:  nil,
			},
			input: params.PageParams{
				Limit:  convert.ToPtr[int32](10),
				Offset: convert.ToPtr[int32](0),
				Total:  convert.ToPtr[int32](10),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, _ := http.NewRequest("GET", address.String(), nil)
			require.Equal(t, tc.expLinks, BuildLinks(r, tc.input))
		})
	}
}
