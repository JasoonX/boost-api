package convert

import (
	"github.com/BOOST-2021/boost-app-back/internal/common"
	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/internal/web/urlvals/params"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func ToResponseNews(news []model.News) []resources.News {
	newsResponse := make([]resources.News, len(news))
	for i, n := range news {
		newsResponse[i] = ToResponseNewsItem(n)
	}
	return newsResponse
}

func ToResponseNewsItem(news model.News) resources.News {
	return resources.News{
		Id:        news.ID.String(),
		AuthorId:  news.AuthorID.String(),
		CreatedAt: news.CreatedAt,
		UpdatedAt: news.UpdatedAt,
		Media:     ToResponseNewsMedia(news.Media),
	}
}

func ToResponseNewsMedia(media *model.NewsMedia) *resources.NewsMedia {
	return &resources.NewsMedia{
		Title:     media.Title,
		Text:      media.Text,
		Resources: ToResponseNewsMediaResources(media.Resources),
	}
}

func ToResponseNewsMediaResources(mediaResources []model.NewsMediaResource) []resources.NewsMediaResource {
	newsMediaResources := make([]resources.NewsMediaResource, len(mediaResources))
	for i, r := range mediaResources {
		newsMediaResources[i] = ToResponseNewsMediaResourceItem(r)
	}
	return newsMediaResources
}

func ToResponseNewsMediaResourceItem(r model.NewsMediaResource) resources.NewsMediaResource {
	var metaMap map[string]interface{}
	if r.Meta != nil {
		metaMap = make(map[string]interface{})
		common.MustUnmarshal(r.Meta, &metaMap)
	}
	return resources.NewsMediaResource{
		Type: r.Type,
		Url:  r.URL,
		Meta: metaMap,
	}
}

func ToResponsePage(page params.PageParams) *resources.Page {
	return &resources.Page{
		Offset: page.Offset,
		Limit:  page.Limit,
		Total:  page.Total,
	}
}
