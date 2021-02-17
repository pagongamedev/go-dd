package godd

import (
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/mcuadros/go-defaults.v1"
)

// =====================================================================
//                              Add On
// =====================================================================

// MustError Func
func MustError(err error, strList ...string) {
	if err != nil {
		if strList != nil {
			log.Fatal(strList)
		} else {
			log.Fatal("Error : ", err)
		}
	}
}

// AddAPIGetHealth Func
func addAPIGetHealth(app InterfaceApp) {
	app.Get("/health", handlerHealth())
}

func handlerHealth() Handler {
	return func(ctx InterfaceContext) error {
		return ctx.Response(Map{"success": true})
	}
}

// =========================================

// ValidateStruct func
func ValidateStruct(context InterfaceContext, i interface{}, iType map[string]interface{}) *Error {
	var errList *map[string]ErrorValidate
	validate := validator.New()
	err := validate.Struct(i)

	if err != nil {
		errList = &map[string]ErrorValidate{}
		for _, err := range err.(validator.ValidationErrors) {

			ss := strings.Split(err.Namespace(), ".")
			length := len(ss) - 2
			if length <= 0 {
				length = 0
			}
			fieldName := getNameField(iType[ss[length]], err.Field())
			(*errList)[fieldName] = ErrorValidate{
				Reason:  err.Tag(),
				Param:   err.Param(),
				Message: context.MustLocalize("validate_"+err.Tag(), Map{"Field": fieldName, "Param": err.Param()}, 0),
			}

		}
	}

	if errList != nil {
		return &Error{
			ErrorValidate: errList,
		}
	}
	return nil
}

// getNameField func
func getNameField(i interface{}, str string) string {
	field, _ := reflect.TypeOf(i).Elem().FieldByName(str)
	s := field.Tag.Get("json")
	if s != "" {
		s = strings.ReplaceAll(s, "omitempty", "")
		s = strings.ReplaceAll(s, ",", "")
		s = strings.ReplaceAll(s, " ", "")
		return s
	}
	return str
}

// =========================================

// SetDefaultStruct func
func SetDefaultStruct(variable interface{}) interface{} {
	defaults.SetDefaults(variable) //<-- This set the defaults values
	return variable
}

// =========================================

// I18N func
type I18N struct {
	bundle        *i18n.Bundle
	localizerList map[string]*i18n.Localizer
}

// NewI18N func
func NewI18N(defaultLanguage language.Tag, formatUnmarshal string, unmarshalFunc i18n.UnmarshalFunc, fileLanguageList Map) *I18N {
	bundle := i18n.NewBundle(defaultLanguage)
	bundle.RegisterUnmarshalFunc(formatUnmarshal, unmarshalFunc)

	localizerList := map[string]*i18n.Localizer{}

	for key, fileLanguage := range fileLanguageList {
		var k string
		keyList := strings.Split(key, "-")
		k = key
		if len(keyList) > 0 {
			k = keyList[0]
		}
		bundle.LoadMessageFile(fileLanguage.(string))
		localizerList[k] = i18n.NewLocalizer(bundle, key)
	}

	return &I18N{
		bundle:        bundle,
		localizerList: localizerList,
	}
}

//MustLocalize func
func (i *I18N) MustLocalize(lang string, id string, data Map, count int, m ...interface{}) string {
	var iMessage *i18n.Message
	if len(m) > 0 {
		iMessage = m[0].(*i18n.Message)
	} else {
		iMessage = &i18n.Message{
			ID: id,
		}
	}
	l := lang
	langList := strings.Split(lang, ",")
	if len(langList) > 0 {
		l = langList[0]
		langList = strings.Split(l, "-")
		if len(langList) > 0 {
			l = langList[0]
		}
	}

	localizer := (i.localizerList[l])
	if localizer != nil {
		return localizer.MustLocalize(
			&i18n.LocalizeConfig{
				DefaultMessage: iMessage,
				TemplateData:   data,
				PluralCount:    count,
			})
	}
	return ""
}
