package data

import "github.com/pkg/errors"

var ErrNotFound = errors.New("not found")

type IDer interface {
	GetID() string
}
