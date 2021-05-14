package api

import (
	"log"
	"reflect"

	godd "github.com/pagongamedev/go-dd"
)

//========================================

// HTTPMappingStandardError
func HTTPMappingStandardError(goddErr *godd.Error) (codeOut int, responseError interface{}, goddErrOut *godd.Error) {
	return codeOut, godd.ResponseDataList{
			Success: false,
			Message: "unsuccess",
			ResponseError: &godd.ResponseError{
				Message:  goddErr.GetMessage(),
				Validate: goddErr.GetErrorValidate(),
			},
		},
		nil
}

// HTTPMappingStandardResponse
func HTTPMappingStandardResponse(context *godd.Context, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseMapping interface{}, goddErr *godd.Error) {
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
