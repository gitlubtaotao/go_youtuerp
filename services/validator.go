package services

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

//错误信息处理返回的Error
type FieldError struct {
	Field       string
	StructField string
	Value       interface{}
	Param       string
	Kind        reflect.Kind
	Type        reflect.Type
	Error       string
	Namespace   string
}

type IValidatorService interface {
	//处理validator 信息
	HandlerError(language string) ([]FieldError, error)
	ResultError(language string) string
}

type ValidatorService struct {
	model interface{}
}

func (v *ValidatorService) ResultError(language string) string {
	handleErr, err := v.HandlerError(language)
	var stringWrite strings.Builder
	if err != nil {
		return err.Error()
	}
	if len(handleErr) <= 0 {
		return ""
	}
	for _, v := range handleErr {
		_, _ = stringWrite.WriteString(v.Error)
	}
	return stringWrite.String()
}

func (v *ValidatorService) HandlerError(language string) (errorsArray []FieldError, err error) {
	fmt.Println(v.model)
	err = validator.New().Struct(v.model)
	if _err := v.registerLanguageService(language); _err != nil {
		return errorsArray, _err
	}
	if _err := v.registerDefaultTranslations(language); _err != nil {
		return errorsArray, _err
	}
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			attr := FieldError{
				Field:       e.Field(),
				StructField: e.StructField(),
				Value:       e.Value(),
				Param:       e.Param(),
				Namespace:   e.Namespace(),
				Error:       e.Translate(trans),
				Kind:        e.Kind(),
				Type:        e.Type(),
			}
			errorsArray = append(errorsArray, attr)
		}
	}
	return
}

func NewValidatorService(model interface{}) IValidatorService {
	return &ValidatorService{model: model}
}

func (v *ValidatorService) registerLanguageService(language string) (err error) {
	var (
		translator locales.Translator
		found      bool
	)
	if language == "zh-CN" {
		language = "zh"
	}
	switch language {
	case "en":
		translator = en.New()
	default:
		translator = zh.New()
	}
	uni = ut.New(translator, translator)
	trans, found = uni.GetTranslator(language)
	fmt.Println(found, language)
	if !found {
		return errors.New("language is not exist")
	}
	return
}

//注册默认的translations
func (v *ValidatorService) registerDefaultTranslations(language string) (err error) {
	validate = validator.New()
	switch language {
	case "en":
		err = en2.RegisterDefaultTranslations(validate, trans)
	case "zh-CN":
		err = zh2.RegisterDefaultTranslations(validate, trans)
	default:
		err = zh2.RegisterDefaultTranslations(validate, trans)
	}
	return
}
