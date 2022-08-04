package urlvals

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/params"
)

type testEmbeded struct {
	params.PageParams
}

type testFields struct {
	Limit  *int32 `page:"limit"`
	Offset *int32 `page:"offset"`
}

type testUnexported struct {
	Limit  *int32 `page:"limit"`
	Offset *int32 `page:"offset"`
	size   *int32 `page:"size"`
}

type testDefaultValues struct {
	Limit  *int32 `page:"limit,default=12"`
	Offset *int32 `page:"offset,default=12"`
	Size   *int32 `page:"size,default=12"`
}

type stringAlias string

type testAlias struct {
	Limit stringAlias `page:"limit"`
}

type testStruct struct {
	Limit int32  `page:"limit"`
	City  string `filter:"city"`
	Ids   []int  `filter:"ids"`
}

func Test_Decode(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  interface{}
	}{
		{
			name: "ok page params embedded struct",
			in:   "page[limit]=12&page[offset]=12",
			out: testEmbeded{
				PageParams: params.PageParams{
					Limit:  convert.ToPtr[int32](12),
					Offset: convert.ToPtr[int32](12),
				},
			},
		},
		{
			name: "ok page params fields",
			in:   "page[limit]=12&page[offset]=12",
			out: testFields{
				Limit:  convert.ToPtr[int32](12),
				Offset: convert.ToPtr[int32](12),
			},
		},
		{
			name: "ok private field not populated",
			in:   "page[limit]=12&page[offset]=12&page[size]=12",
			out: testUnexported{
				Limit:  convert.ToPtr[int32](12),
				Offset: convert.ToPtr[int32](12),
				size:   nil,
			},
		},
		{
			name: "ok default value",
			in:   "",
			out: testDefaultValues{
				Limit:  convert.ToPtr[int32](12),
				Offset: convert.ToPtr[int32](12),
				Size:   convert.ToPtr[int32](12),
			},
		},
		{
			name: "ok alias",
			in:   "page[limit]=12",
			out: testAlias{
				Limit: "12",
			},
		},
		{
			name: "ok zero value",
			in:   "page[limit]=0&filter[city]=&filter[ids]=1,2,3",
			out: testStruct{
				Limit: 0,
				City:  "",
				Ids:   []int{1, 2, 3},
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
			case testDefaultValues:
				out := testDefaultValues{}
				err = Decode(values, &out)
				require.NoError(t, err)
				require.Equal(t, tc.out, out)
			case testAlias:
				out := testAlias{}
				err = Decode(values, &out)
				require.NoError(t, err)
				require.Equal(t, tc.out, out)
			case testStruct:
				out := testStruct{}
				err = Decode(values, &out)
				require.NoError(t, err)
				require.Equal(t, tc.out, out)
			}
		})
	}
}
