package godd

import (
	"net/http"
)

// APIHTTP Struct
type APIHTTP struct {
	handlerByPass     HandlerByPass
	api               *apiLifeCycle
	service           interface{}
	serviceOptionList map[string]interface{}
}

// apiLifeCycle Struct
type apiLifeCycle struct {
	OnStart                 HandlerCycle
	ValidateLanguage        HandlerCycle
	OnPreAuth               HandlerCycle
	ValidateAuth            HandlerCycle
	OnPostAuth              HandlerCycle
	ValidateHeaders         HandlerCycle
	ValidateParams          HandlerCycle
	ValidateQuery           HandlerCycle
	ParseRequest            ParseRequest
	ValidateRequest         ValidateRequest
	OnPreHandler            OnPreHandler
	HandlerLogic            HandlerLogic
	OnPostHandler           OnPostHandler
	MappingResponse         MappingResponse
	ValidateResponse        ValidateResponse
	MappingResponseStandard MappingResponseStandard
	OnPreResponse           OnPreResponse
	SendResponse            SendResponse
}

// NewAPIHTTP API
func NewAPIHTTP() *APIHTTP {
	api := apiLifeCycle{
		OnStart:                 handlerDefault(),
		ValidateLanguage:        handlerDefault(),
		OnPreAuth:               handlerDefault(),
		ValidateAuth:            handlerDefault(),
		OnPostAuth:              handlerDefault(),
		ValidateHeaders:         handlerDefault(),
		ValidateParams:          handlerDefault(),
		ValidateQuery:           handlerDefault(),
		ParseRequest:            handlerParseRequestDefault(),
		ValidateRequest:         handlerValidateRequestDefault(),
		OnPreHandler:            handlerOnPreHandlerDefault(),
		HandlerLogic:            handlerHandlerLogicDefault(),
		OnPostHandler:           handlerOnPostHandlerDefault(),
		MappingResponse:         handlerMappingResponseDefault(),
		ValidateResponse:        handlerValidateResponseDefault(),
		MappingResponseStandard: handlerMappingResponseStandardDefault(),
		OnPreResponse:           handlerOnPreResponseDefault(),
		SendResponse:            handlerSendResponseDefault(),
	}

	return &APIHTTP{api: &api}
}

// HandlerByPass Type
type HandlerByPass = func(service interface{}, serviceOptionList map[string]interface{}) Handler

// HandlerCycle Type
type HandlerCycle = func(context InterfaceContext) (err *Error)

// ParseRequest Type
type ParseRequest = func(context InterfaceContext) (requestMapping interface{}, err *Error)

// ValidateRequest Type
type ValidateRequest = func(context InterfaceContext, requestMapping interface{}) (requestValidated interface{}, err *Error)

// OnPreHandler Type
type OnPreHandler = func(context InterfaceContext, requestValidatedIn interface{}) (requestValidatedOut interface{}, err *Error)

// HandlerLogic Type
type HandlerLogic = func(context InterfaceContext, requestValidated interface{}) (code int, responseRaw interface{}, responsePagination *ResponsePagination, err *Error)

// OnPostHandler Type
type OnPostHandler = func(context InterfaceContext, code int, responseRawIn interface{}, responsePagination *ResponsePagination) (codeOut int, responseRawOut interface{}, responsePaginationOut *ResponsePagination, err *Error)

// MappingResponse Type
type MappingResponse = func(context InterfaceContext, code int, responseRaw interface{}, responsePagination *ResponsePagination) (codeOut int, responseMapping interface{}, responsePaginationOut *ResponsePagination, err *Error)

// ValidateResponse Type
type ValidateResponse = func(context InterfaceContext, code int, responseMapping interface{}, responsePagination *ResponsePagination) (codeOut int, responseValidated interface{}, responsePaginationOut *ResponsePagination, err *Error)

// MappingResponseStandard Type
type MappingResponseStandard = func(context InterfaceContext, code int, responseRaw interface{}, responsePagination *ResponsePagination) (codeOut int, responseMapping interface{}, err *Error)

