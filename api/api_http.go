package api

import (
	"log"
	"reflect"

	godd "github.com/pagongamedev/go-dd"
	mdw "github.com/pagongamedev/go-dd/middleware"
)

// HTTP struct
type HTTP struct {
	handlerByPass     godd.HandlerByPass
	LifeCycle         *godd.APILifeCycle
	service           interface{}
	serviceOptionList map[string]interface{}
	i18n              *godd.I18N
	middleware        *mdw.Middleware
}

// NewAPIHTTP API
func NewAPIHTTP() *HTTP {

	return &HTTP{LifeCycle: &godd.APILifeCycle{}}
}

func handlerDefault() godd.HandlerCycle {
	return func(context godd.InterfaceContext) (goddgoddErr *godd.Error) {
		return nil
	}
}

func handlerValidateAuth() godd.ValidateAuth {
	return func(context godd.InterfaceContext) (roleData interface{}, goddgoddErr *godd.Error) {
		return nil, nil
	}
}

func handlerValidateRole() godd.ValidateRole {
	return func(context godd.InterfaceContext, roleData interface{}) (goddgoddErr *godd.Error) {
		return nil
	}
}

func handlerValidateParam() godd.ValidateParam {
	return func(context godd.InterfaceContext) (requestValidatedParam interface{}, goddgoddErr *godd.Error) {
		return nil, nil
	}
}
func handlerValidateQuery() godd.ValidateQuery {
	return func(context godd.InterfaceContext) (requestValidatedQuery interface{}, goddgoddErr *godd.Error) {
		return nil, nil
	}
}

func handlerParseLanguage() godd.HandlerCycle {
	return func(context godd.InterfaceContext) (goddgoddErr *godd.Error) {

		if context != nil {
			acceptLanguage := context.GetHeader("Accept-Language")
			if acceptLanguage == "" {
				acceptLanguage = "en-US"
			}

			context.SetLang(acceptLanguage)
		}
		return nil
	}
}

func handlerParseRequestDefault() godd.ParseRequest {
	return func(context godd.InterfaceContext) (requestMapping interface{}, goddgoddErr *godd.Error) {
		return nil, nil
	}
}

func handlerValidateRequestDefault() godd.ValidateRequest {
	return func(context godd.InterfaceContext, requestMapping interface{}) (requestValidated interface{}, goddgoddErr *godd.Error) {
		return requestMapping, nil
	}
}

func handlerOnPreHandlerDefault() godd.OnPreHandler {
	return func(context godd.InterfaceContext, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (requestValidatedBodyOut interface{}, requestValidatedParamOut interface{}, requestValidatedQueryOut interface{}, goddgoddErr *godd.Error) {
		return requestValidatedBody, requestValidatedParam, requestValidatedQuery, nil
	}
}

func handlerHandlerLogicDefault() godd.HandlerLogic {
	return func(context godd.InterfaceContext, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, goddgoddErr *godd.Error) {
		return 200, nil, nil, nil
	}
}

func handlerOnPostHandlerDefault() godd.OnPostHandler {
	return func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePaginationIn *godd.ResponsePagination) (codeOut int, responseRawOut interface{}, responsePaginationOut *godd.ResponsePagination, goddgoddErr *godd.Error) {
		return code, responseRaw, responsePaginationIn, nil
	}
}

func handlerMappingResponseDefault() godd.MappingResponse {
	return func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseMapping interface{}, responsePaginationOut *godd.ResponsePagination, goddgoddErr *godd.Error) {
		return code, responseRaw, responsePagination, nil
	}
}

func handlerValidateResponseDefault() godd.ValidateResponse {
	return func(context godd.InterfaceContext, code int, responseMapping interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseValidated interface{}, responsePaginationOut *godd.ResponsePagination, goddgoddErr *godd.Error) {
		return code, responseMapping, responsePagination, nil
	}
}

func handlerMappingResponseStandardDefault() godd.MappingResponseStandard {
	return func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseStandard interface{}, goddgoddErr *godd.Error) {

		response, goddErr := MappingStandard(code, responseRaw, responsePagination)
		return code, response, goddErr
	}
}

func handlerOnPreResponseDefault() godd.OnPreResponse {
	return func(context godd.InterfaceContext, code int, responseStandard interface{}) (codeOut int, responseStandardOut interface{}, goddgoddErr *godd.Error) {
		return code, responseStandard, nil
	}
}

func handlerSendResponseDefault() godd.SendResponse {
	return func(context godd.InterfaceContext, code int, responseStandard interface{}) (goddgoddErr *godd.Error) {
		if !godd.IsInterfaceIsNil(responseStandard) {
			if context != nil {
				context.Response(responseStandard, code)
			}
		}
		return nil
	}
}

