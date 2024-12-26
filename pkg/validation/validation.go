package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/sajadsalem/startkit/pkg/modules"
)

func ValidateStruct(validate *validator.Validate, request interface{}) ([]map[string]string, error) {
	errors := []map[string]string{}

	// Validate the input struct
	err := validate.Struct(request)
	if err != nil {
		// Iterate over the validation errors
		for _, err := range err.(validator.ValidationErrors) {
			// Add an error to the errors array
			errors = append(errors, map[string]string{
				err.Field(): err.Error(),
			})
		}

		return errors, err
	}

	return nil, nil
}

func phoneValidation(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	_, err := modules.CleanPhoneNumber(phone)

	return err == nil
}

func RegisterCustomValidations(v *validator.Validate) {
	// phone validation
	v.RegisterValidation("phone", phoneValidation)
}
