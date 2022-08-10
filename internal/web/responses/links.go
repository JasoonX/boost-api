package responses

import (
	"net/http"
	"net/url"

	"github.com/BOOST-2021/boost-app-back/internal/common"
	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/params"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func BuildLinks(r *http.Request, page params.PageParams) *resources.Links {
	urlBuilder := func(rawQuery string) *url.URL {
		return &url.URL{
			Path:     r.URL.Path,
			RawQuery: rawQuery,
		}
	}

	selfQuery := urlvals.Encode(page)
	nextQuery := urlvals.Encode(params.PageParams{
		Limit:  page.Limit,
		Offset: common.MaxOrCeil(common.SumPtr(page.Offset, page.Limit), page.Total),
	})
	prevQuery := urlvals.Encode(params.PageParams{
		Limit:  page.Limit,
		Offset: common.MinOrFloor(common.SubPtr(page.Offset, page.Limit), convert.ToPtr[int64](0)),
	})

	firstQuery := urlvals.Encode(params.PageParams{
		Limit:  page.Limit,
		Offset: convert.ToPtr[int64](0),
	})

	var lastQuery *string
	if page.Total != nil && page.Limit != nil {
		lastQuery = convert.ToPtr(
			urlBuilder(urlvals.Encode(params.PageParams{
				Limit:  page.Limit,
				Offset: convert.ToPtr(lastPageOffset(*page.Total, *page.Limit)),
			})).String(),
		)
	}

	return &resources.Links{
		Self:  urlBuilder(selfQuery).String(),
		Next:  urlBuilder(nextQuery).String(),
		Prev:  urlBuilder(prevQuery).String(),
		First: convert.ToPtr(urlBuilder(firstQuery).String()),
		Last:  lastQuery,
	}
}

func lastPageOffset(total, limit int64) int64 {
	return total - (total % limit)
}
