package api

import (
	"net/http"

	godd "github.com/pagongamedev/go-dd"
)

// HTTP struct
type HTTP struct {
	HandlerByPass     HandlerByPass
	API               *apiLifeCycle
	Service           interface{}
	ServiceOptionList map[string]interface{}
	I18n              *godd.I18N
}

// apiLifeCycle Struct
type apiLifeCycle struct {
	OnStart                 HandlerCycle
	ParseLanguage           HandlerCycle
	OnPreAuth               HandlerCycle
	ValidateAuth            HandlerCycle
	OnPostAuth              HandlerCycle
	ValidateHeader          HandlerCycle
	ValidateParam           ValidateParam
	ValidateQuery           ValidateQuery
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
func NewAPIHTTP() *HTTP {
	api := apiLifeCycle{
		OnStart:                 handlerDefault(),
		ParseLanguage:           handlerParseLanguage(),
		OnPreAuth:               handlerDefault(),
		ValidateAuth:            handlerDefault(),
		OnPostAuth:              handlerDefault(),
		ValidateHeader:          handlerDefault(),
		ValidateParam:           handlerValidateParam(),
		ValidateQuery:           handlerValidateQuery(),
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

	return &HTTP{API: &api}
}

// HandlerByPass Type
type HandlerByPass = func(service interface{}, serviceOptionList map[string]interface{}) godd.Handler

// HandlerCycle Type
type HandlerCycle = func(context godd.InterfaceContext) (err *godd.Error)

// ValidateParam Type
type ValidateParam = func(context godd.InterfaceContext) (requestValidatedParam interface{}, err *godd.Error)

// ValidateQuery Type
type ValidateQuery = func(context godd.InterfaceContext) (requestValidatedQuery interface{}, err *godd.Error)

// ParseRequest Type
type ParseRequest = func(context godd.InterfaceContext) (requestMappingBody interface{}, err *godd.Error)

// ValidateRequest Type
type ValidateRequest = func(context godd.InterfaceContext, requestMappingBody interface{}) (requestValidatedBody interface{}, err *godd.Error)

// OnPreHandler Type
type OnPreHandler = func(context godd.InterfaceContext, requestValidatedBody interface{}, requestValidatedParam interface{}, requestValidatedQuery interface{}) (requestValidatedBodyOut interface{}, requestValidatedParamOut interface{}, requestValidatedQueryOut interface{}, err *godd.Error)

// HandlerLogic Type
type HandlerLogic = func(context godd.InterfaceContext, requestValidatedBody interface{}, requestValidatedParam interface{}, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, err *godd.Error)

// OnPostHandler Type
type OnPostHandler = func(context godd.InterfaceContext, code int, responseRawIn interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseRawOut interface{}, responsePaginationOut *godd.ResponsePagination, err *godd.Error)

// MappingResponse Type
type MappingResponse = func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseMapping interface{}, responsePaginationOut *godd.ResponsePagination, err *godd.Error)

// ValidateResponse Type
type ValidateResponse = func(context godd.InterfaceContext, code int, responseMapping interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseValidated interface{}, responsePaginationOut *godd.ResponsePagination, err *godd.Error)

// MappingResponseStandard Type
type MappingResponseStandard = func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseMapping interface{}, err *godd.Error)

// OnPreResponse Type
type OnPreResponse = func(context godd.InterfaceContext, code int, requestValidatedIn interface{}) (codeOut int, requestValidatedOut interface{}, err *godd.Error)

// SendResponse Type
type SendResponse = func(context godd.InterfaceContext, code int, requestValidated interface{}) (err *godd.Error)

func handlerDefault() HandlerCycle {
	return func(context godd.InterfaceContext) (err *godd.Error) {
		return nil
	}
}

func handlerValidateParam() ValidateParam {
	return func(context godd.InterfaceContext) (requestValidatedParam interface{}, err *godd.Error) {
		return nil, nil
	}
}
func handlerValidateQuery() ValidateQuery {
	return func(context godd.InterfaceContext) (requestValidatedQuery interface{}, err *godd.Error) {
		return nil, nil
	}
}

func handlerParseLanguage() HandlerCycle {
	return func(context godd.InterfaceContext) (err *godd.Error) {

		acceptLanguage := context.GetHeader("Accept-Language")
		if acceptLanguage == "" {
			acceptLanguage = "en-US"
		}

		context.SetLang(acceptLanguage)

		return nil
	}
}

func handlerParseRequestDefault() ParseRequest {
	return func(context godd.InterfaceContext) (requestMapping interface{}, err *godd.Error) {
		return nil, nil
	}
}

func handlerValidateRequestDefault() ValidateRequest {
	return func(context godd.InterfaceContext, requestMapping interface{}) (requestValidated interface{}, err *godd.Error) {
		return nil, nil
	}
}

