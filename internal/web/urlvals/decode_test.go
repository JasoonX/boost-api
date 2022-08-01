package urlvals

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
)

type testEmbeded struct {
	PageParams
}

type testFields struct {
	Limit  *int32 `page:"limit"`
	Offset *int32 `page:"offset"`
}

type testUnexported struct {
	Limit    *int32 `page:"limit"`
	Offset   *int32 `page:"offset"`
	pageSize *int32 `page:"size"`
}

func TestDecode(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  interface{}
	}{
		{
			name: "ok page params embedded struct",
			in:   "page[limit]=12&page[offset]=12",
			out: testEmbeded{
				PageParams: PageParams{
					Limit:  convert.Int32Ptr(12),
					Offset: convert.Int32Ptr(12),
				},
			},
		},
		{
			name: "ok page params fields",
			in:   "page[limit]=12&page[offset]=12",
			out: testFields{
				Limit:  convert.Int32Ptr(12),
				Offset: convert.Int32Ptr(12),
			},
		},
		{
			name: "ok private field not populated",
			in:   "page[limit]=12&page[offset]=12&page[size]=12",
			out: testUnexported{
				Limit:    convert.Int32Ptr(12),
				Offset:   convert.Int32Ptr(12),
				pageSize: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := url.ParseQuery(tc.in)
			require.NoError(t, err)

			switch tc.out.(type) {
			case testEmbeded:
				out := testEmbeded{}
				err = Decode(values, &out)
				require.NoError(t, err)
				require.Equal(t, tc.out, out)
			case testFields:
				out := testFields{}
				err = Decode(values, &out)
				require.NoError(t, err)
				require.Equal(t, tc.out, out)
			case testUnexported:
				out := testUnexported{}
				err = Decode(values, &out)
				require.NoError(t, err)
				require.Equal(t, tc.out, out)
			}
		})
	}
}