// OnPreResponse Type
type OnPreResponse = func(context InterfaceContext, code int, requestValidatedIn interface{}) (codeOut int, requestValidatedOut interface{}, err *Error)

// SendResponse Type
type SendResponse = func(context InterfaceContext, code int, requestValidated interface{}) (err *Error)

func handlerDefault() HandlerCycle {
	return func(context InterfaceContext) (err *Error) {
		return nil
	}
}

func handlerParseRequestDefault() ParseRequest {
	return func(context InterfaceContext) (requestMapping interface{}, err *Error) {
		return nil, nil
	}
}

func handlerValidateRequestDefault() ValidateRequest {
	return func(context InterfaceContext, requestMapping interface{}) (requestValidated interface{}, err *Error) {
		return nil, nil
	}
}

func handlerOnPreHandlerDefault() OnPreHandler {
	return func(context InterfaceContext, requestValidatedIn interface{}) (requestValidatedOut interface{}, err *Error) {
		return requestValidatedIn, nil
	}
}

func handlerHandlerLogicDefault() HandlerLogic {
	return func(context InterfaceContext, requestValidated interface{}) (code int, responseRaw interface{}, responsePagination *ResponsePagination, err *Error) {
		return 200, nil, nil, nil
	}
}

func handlerOnPostHandlerDefault() OnPostHandler {
	return func(context InterfaceContext, code int, responseRaw interface{}, responsePaginationIn *ResponsePagination) (codeOut int, responseRawOut interface{}, responsePaginationOut *ResponsePagination, err *Error) {
		return code, responseRaw, responsePaginationIn, nil
	}
}

func handlerMappingResponseDefault() MappingResponse {
	return func(context InterfaceContext, code int, responseRaw interface{}, responsePagination *ResponsePagination) (codeOut int, responseMapping interface{}, responsePaginationOut *ResponsePagination, err *Error) {
		return code, responseRaw, responsePagination, nil
	}
}

func handlerValidateResponseDefault() ValidateResponse {
	return func(context InterfaceContext, code int, responseMapping interface{}, responsePagination *ResponsePagination) (codeOut int, responseValidated interface{}, responsePaginationOut *ResponsePagination, err *Error) {
		return code, responseMapping, responsePagination, nil
	}
}

func handlerMappingResponseStandardDefault() MappingResponseStandard {
	return func(context InterfaceContext, code int, responseRaw interface{}, responsePagination *ResponsePagination) (codeOut int, responseStandard interface{}, err *Error) {

		response, err := MappingStandard(code, responseRaw, responsePagination)
		return code, response, err
	}
}

func handlerOnPreResponseDefault() OnPreResponse {
	return func(context InterfaceContext, code int, responseStandard interface{}) (codeOut int, responseStandardOut interface{}, err *Error) {
		return code, responseStandard, nil
	}
}

func handlerSendResponseDefault() SendResponse {
	return func(context InterfaceContext, code int, responseStandard interface{}) (err *Error) {
		if responseStandard != nil {
			context.Response(responseStandard, code)
		}
		return nil
	}
}

// ================================================================

// SetHandlerByPassLifeCycle API
func (api *APIHTTP) SetHandlerByPassLifeCycle(handler HandlerByPass) {
	api.handlerByPass = handler
}

// SetupHandlerHTTP API
func (api *APIHTTP) SetupHandlerHTTP(service interface{}, serviceOptionList map[string]interface{}) Handler {
	api.service = service
	api.serviceOptionList = serviceOptionList
	return api.handlerLifeCycle()
}

// ================================================================

