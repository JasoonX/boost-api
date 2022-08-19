package convert

import (
	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/common"
	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
	"github.com/BOOST-2021/boost-app-back/internal/data"
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
		Id:   news.ID.String(),
		Type: resources.NEWS,
		Attributes: resources.NewsAttributes{
			CreatedAt: news.CreatedAt,
			UpdatedAt: news.UpdatedAt,
			Media:     ToResponseNewsMedia(news.Media),
		},
		Relationships: resources.NewsRelationships{
			Author: resources.NewsRelationshipsAuthor{
				Links: ToResponseRelationLinks(resources.NEWS, resources.USERS, news.ID.String(), news.AuthorID.String()),
				Data:  ToResponseRelationItem(resources.USERS, news.AuthorID.String()),
			},
		},
	}
}

func ToResponseRelations[T data.IDer](relationType resources.EntityType, relations []T) []resources.RelatedEntity {
	relationsResponse := make([]resources.RelatedEntity, len(relations))
	for i, r := range relations {
		relationsResponse[i] = *ToResponseRelationItem(relationType, r.GetID())
	}
	return relationsResponse
}

func ToResponseRelationItem(relationType resources.EntityType, relationID string) *resources.RelatedEntity {
	return &resources.RelatedEntity{
		Type: &relationType,
		Id:   &relationID,
	}
}

func ToResponseRelationLinks(selfType, relationType resources.EntityType, selfID, relationID string) *resources.RelatedLinks {
	return &resources.RelatedLinks{
		Self:    EntityTypeToPath(selfType, selfID),
		Related: EntityTypeToPath(relationType, relationID),
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

func ToResponseUser(user *model.User) *resources.User {
	return &resources.User{
		Id:   user.ID.String(),
		Type: resources.USERS,
		Attributes: resources.UserAttributes{
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Status:    ToResponseUserStatus(user.Status),
			Role:      ToResponseUserRoles(user.Roles()),
		},
		Relationships: resources.UserRelationships{
			Emails: resources.UserRelationshipsEmails{
				Data: ToResponseRelations(resources.EMAILS, user.Emails),
			},
			Phones: &resources.UserRelationshipsPhones{
				Data: ToResponseRelations(resources.PHONES, user.Phones),
			},
		},
	}
}

func ToResponseUserEmails(emails []model.Email) []resources.Email {
	emailsResponse := make([]resources.Email, len(emails))
	for i, e := range emails {
		emailsResponse[i] = ToResponseEmailItem(e)
	}
	return emailsResponse
}

func ToResponseEmailItem(email model.Email) resources.Email {
	return resources.Email{
		Id:   email.ID.String(),
		Type: resources.EMAILS,
		Attributes: resources.EmailAttributes{
			Email:      email.Email,
			IsVerified: &email.IsVerified,
			IsPrimary:  &email.IsPrimary,
		},
		Relationships: resources.EmailRelationships{
			User: resources.EmailRelationshipsUser{
				Data: ToResponseRelationItem(resources.USERS, email.UserID.String()),
			},
		},
	}
}

func ToResponseTokenPair(tokenPair *auth.TokenPair) resources.TokenPair {
	return resources.TokenPair{
		Type: resources.TOKEN_PAIR,
		Attributes: resources.TokenPairAttributes{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		},
	}
}

func FromResponsePage(params params.PageParams) model.PageParams {
	return model.PageParams{
		Limit:  int(convert.FromPtr(params.Limit)),
		Offset: int(convert.FromPtr(params.Offset)),
	}
}
