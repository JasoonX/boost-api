package utils

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/BOOST-2021/boost-app-back/internal/common/convert"
	"github.com/BOOST-2021/boost-app-back/internal/web/responses"
	"github.com/BOOST-2021/boost-app-back/resources"
)

// UnwrapValidationErrors returns JSONServerErrors and nil, or nil and passed error if error is not a valid validation.Errors
func UnwrapValidationErrors(err error) (responses.JSONServerErrors, error) {
	res := responses.JSONServerErrors{}
	errorsMap := make(map[string]validation.Error)
	uErr := unwrapValidationErrors("", err, errorsMap)
	if uErr != nil {
		return nil, uErr
	}
	for source, valErr := range errorsMap {
		res = append(res, &resources.Error{
			Code: valErr.Code(),
			Source: map[string]interface{}{
				"source": source,
			},
			Title:  convert.ToPtr(valErr.Message()),
			Detail: valErr.Params(),
		})
	}
	return res, nil
}

func unwrapValidationErrors(path string, err error, out map[string]validation.Error) error {
	if valErr, ok := err.(validation.Error); ok {
		out[path] = valErr
		return nil
	}
	if valErrors, ok := err.(validation.Errors); ok {
		for source, valErrRaw := range valErrors {
			uErr := unwrapValidationErrors(fmt.Sprintf("%s/%s", path, source), valErrRaw, out)
			if uErr != nil {
				return uErr
			}
		}
		return nil
	}
	return err
}
