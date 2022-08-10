package news

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/BOOST-2021/boost-app-back/internal/common"
	"github.com/BOOST-2021/boost-app-back/internal/common/enum"
	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres"
	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/internal/data/queriers"
)

const News = "news"

var newsModel = model.News{}

// fields
const newsTsvText = "tsv_text"

type newsProvider struct {
	log *logrus.Entry
	db  *gorm.DB
}

func New(cfg config.Config) queriers.NewsProvider {
	return &newsProvider{
		db:  cfg.DB(),
		log: cfg.Logging().WithField(postgres.Querier, News),
	}
}

func (p newsProvider) CountNews(_ context.Context) (int64, error) {
	var count int64
	if err := p.db.Model(newsModel).Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, "failed to count news")
	}
	return count, nil
}

func (p newsProvider) OfPage(_ context.Context, pageParams model.PageParams) queriers.NewsProvider {
	p.db = p.db.Offset(pageParams.Offset).Limit(pageParams.Limit)
	return p
}

func (p newsProvider) ListNews(_ context.Context) ([]model.News, error) {
	var out []model.News
	if err := p.db.Model(newsModel).Find(&out).Error; err != nil {
		return nil, errors.Wrap(err, "failed to list news")
	}
	return out, nil
}

func (p newsProvider) IndexNews(ctx context.Context, id uuid.UUID, title, text, meta string) error {
	locale := enum.LocaleFromISO2(common.GetLocaleFromContext(ctx))
	err := p.db.Model(model.News{
		ID: id,
	}).UpdateColumn(
		fmt.Sprintf("%s_%s", newsTsvText, locale),
		gorm.Expr(fmt.Sprintf("%s||%s||%s"),
			postgres.SetWeight(postgres.ToTsVector(locale, title), postgres.WeightA),
			postgres.SetWeight(postgres.ToTsVector(locale, text), postgres.WeightB),
			postgres.SetWeight(postgres.ToTsVector(locale, meta), postgres.WeightC),
		),
	).Error
	if err != nil {
		return errors.Wrap(err, "failed to index news")
	}
	return nil
}

func (p newsProvider) AddNewsBatch(_ context.Context, news []model.News) error {
	if err := p.db.Omit(postgres.ID).Create(&news).Error; err != nil {
		return errors.Wrap(err, "failed to insert news")
	}
	return nil
}
