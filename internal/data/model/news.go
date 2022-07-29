package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type News struct {
	ID        uuid.UUID  `gorm:"column:id"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`

	// Metrics
	HideMetrics        bool  `gorm:"column:hide_metrics"`
	ViewsCount         int64 `gorm:"column:views_count"`
	ReactionsCount     int64 `gorm:"column:reactions_count"`
	HideViewsCount     bool  `gorm:"column:hide_views_count"`
	HideReactionsCount bool  `gorm:"column:hide_reactions_count"`

	// Data
	Media    *NewsMedia `gorm:"column:media"`
	AuthorID uuid.UUID  `gorm:"column:author_id"`
}

// NewsMedia struct for media content, can be a nearly free form json
type NewsMedia struct {
	Title     *string             `json:"title"`
	Text      *string             `json:"text"`
	Resources []NewsMediaResource `json:"resources"`
}

type NewsMediaResource struct {
	Type *string         `json:"type"`
	URL  *string         `json:"url"`
	Meta json.RawMessage `json:"meta"`
}

func (a NewsMedia) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *NewsMedia) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func (News) TableName() string {
	return "news"
}
