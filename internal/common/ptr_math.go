package common

import (
	"golang.org/x/exp/constraints"

	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
)

func SumPtr[T constraints.Integer](a, b *T) *T {
	return convert.ToPtr(convert.FromPtr(a) + convert.FromPtr(b))
}

func SubPtr[T constraints.Integer](a, b *T) *T {
	return convert.ToPtr(convert.FromPtr(a) - convert.FromPtr(b))
}

func MinOrFloor[T constraints.Integer](a *T, floor *T) *T {
	if a == nil {
		return floor
	}
	if floor != nil && *a < *floor {
		return floor
	}
	return a
}

func MaxOrCeil[T constraints.Integer](a *T, ceil *T) *T {
	if a == nil {
		return ceil
	}
	if ceil != nil && *a > *ceil {
		return ceil
	}
	return a
}
