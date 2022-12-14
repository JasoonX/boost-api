package postgres

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("not found")

func ToTsVector(locale, text string) string {
	return fmt.Sprintf("to_tsvector('%s', %s)", locale, text)
}

func SetWeight(text, weight string) string {
	return fmt.Sprintf("setweight(%s, '%s')", text, weight)
}
