package godd

import (
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
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
func ValidateStruct(i interface{}, iType map[string]interface{}) *Error {
	var errValidateList *map[string]ErrorValidate
	validate := validator.New()
	err := validate.Struct(i)

	if err != nil {
		errValidateList = &map[string]ErrorValidate{}
		for _, err := range err.(validator.ValidationErrors) {

			ss := strings.Split(err.Namespace(), ".")
			length := len(ss) - 2
			if length <= 0 {
				length = 0
			}
			fieldName := getNameField(iType[ss[length]], err.Field())
			v := (*errValidateList)[fieldName]

			if v.ReasonList == nil {
				v.ReasonList = map[string]ErrorValidateReason{}
			}
			v.ReasonList[err.Tag()] = ErrorValidateReason{
				Message: "",
				Param:   err.Param(),
			}

			(*errValidateList)[fieldName] = v
		}
	}

	if errValidateList != nil {
		return &Error{
			ErrorValidate: errValidateList,
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
