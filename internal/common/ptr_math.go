package common

import "github.com/BOOST-2021/boost-app-back/internal/common/convert"

func SumPtr[T Integer](a, b *T) *T {
	return convert.ToPtr(convert.FromPtr(a) + convert.FromPtr(b))
}

func SubPtr[T Integer](a, b *T) *T {
	return convert.ToPtr(convert.FromPtr(a) - convert.FromPtr(b))
}

func MinOrFloor[T Integer](a *T, floor *T) *T {
	if a == nil {
		return floor
	}
	if floor != nil && *a < *floor {
		return floor
	}
	return a
}

func MaxOrCeil[T Integer](a *T, ceil *T) *T {
	if a == nil {
		return ceil
	}
	if ceil != nil && *a > *ceil {
		return ceil
	}
	return a
}
