package common

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

func RunEvery(d time.Duration, fs ...func() error) error {
	if err := runFuncs(time.Now(), fs...); err != nil {
		return errors.Wrap(err, "failed to run funcs initial")
	}
	for x := range time.Tick(d) {
		if err := runFuncs(x, fs...); err != nil {
			return errors.Wrap(err, "failed to run funcs")
		}
	}
	return nil
}

func runFuncs(x time.Time, fs ...func() error) error {
	for i, f := range fs {
		if err := f(); err != nil {
			return errors.Wrapf(err, "failed to run function: %v -> %v", x, i)
		}
	}
	return nil
}

func MustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func MustUnmarshal(b []byte, v interface{}) {
	if err := json.Unmarshal(b, v); err != nil {
		panic(err)
	}
}
