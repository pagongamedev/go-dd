package http

import (
	"log"
	"reflect"

	godd "github.com/pagongamedev/go-dd"
)

//===============================================================================================================================================================

// MappingStandardError Func
func MappingStandardError(goddErr *godd.Error) (codeOut int, responseError interface{}, goddErrOut *godd.Error) {
	// errorreporting.LogError(errors.New(errorMessage))
	return goddErr.Code, &godd.ResponseDataList{
		Code:    goddErr.Code,
		Success: false,
		Message: "unsuccess",
		ResponseError: &godd.ResponseError{
			Message:  goddErr.GetMessage(),
			Validate: goddErr.GetErrorValidate(),
		},
	}, nil
}

// MappingStandard Func
func MappingStandardResponse(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseStandard interface{}, goddErr *godd.Error) {
	if responseRaw == nil {
		responseRaw = []string{}
	}

	if reflect.ValueOf(responseRaw).Kind() != reflect.Slice {
		log.Println("Warning : Input Data Response Not Array")
		responseRaw = append(([]interface{}{}), responseRaw)
	}

	isSuccess := code/100 == 2
	var message string
	if isSuccess {
		message = "success"
	} else {
		message = "unsuccess"
	}
	response := godd.ResponseDataList{
		Code:               code,
		Success:            isSuccess,
		Message:            message,
		Data:               responseRaw,
		ResponsePagination: responsePagination,
	}
	return code, response, nil
}
