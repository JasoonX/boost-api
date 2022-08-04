package urlvals

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/params"
)

func Test_Encode(t *testing.T) {
	testCases := []struct {
		name string
		in   interface{}
		out  string
	}{
		{
			name: "ok page params embedded struct",
			in: struct {
				params.PageParams
			}{
				PageParams: params.PageParams{
					Limit:  convert.ToPtr[int32](12),
					Offset: convert.ToPtr[int32](12),
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
				Limit:  convert.ToPtr[int32](12),
				Offset: convert.ToPtr[int32](12),
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
				Limit:    convert.ToPtr[int32](12),
				Offset:   convert.ToPtr[int32](12),
				pageSize: convert.ToPtr[int32](12),
			},
			out: "page[limit]=12&page[offset]=12",
		},
		{
			name: "ok not pointer",
			in: struct {
				Limit int32 `page:"limit"`
			}{
				Limit: 10,
			},
			out: "page[limit]=10",
		},
		{
			name: "ok zero value",
			in: struct {
				Limit int32 `page:"limit"`
			}{
				Limit: 0,
			},
			out: "page[limit]=0",
		},
		{
			name: "ok slices",
			in: struct {
				Cities []string `page:"cities"`
				Ids    []int64  `page:"ids"`
			}{
				Cities: []string{"Moscow", "London"},
				Ids:    []int64{1, 2, 3},
			},
			out: "page[cities]=Moscow,London&page[ids]=1,2,3",
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