// ================================================================

// SetHandlerByPassLifeCycle API
func (api *HTTP) SetHandlerByPassLifeCycle(handler godd.HandlerByPass) {
	api.handlerByPass = handler
}

// SetupHandlerHTTP API
func (api *HTTP) SetupHandlerHTTP(service interface{}, serviceOptionList map[string]interface{}, i18n *godd.I18N, middleware *mdw.Middleware) {
	api.service = service
	api.serviceOptionList = serviceOptionList
	api.i18n = i18n
	api.middleware = middleware
	api.middlewareLifeCycleChecker()
}

// ================================================================

// HandlerLifeCycle func
func (api *HTTP) HandlerLifeCycle() godd.Handler {
	if api.handlerByPass != nil {
		return api.handlerByPass(api.service, api.serviceOptionList)
	}

	return func(context godd.InterfaceContext) error {
		var goddErr *godd.Error
		context.SetContext(api.service, api.serviceOptionList, api.i18n, map[string]interface{}{})
		// ================== Mdw =======================

		if api.middleware.HandlerStartList != nil {
			for _, funcMdw := range api.middleware.HandlerStartList {
				goddErr = funcMdw(context)
				if goddErr != nil {
					return encodeErrorHandler(context, goddErr)
				}
			}
		}
		// ================== Start =======================
		code, responseStandard, goddErr := api.LifeCycle.RunLifeCycle(context)
		if goddErr != nil {
			return encodeErrorHandler(context, goddErr)
		}

		goddErr = api.LifeCycle.GetSendResponse()(context, code, responseStandard)
		if goddErr != nil {
			return encodeErrorHandler(context, goddErr)
		}
		// ================== Mdw =======================

		if api.middleware.HandlerEndList != nil {
			for _, funcMdw := range api.middleware.HandlerEndList {
				goddErr = funcMdw(context)
				if goddErr != nil {
					return encodeErrorHandler(context, goddErr)
				}
			}
		}
		// ================== Start =======================
		return nil
	}
}

