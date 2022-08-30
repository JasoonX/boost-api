package resources

import (
	"unicode"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var _ validation.Validatable = AddUser{}
var _ validation.Validatable = AddUserAttributes{}
var _ validation.Validatable = AddPhoneAttributes{}

const (
	passwordMinLength = 8
)

func isValidPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= passwordMinLength {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// TODO: https://github.com/datasets/country-codes/blob/master/data/country-codes.csv
func isCountryCodeNumeric(c string) bool {
	return true
}

func (o AddUser) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Type, validation.In(USERS)),
		validation.Field(&o.Attributes),
	)
}

func (o AddUserAttributes) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Email, validation.Required, is.Email),
		validation.Field(&o.Password, validation.Required, validation.NewStringRule(isValidPassword, "Invalid password")),
		validation.Field(&o.FirstName, validation.When(o.HasFirstName(), is.Alpha)),
		validation.Field(&o.LastName, validation.When(o.HasLastName(), is.Alpha)),
		validation.Field(&o.Phone, validation.When(o.HasPhone())),
	)
}

func (o AddPhoneAttributes) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.SubscriberNumber, is.E164),
		validation.Field(&o.CountryCode, validation.NewStringRule(isCountryCodeNumeric, "Not a valid country code")),
		validation.Field(&o.Extension, validation.When(o.HasExtension(), is.E164)),
	)
}
