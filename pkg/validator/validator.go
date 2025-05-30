package validator

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ahargunyllib/thera-be/pkg/log"
	"github.com/bytedance/sonic"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type CustomValidatorInterface interface {
	Validate(data interface{}) ValidationErrors
}

type CustomValidatorStruct struct {
	validator *validator.Validate
	trans     ut.Translator
}

var Validator = getValidator()

func getValidator() CustomValidatorInterface {
	en := en.New()
	ut := ut.New(en, en)

	trans, found := ut.GetTranslator("en")
	if !found {
		log.Error(log.CustomLogInfo{
			"error": "translator not found",
		}, "[VALIDATOR][getValidator] Translator not found")
	}

	validator := validator.New()
	err := enTranslations.RegisterDefaultTranslations(validator, trans)
	if err != nil {
		log.Error(log.CustomLogInfo{
			"error": err.Error(),
		}, "[VALIDATOR][getValidator] Failed to register default translations")
	}

	return &CustomValidatorStruct{
		validator: validator,
		trans:     trans,
	}
}

func (v *CustomValidatorStruct) Validate(data interface{}) ValidationErrors {
	valErr := v.validator.Struct(data)
	if valErr != nil {
		var valErrs validator.ValidationErrors
		if errors.As(valErr, &valErrs) {
			// Get the type of the data struct
			dataType := reflect.TypeOf(data)
			if dataType.Kind() == reflect.Ptr {
				dataType = dataType.Elem()
			}

			body := ValidationError{
				Fields: make(map[string]FieldError),
			}
			query := ValidationError{
				Fields: make(map[string]FieldError),
			}
			param := ValidationError{
				Fields: make(map[string]FieldError),
			}
			other := ValidationError{
				Fields: make(map[string]FieldError),
			}

			for _, err := range valErrs {
				field, _ := dataType.FieldByName(err.StructField())
				tag, fieldName := getTagAndFieldName(field)

				if tag == "json" {
					body.Fields = map[string]FieldError{
						fieldName: {
							Tag:     err.Tag(),
							Message: err.Translate(v.trans),
						},
					}
					continue
				}

				if tag == "param" {
					param.Fields = map[string]FieldError{
						fieldName: {
							Tag:     err.Tag(),
							Message: err.Translate(v.trans),
						},
					}
					continue
				}

				if tag == "query" {
					query.Fields = map[string]FieldError{
						fieldName: {
							Tag:     err.Tag(),
							Message: err.Translate(v.trans),
						},
					}
					continue
				}

				other.Fields = map[string]FieldError{
					fieldName: {
						Tag:     err.Tag(),
						Message: err.Translate(v.trans),
					},
				}
			}

			body.Message = fmt.Sprintf("%d validation error(s) in body", len(body.Fields))
			param.Message = fmt.Sprintf("%d validation error(s) in param", len(param.Fields))
			query.Message = fmt.Sprintf("%d validation error(s) in query", len(query.Fields))
			other.Message = fmt.Sprintf("%d validation error(s) in others", len(other.Fields))

			res := ValidationErrors{
				"body":   body,
				"param":  param,
				"query":  query,
				"others": other,
			}

			return res
		}

		log.Error(log.CustomLogInfo{
			"error": valErr.Error(),
		}, "[VALIDATOR][Validate] Failed to validate data")
	}

	return nil
}

type FieldError struct {
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

type ValidationError struct {
	Message string                `json:"message"`
	Fields  map[string]FieldError `json:"fields"`
}

type ValidationErrors map[string]ValidationError

func (v ValidationErrors) Error() string {
	j, err := sonic.Marshal(v)
	if err != nil {
		return ""
	}

	return string(j)
}

func (v ValidationErrors) Serialize() any {
	return v
}

func getTagAndFieldName(field reflect.StructField) (string, string) {
	checkTags := []string{"json", "query", "param"}
	for _, tag := range checkTags {
		fieldName, ok := field.Tag.Lookup(tag)
		if ok {
			return tag, fieldName
		}
	}

	return "", field.Name
}