func handlerOnPreHandlerDefault() OnPreHandler {
	return func(context godd.InterfaceContext, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (requestValidatedBodyOut interface{}, requestValidatedParamOut interface{}, requestValidatedQueryOut interface{}, err *godd.Error) {
		return requestValidatedBody, requestValidatedParam, requestValidatedQuery, nil
	}
}

func handlerHandlerLogicDefault() HandlerLogic {
	return func(context godd.InterfaceContext, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, err *godd.Error) {
		return 200, nil, nil, nil
	}
}

func handlerOnPostHandlerDefault() OnPostHandler {
	return func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePaginationIn *godd.ResponsePagination) (codeOut int, responseRawOut interface{}, responsePaginationOut *godd.ResponsePagination, err *godd.Error) {
		return code, responseRaw, responsePaginationIn, nil
	}
}

func handlerMappingResponseDefault() MappingResponse {
	return func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseMapping interface{}, responsePaginationOut *godd.ResponsePagination, err *godd.Error) {
		return code, responseRaw, responsePagination, nil
	}
}

func handlerValidateResponseDefault() ValidateResponse {
	return func(context godd.InterfaceContext, code int, responseMapping interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseValidated interface{}, responsePaginationOut *godd.ResponsePagination, err *godd.Error) {
		return code, responseMapping, responsePagination, nil
	}
}

func handlerMappingResponseStandardDefault() MappingResponseStandard {
	return func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseStandard interface{}, err *godd.Error) {

		response, err := MappingStandard(code, responseRaw, responsePagination)
		return code, response, err
	}
}

func handlerOnPreResponseDefault() OnPreResponse {
	return func(context godd.InterfaceContext, code int, responseStandard interface{}) (codeOut int, responseStandardOut interface{}, err *godd.Error) {
		return code, responseStandard, nil
	}
}

func handlerSendResponseDefault() SendResponse {
	return func(context godd.InterfaceContext, code int, responseStandard interface{}) (err *godd.Error) {
		if responseStandard != nil {
			context.Response(responseStandard, code)
		}
		return nil
	}
}

// ================================================================

// SetHandlerByPassLifeCycle API
func (api *HTTP) SetHandlerByPassLifeCycle(handler HandlerByPass) {
	api.HandlerByPass = handler
}

// SetupHandlerHTTP API
func (api *HTTP) SetupHandlerHTTP(service interface{}, serviceOptionList map[string]interface{}, i18n *godd.I18N) godd.Handler {
	api.Service = service
	api.ServiceOptionList = serviceOptionList
	api.I18n = i18n
	return api.HandlerLifeCycle()
}

// ================================================================