func (api *HTTP) middlewareLifeCycleChecker() {
	var i interface{}

	i = handlerlifeCycleChecker("Onstart", api.LifeCycle.GetOnStart(), api.middleware.LifeCycle.GetOnStart(), handlerDefault())
	api.LifeCycle.OnStart(i.(godd.HandlerCycle))

	i = handlerlifeCycleChecker("ParseLanguage", api.LifeCycle.GetParseLanguage(), api.middleware.LifeCycle.GetParseLanguage(), handlerParseLanguage())
	api.LifeCycle.ParseLanguage(i.(godd.HandlerCycle))

	i = handlerlifeCycleChecker("OnPreAuth", api.LifeCycle.GetOnPreAuth(), api.middleware.LifeCycle.GetOnPreAuth(), handlerDefault())
	api.LifeCycle.OnPreAuth(i.(godd.HandlerCycle))

	i = handlerlifeCycleChecker("ValidateAuth", api.LifeCycle.GetValidateAuth(), api.middleware.LifeCycle.GetValidateAuth(), handlerValidateAuth())
	api.LifeCycle.ValidateAuth(i.(godd.ValidateAuth))

	i = handlerlifeCycleChecker("ValidateRole", api.LifeCycle.GetValidateRole(), api.middleware.LifeCycle.GetValidateRole(), handlerValidateRole())
	api.LifeCycle.ValidateRole(i.(godd.ValidateRole))

	i = handlerlifeCycleChecker("OnPostAuth", api.LifeCycle.GetOnPostAuth(), api.middleware.LifeCycle.GetOnPostAuth(), handlerDefault())
	api.LifeCycle.OnPostAuth(i.(godd.HandlerCycle))

	i = handlerlifeCycleChecker("ValidateHeader", api.LifeCycle.GetValidateHeader(), api.middleware.LifeCycle.GetValidateHeader(), handlerDefault())
	api.LifeCycle.ValidateHeader(i.(godd.HandlerCycle))

	i = handlerlifeCycleChecker("ValidateParam", api.LifeCycle.GetValidateParam(), api.middleware.LifeCycle.GetValidateParam(), handlerValidateParam())
	api.LifeCycle.ValidateParam(i.(godd.ValidateParam))

	i = handlerlifeCycleChecker("ValidateQuery", api.LifeCycle.GetValidateQuery(), api.middleware.LifeCycle.GetValidateQuery(), handlerValidateQuery())
	api.LifeCycle.ValidateQuery(i.(godd.ValidateQuery))

	i = handlerlifeCycleChecker("ParseRequest", api.LifeCycle.GetParseRequest(), api.middleware.LifeCycle.GetParseRequest(), handlerParseRequestDefault())
	api.LifeCycle.ParseRequest(i.(godd.ParseRequest))

	i = handlerlifeCycleChecker("ValidateRequest", api.LifeCycle.GetValidateRequest(), api.middleware.LifeCycle.GetValidateRequest(), handlerValidateRequestDefault())
	api.LifeCycle.ValidateRequest(i.(godd.ValidateRequest))

	i = handlerlifeCycleChecker("OnPreHandler", api.LifeCycle.GetOnPreHandler(), api.middleware.LifeCycle.GetOnPreHandler(), handlerOnPreHandlerDefault())
	api.LifeCycle.OnPreHandler(i.(godd.OnPreHandler))

	i = handlerlifeCycleChecker("HandlerLogic", api.LifeCycle.GetHandlerLogic(), api.middleware.LifeCycle.GetHandlerLogic(), handlerHandlerLogicDefault())
	api.LifeCycle.HandlerLogic(i.(godd.HandlerLogic))

	i = handlerlifeCycleChecker("OnPostHandler", api.LifeCycle.GetOnPostHandler(), api.middleware.LifeCycle.GetOnPostHandler(), handlerOnPostHandlerDefault())
	api.LifeCycle.OnPostHandler(i.(godd.OnPostHandler))

	i = handlerlifeCycleChecker("MappingResponse", api.LifeCycle.GetMappingResponse(), api.middleware.LifeCycle.GetMappingResponse(), handlerMappingResponseDefault())
	api.LifeCycle.MappingResponse(i.(godd.MappingResponse))

	i = handlerlifeCycleChecker("ValidateResponse", api.LifeCycle.GetValidateResponse(), api.middleware.LifeCycle.GetValidateResponse(), handlerValidateResponseDefault())
	api.LifeCycle.ValidateResponse(i.(godd.ValidateResponse))

	i = handlerlifeCycleChecker("MappingResponseStandard", api.LifeCycle.GetMappingResponseStandard(), api.middleware.LifeCycle.GetMappingResponseStandard(), handlerMappingResponseStandardDefault())
	api.LifeCycle.MappingResponseStandard(i.(godd.MappingResponseStandard))

	i = handlerlifeCycleChecker("OnPreResponse", api.LifeCycle.GetOnPreResponse(), api.middleware.LifeCycle.GetOnPreResponse(), handlerOnPreResponseDefault())
	api.LifeCycle.OnPreResponse(i.(godd.OnPreResponse))

	i = handlerlifeCycleChecker("SendResponse", api.LifeCycle.GetSendResponse(), api.middleware.LifeCycle.GetSendResponse(), handlerSendResponseDefault())
	api.LifeCycle.SendResponse(i.(godd.SendResponse))
}

func handlerlifeCycleChecker(name string, api interface{}, mdw interface{}, setDefault interface{}) interface{} {

	if !godd.IsInterfaceIsNil(mdw) && !godd.IsInterfaceIsNil(api) {
		log.Println(name + " : Exist in Middleware and API. Finally Override by Middleware")
	}

	if !godd.IsInterfaceIsNil(mdw) {
		// log.Println(name + " : SetMiddleware")
		return mdw

	}

	if !godd.IsInterfaceIsNil(api) {
		// log.Println(name + " : SetAPI")
		return api
	}

	// log.Println(name + " : SetDefault")
	return setDefault
}

//========================================

func encodeErrorHandler(context godd.InterfaceContext, goddErr *godd.Error) error {

	// errorreporting.LogError(errors.New(errorMessage))
	context.SetContentType("application/json; charset=utf-8")

	context.Response(
		godd.ResponseDataList{
			Success: false,
			Message: "unsuccess",
			ResponseError: &godd.ResponseError{
				Message:  goddErr.GetMessage(),
				Validate: goddErr.GetErrorValidate(),
			},
		},
		goddErr.Code,
	)
	return nil
}

// MappingStandard Func
func MappingStandard(code int, dataList interface{}, responsePagination *godd.ResponsePagination) (response godd.ResponseDataList, goddgoddErr *godd.Error) {
	if dataList == nil {
		dataList = []string{}
	}

	if reflect.ValueOf(dataList).Kind() != reflect.Slice {
		log.Println("Warning : Input Data Response Not Array")
		dataList = append(([]interface{}{}), dataList)
	}

	isSuccess := code/100 == 2
	var message string
	if isSuccess {
		message = "success"
	} else {
		message = "unsuccess"
	}
	response = godd.ResponseDataList{
		Code:               code,
		Success:            isSuccess,
		Message:            message,
		Data:               dataList,
		ResponsePagination: responsePagination,
	}
	return response, nil
}