func (api *APIHTTP) handlerLifeCycle() Handler {
	if api.handlerByPass != nil {
		return api.handlerByPass(api.service, api.serviceOptionList)
	}

	return func(context InterfaceContext) error {
		// ===============================================
		context.SetContext(api, map[string]interface{}{})
		var err *Error

		// ================== Start =======================

		err = api.api.OnStart(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.ValidateLanguage(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Auth =======================

		err = api.api.OnPreAuth(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.ValidateAuth(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.OnPostAuth(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Validate Request =======================

		err = api.api.ValidateHeaders(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.ValidateParams(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.ValidateQuery(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		requestMapping, err := api.api.ParseRequest(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		requestValidated, err := api.api.ValidateRequest(context, requestMapping)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Handler =======================

		requestValidated, err = api.api.OnPreHandler(context, requestValidated)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseRaw, responsePagination, err := api.api.HandlerLogic(context, requestValidated)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseRaw, responsePagination, err = api.api.OnPostHandler(context, code, responseRaw, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Validate Response =======================

		code, responseMapping, responsePagination, err := api.api.MappingResponse(context, code, responseRaw, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseValidated, responsePagination, err := api.api.ValidateResponse(context, code, responseMapping, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseStandard, err := api.api.MappingResponseStandard(context, code, responseValidated, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseStandard, err = api.api.OnPreResponse(context, code, responseStandard)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.SendResponse(context, code, responseStandard)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		return nil
	}
}

//========================================

// OnPreAuth func
func (api *APIHTTP) OnPreAuth(handler HandlerCycle) {
	api.api.OnPreAuth = handler
}

// ValidateAuth func
func (api *APIHTTP) ValidateAuth(handler HandlerCycle) {
	api.api.ValidateAuth = handler
}

// ValidateHeaders func
func (api *APIHTTP) ValidateHeaders(handler HandlerCycle) {
	api.api.ValidateHeaders = handler
}

// ValidateParams func
func (api *APIHTTP) ValidateParams(handler HandlerCycle) {
	api.api.ValidateParams = handler
}

// ValidateQuery func
func (api *APIHTTP) ValidateQuery(handler HandlerCycle) {
	api.api.ValidateQuery = handler
}

// ParseRequest func
func (api *APIHTTP) ParseRequest(handler ParseRequest) {
	api.api.ParseRequest = handler
}

// ValidateRequest func
func (api *APIHTTP) ValidateRequest(handler ValidateRequest) {
	api.api.ValidateRequest = handler
}

// OnPreHandler func
func (api *APIHTTP) OnPreHandler(handler OnPreHandler) {
	api.api.OnPreHandler = handler
}

// HandlerLogic func
func (api *APIHTTP) HandlerLogic(handler HandlerLogic) {
	api.api.HandlerLogic = handler
}

// OnPostHandler func
func (api *APIHTTP) OnPostHandler(handler OnPostHandler) {
	api.api.OnPostHandler = handler
}

// MappingResponse func
func (api *APIHTTP) MappingResponse(handler MappingResponse) {
	api.api.MappingResponse = handler
}

// ValidateResponse func
func (api *APIHTTP) ValidateResponse(handler ValidateResponse) {
	api.api.ValidateResponse = handler
}

// OnPreResponse func
func (api *APIHTTP) OnPreResponse(handler OnPreResponse) {
	api.api.OnPreResponse = handler
}

// SendResponse func
func (api *APIHTTP) SendResponse(handler SendResponse) {
	api.api.SendResponse = handler
}

//========================================

func encodeErrorHandler(context InterfaceContext, err *Error) error {
	var errorMessage string
	if err.Error != nil {
		errorMessage = err.Error.Error()
	} else {
		errorMessage = http.StatusText(err.Code)
	}
	// errorreporting.LogError(errors.New(errorMessage))
	context.SetContentType("application/json; charset=utf-8")

	context.Response(
		ResponseDataList{
			Success: false,
			Message: "unsuccess",
			ResponseError: &ResponseError{
				Message: errorMessage,
			},
		},
		err.Code,
	)
	return nil
}

// MappingStandard Func
func MappingStandard(code int, dataList interface{}, responsePagination *ResponsePagination) (response ResponseDataList, err *Error) {
	if dataList == nil {
		dataList = []string{}
	}
	isSuccess := code/100 == 2
	var message string
	if isSuccess {
		message = "success"
	} else {
		message = "unsuccess"
	}
	response = ResponseDataList{
		Success:            isSuccess,
		Message:            message,
		Data:               dataList,
		ResponsePagination: responsePagination,
	}
	return response, nil
}