// HandlerLifeCycle func
func (api *HTTP) HandlerLifeCycle() godd.Handler {
	if api.HandlerByPass != nil {
		return api.HandlerByPass(api.Service, api.ServiceOptionList)
	}

	return func(context godd.InterfaceContext) error {
		// ===============================================
		context.SetContext(api.Service, api.ServiceOptionList, api.I18n, map[string]interface{}{})
		var err *godd.Error

		// ================== Start =======================

		err = api.API.OnStart(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.API.ParseLanguage(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Auth =======================

		err = api.API.OnPreAuth(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.API.ValidateAuth(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.API.OnPostAuth(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Validate Request =======================

		err = api.API.ValidateHeader(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		requestValidatedParam, err := api.API.ValidateParam(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		requestValidatedQuery, err := api.API.ValidateQuery(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		requestMappingBody, err := api.API.ParseRequest(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		requestValidatedBody, err := api.API.ValidateRequest(context, requestMappingBody)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Handler =======================

		requestValidatedBody, requestValidatedParam, requestValidatedQuery, err = api.API.OnPreHandler(context, requestValidatedBody, requestValidatedParam, requestValidatedQuery)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseRaw, responsePagination, err := api.API.HandlerLogic(context, requestValidatedBody, requestValidatedParam, requestValidatedQuery)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseRaw, responsePagination, err = api.API.OnPostHandler(context, code, responseRaw, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Validate Response =======================

		code, responseMapping, responsePagination, err := api.API.MappingResponse(context, code, responseRaw, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseValidated, responsePagination, err := api.API.ValidateResponse(context, code, responseMapping, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseStandard, err := api.API.MappingResponseStandard(context, code, responseValidated, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseStandard, err = api.API.OnPreResponse(context, code, responseStandard)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.API.SendResponse(context, code, responseStandard)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		return nil
	}
}

//========================================

// OnStart func
func (api *HTTP) OnStart(handler HandlerCycle) {
	api.API.OnStart = handler
}

// ParseLanguage func
func (api *HTTP) ParseLanguage(handler HandlerCycle) {
	api.API.ParseLanguage = handler
}

// OnPreAuth func
func (api *HTTP) OnPreAuth(handler HandlerCycle) {
	api.API.OnPreAuth = handler
}

// ValidateAuth func
func (api *HTTP) ValidateAuth(handler HandlerCycle) {
	api.API.ValidateAuth = handler
}

// ValidateHeader func
func (api *HTTP) ValidateHeader(handler HandlerCycle) {
	api.API.ValidateHeader = handler
}

// ValidateParam func
func (api *HTTP) ValidateParam(handler ValidateParam) {
	api.API.ValidateParam = handler
}

// ValidateQuery func
func (api *HTTP) ValidateQuery(handler ValidateQuery) {
	api.API.ValidateQuery = handler
}

// ParseRequest func
func (api *HTTP) ParseRequest(handler ParseRequest) {
	api.API.ParseRequest = handler
}

// ValidateRequest func
func (api *HTTP) ValidateRequest(handler ValidateRequest) {
	api.API.ValidateRequest = handler
}

// OnPreHandler func
func (api *HTTP) OnPreHandler(handler OnPreHandler) {
	api.API.OnPreHandler = handler
}

// HandlerLogic func
func (api *HTTP) HandlerLogic(handler HandlerLogic) {
	api.API.HandlerLogic = handler
}

// OnPostHandler func
func (api *HTTP) OnPostHandler(handler OnPostHandler) {
	api.API.OnPostHandler = handler
}

// MappingResponse func
func (api *HTTP) MappingResponse(handler MappingResponse) {
	api.API.MappingResponse = handler
}

// ValidateResponse func
func (api *HTTP) ValidateResponse(handler ValidateResponse) {
	api.API.ValidateResponse = handler
}

// OnPreResponse func
func (api *HTTP) OnPreResponse(handler OnPreResponse) {
	api.API.OnPreResponse = handler
}

// SendResponse func
func (api *HTTP) SendResponse(handler SendResponse) {
	api.API.SendResponse = handler
}

//========================================

// GetOnStart func
func (api *HTTP) GetOnStart() HandlerCycle {
	return api.API.OnStart
}

// GetParseLanguage func
func (api *HTTP) GetParseLanguage() HandlerCycle {
	return api.API.ParseLanguage
}

// GetOnPreAuth func
func (api *HTTP) GetOnPreAuth() HandlerCycle {
	return api.API.OnPreAuth
}

// GetValidateAuth func
func (api *HTTP) GetValidateAuth() HandlerCycle {
	return api.API.ValidateAuth
}

// GetOnPostAuth func
func (api *HTTP) GetOnPostAuth() HandlerCycle {
	return api.API.OnPostAuth
}

// GetValidateHeader func
func (api *HTTP) GetValidateHeader() HandlerCycle {
	return api.API.ValidateHeader
}

// GetValidateParam func
func (api *HTTP) GetValidateParam() ValidateParam {
	return api.API.ValidateParam
}

// GetValidateQuery func
func (api *HTTP) GetValidateQuery() ValidateQuery {
	return api.API.ValidateQuery
}

// GetParseRequest func
func (api *HTTP) GetParseRequest() ParseRequest {
	return api.API.ParseRequest
}

// GetValidateRequest func
func (api *HTTP) GetValidateRequest() ValidateRequest {
	return api.API.ValidateRequest
}

// GetOnPreHandler func
func (api *HTTP) GetOnPreHandler() OnPreHandler {
	return api.API.OnPreHandler
}

// GetHandlerLogic func
func (api *HTTP) GetHandlerLogic() HandlerLogic {
	return api.API.HandlerLogic
}

// GetOnPostHandler func
func (api *HTTP) GetOnPostHandler() OnPostHandler {
	return api.API.OnPostHandler
}

// GetMappingResponse func
func (api *HTTP) GetMappingResponse() MappingResponse {
	return api.API.MappingResponse
}

// GetValidateResponse func
func (api *HTTP) GetValidateResponse() ValidateResponse {
	return api.API.ValidateResponse
}

// GetMappingResponseStandard func
func (api *HTTP) GetMappingResponseStandard() MappingResponseStandard {
	return api.API.MappingResponseStandard
}

// GetOnPreResponse func
func (api *HTTP) GetOnPreResponse() OnPreResponse {
	return api.API.OnPreResponse
}

// GetSendResponse func
func (api *HTTP) GetSendResponse() SendResponse {
	return api.API.SendResponse
}

//========================================

func encodeErrorHandler(context godd.InterfaceContext, err *godd.Error) error {
	if err.ErrorValidate != nil {
		err.Code = http.StatusBadRequest
	}

	var errorMessage string
	if err.Error != nil {
		errorMessage = err.Error.Error()
	} else {
		errorMessage = http.StatusText(err.Code)
	}
	// errorreporting.LogError(errors.New(errorMessage))
	context.SetContentType("application/json; charset=utf-8")

	context.Response(
		godd.ResponseDataList{
			Success: false,
			Message: "unsuccess",
			ResponseError: &godd.ResponseError{
				Message:  errorMessage,
				Validate: err.ErrorValidate,
			},
		},
		err.Code,
	)
	return nil
}

// MappingStandard Func
func MappingStandard(code int, dataList interface{}, responsePagination *godd.ResponsePagination) (response godd.ResponseDataList, err *godd.Error) {
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
	response = godd.ResponseDataList{
		Success:            isSuccess,
		Message:            message,
		Data:               dataList,
		ResponsePagination: responsePagination,
	}
	return response, nil
}
