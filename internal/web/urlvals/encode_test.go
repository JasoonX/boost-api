package urlvals

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
)

func TestEncode(t *testing.T) {
	testCases := []struct {
		name string
		in   interface{}
		out  string
	}{
		{
			name: "ok page params embedded struct",
			in: struct {
				PageParams
			}{
				PageParams: PageParams{
					Limit:  convert.Int32Ptr(12),
					Offset: convert.Int32Ptr(12),
				},
			},
			out: "page[limit]=12&page[offset]=12",
		},
		{
			name: "ok page params fields",
			in: struct {
				Limit  *int32 `page:"limit"`
				Offset *int32 `page:"offset"`
			}{
				Limit:  convert.Int32Ptr(12),
				Offset: convert.Int32Ptr(12),
			},
			out: "page[limit]=12&page[offset]=12",
		},
		{
			name: "ok private field not populated",
			in: struct {
				Limit    *int32 `page:"limit"`
				Offset   *int32 `page:"offset"`
				pageSize *int32 `page:"page_size"`
			}{
				Limit:    convert.Int32Ptr(12),
				Offset:   convert.Int32Ptr(12),
				pageSize: convert.Int32Ptr(12),
			},
			out: "page[limit]=12&page[offset]=12",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			decodedValue, err := url.QueryUnescape(Encode(tc.in))
			require.NoError(t, err)
			require.Equal(t, tc.out, decodedValue)
		})
	}
}
