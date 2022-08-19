////go:build fake

package fake

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/BOOST-2021/boost-app-back/internal/auth"
	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
	"github.com/BOOST-2021/boost-app-back/internal/common/debug"
	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/internal/data/store"
)

// this package is ment to be used for testing purposes
// this allows easily run local dev env, with local data
// useful for the testing, which might affect the database
// as well as for integration testing

type Generator interface {
	Users(n int) ([]model.User, error)
	News(n int) ([]model.News, error)
	Emails(n int) ([]model.Email, error)
	Phones(n int) ([]model.Phone, error)

	Default() error
}

const (
	seed = 123
)

type generator struct {
	data  store.DataProvider
	cfg   *Config
	log   *logrus.Entry
	faker *gofakeit.Faker
}

func (g *generator) Users(n int) ([]model.User, error) {
	ctx := context.TODO()
	users := make([]model.User, n)
	for i := 0; i < n; i++ {
		passHash, _ := auth.HashPassword(g.faker.Password(true, true, true, true, true, 20))
		users[i] = model.User{
			Username:     convert.ToPtr(g.faker.Username()),
			Status:       model.UserStatuses[rand.Int()%len(model.UserStatuses)],
			Role:         pq.StringArray{string(model.UserRoles[rand.Int()%len(model.UserRoles)])},
			PasswordHash: passHash,
		}
	}

	if err := g.data.UsersProvider().AddUsersBatch(ctx, users); err != nil {
		return nil, errors.Wrap(err, "failed to add users")
	}
	return users, nil
}

func (g *generator) News(n int) ([]model.News, error) {
	ctx := context.TODO()
	users, err := g.data.UsersProvider().ListUsers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list users")

	}
	news := make([]model.News, n)
	for i := 0; i < n; i++ {
		news[i] = model.News{
			HideMetrics:        rand.Int()%2 == 0,
			ViewsCount:         rand.Int63(),
			ReactionsCount:     rand.Int63(),
			HideViewsCount:     rand.Int()%2 == 0,
			HideReactionsCount: rand.Int()%2 == 0,
			Media: &model.NewsMedia{
				Title: convert.ToPtr(fmt.Sprintf("%s %s %s", g.faker.Noun(), g.faker.VerbAction(), g.faker.NounConcrete())),
				Text:  convert.ToPtr(g.faker.Quote()),
			},
			AuthorID: users[rand.Int()%len(users)].ID,
		}
	}

	if err := g.data.NewsProvider().AddNewsBatch(ctx, news); err != nil {
		return nil, errors.Wrap(err, "failed to add news")
	}
	return news, nil
}

func (g *generator) Emails(n int) ([]model.Email, error) {
	ctx := context.TODO()
	users, err := g.data.UsersProvider().ListUsers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}

	emails := make([]model.Email, n)
	for i := 0; i < n; i++ {
		emails[i] = model.Email{
			Email:      g.faker.Email(),
			IsVerified: rand.Int()%2 == 0,
			IsPrimary:  false,
			UserID:     users[rand.Int()%len(users)].ID,
		}
	}

	// if is set to primary, then should always be verified
	emails[0].IsPrimary = true
	emails[0].IsVerified = true

	if err := g.data.EmailsProvider().AddEmailsBatch(ctx, emails); err != nil {
		return nil, errors.Wrap(err, "failed to add emails")
	}
	return emails, nil
}

func (g *generator) Phones(n int) ([]model.Phone, error) {
	ctx := context.TODO()
	users, err := g.data.UsersProvider().ListUsers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}

	phones := make([]model.Phone, n)
	for i := 0; i < n; i++ {
		phones[i] = model.Phone{
			SubscriberNumber: g.faker.Phone(),
			CountryCode:      convert.ToPtr("380"),
			IsVerified:       rand.Int()%2 == 0,
			IsPrimary:        false,
			UserID:           users[rand.Int()%len(users)].ID,
		}
	}

	// if is set to primary, then should always be verified
	phones[0].IsPrimary = true
	phones[0].IsVerified = true

	if err := g.data.PhonesProvider().AddPhonesBatch(ctx, phones); err != nil {
		return nil, errors.Wrap(err, "failed to add phones")
	}
	return phones, nil
}

func (g *generator) Locations(n int) ([]model.Location, error) {
	ctx := context.TODO()
	countries, err := g.data.EnumProvider().ListCountries(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list countries")
	}

	locations := make([]model.Location, n)
	for i := 0; i < n; i++ {
		locations[i] = model.Location{
			CountryCode: convert.ToPtr(countries[rand.Int()%len(countries)].Code),
			Region:      convert.ToPtr(g.faker.State()),
			City:        convert.ToPtr(g.faker.City()),
			Street:      convert.ToPtr(g.faker.Street()),
		}
	}

	if err := g.data.LocationsProvider().AddLocationsBatch(ctx, locations); err != nil {
		return nil, errors.Wrap(err, "failed to add locations")
	}
	return locations, nil
}

func (g *generator) UserLocations(n int) ([]model.UserLocation, error) {
	ctx := context.TODO()
	users, err := g.data.UsersProvider().ListUsers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}

	locations, err := g.data.LocationsProvider().ListLocations(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list locations")
	}

	userLocations := make([]model.UserLocation, n)
	for i := 0; i < n; i++ {
		userLocations[i] = model.UserLocation{
			UserID:     users[rand.Int()%len(users)].ID,
			LocationID: locations[rand.Int()%len(locations)].ID,
		}
	}

	if err := g.data.UserLocationsProvider().AddUserLocationsBatch(ctx, userLocations); err != nil {
		return nil, errors.Wrap(err, "failed to add user locations")
	}
	return userLocations, nil
}

func logGenerated[T interface{}](log *logrus.Entry, entities []T, entityName string) {
	log.Infof("%s generated: %d", entityName, len(entities))
	for _, u := range entities {
		log.Debug(debug.Struct(u))
	}
}

func (g *generator) Default() error {
	rand.Seed(seed)

	users, err := g.Users(g.cfg.Users)
	if err != nil {
		return errors.Wrap(err, "failed to generate users")
	}
	logGenerated(g.log, users, "users")

	news, err := g.News(g.cfg.News)
	if err != nil {
		return errors.Wrap(err, "failed to generate news")
	}
	logGenerated(g.log, news, "news")

	emails, err := g.Emails(g.cfg.Emails)
	if err != nil {
		return errors.Wrap(err, "failed to generate emails")
	}
	logGenerated(g.log, emails, "emails")

	phones, err := g.Phones(g.cfg.Phones)
	if err != nil {
		return errors.Wrap(err, "failed to generate phones")
	}
	logGenerated(g.log, phones, "phones")

	locations, err := g.Locations(g.cfg.Locations)
	if err != nil {
		return errors.Wrap(err, "failed to generate locations")
	}
	logGenerated(g.log, locations, "locations")

	userLocations, err := g.UserLocations(g.cfg.UserLocations)
	if err != nil {
		return errors.Wrap(err, "failed to generate user locations")
	}
	logGenerated(g.log, userLocations, "user locations")

	return nil
}

func New(appCfg config.Config, cfg *Config) Generator {
	return &generator{
		data:  store.New(appCfg),
		cfg:   cfg,
		log:   appCfg.Logging(),
		faker: gofakeit.New(seed),
	}
}
