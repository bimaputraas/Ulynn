package helper

import "github.com/go-playground/validator/v10"

func Validate(s interface{}) error{
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}