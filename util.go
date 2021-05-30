package godd

import (
	"log"
	"math"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	defaults "github.com/pagongamedev/temp-go-defaults"
	"golang.org/x/text/language"
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

// =========================================

// ValidateStruct func
func ValidateStruct(i18n *I18N, i interface{}, iType map[string]interface{}) *Error {
	var errList *map[string]ResponseErrorValidate
	validate := validator.New()
	err := validate.Struct(i)

	if err != nil {
		errList = &map[string]ResponseErrorValidate{}
		for _, err := range err.(validator.ValidationErrors) {

			ss := strings.Split(err.Namespace(), ".")
			length := len(ss) - 2
			if length <= 0 {
				length = 0
			}
			fieldName := getNameField(iType[ss[length]], err.Field())
			(*errList)[fieldName] = ResponseErrorValidate{
				Reason:  err.Tag(),
				Param:   err.Param(),
				Message: i18n.MustLocalize("validate_"+err.Tag(), Map{"Field": fieldName, "Param": err.Param()}, 0),
			}

		}
	}

	if errList != nil {
		return &Error{
			Code:          http.StatusBadRequest,
			errorValidate: errList,
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

func GetOneFromArray(list ...interface{}) interface{} {
	if len(list) > 0 {
		return list[0]
	}
	return nil
}

// =========================================

// I18N func
type I18N struct {
	bundle        *i18n.Bundle
	localizerList map[string]*i18n.Localizer
	lang          string
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
func (i *I18N) MustLocalize(id string, data Map, count int, m ...interface{}) string {
	var iMessage *i18n.Message
	if len(m) > 0 {
		iMessage = m[0].(*i18n.Message)
	} else {
		iMessage = &i18n.Message{
			ID: id,
		}
	}
	l := i.lang
	langList := strings.Split(i.lang, ",")
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

// SetLang func
func (i *I18N) SetLang(lang string) {
	i.lang = lang
}

// GetLang func
func (i *I18N) GetLang() string {
	return i.lang
}

//====================

// EnvironmentSwitcher func
func EnvironmentSwitcher(env string, Localhost int, Development int, Testing int, Staging int, Production int, i ...interface{}) interface{} {
	var index int
	switch Env(env) {
	case EnvLocalhost:
		index = Localhost
	case EnvDevelopment:
		index = Development
	case EnvTesting:
		index = Testing
	case EnvStaging:
		index = Staging
	case EnvProduction:
		index = Production
	}

	if index < 0 || index > len(i)-1 {
		log.Println("Out of Length")
		return nil
	}

	return i[index]
}

// IsInterfaceIsNil func
func IsInterfaceIsNil(i interface{}) bool {
	if i == nil {
		return true
	}

	switch reflect.TypeOf(i).Kind() {
	// case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice, reflect.Func:
	// 	return reflect.ValueOf(i).IsNil()

	case reflect.Struct:
		return false
	}

	return reflect.ValueOf(i).IsNil()

}

// ConvertToArray func
func ConvertToArray(dataList interface{}) []interface{} {
	return append(([]interface{}{}), dataList)
}

// Helper Error
func ErrorNew(code int, err error) *Error {
	return &Error{
		Code:  code,
		Error: err,
	}
}

//====================

// Pagination

func GetResponsePagination(requestPagination RequestPagination, totalCount int, itemCount int) *ResponsePagination {
	var pageCount int

	if totalCount != 0 {
		pageCount = (int)(math.Ceil((float64)(totalCount) / (float64)(requestPagination.PageSize)))
	}

	responsePagination := &ResponsePagination{
		Page:       requestPagination.Page,
		PageSize:   requestPagination.PageSize,
		PageCount:  pageCount,
		ItemCount:  itemCount,
		TotalCount: totalCount,
	}

	return responsePagination
}

//====================
