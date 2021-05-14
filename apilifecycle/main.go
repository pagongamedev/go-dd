package apilifecycle

import (
	"log"

	godd "github.com/pagongamedev/go-dd"
)

// APILifeCycle Struct
type APILifeCycle struct {
	// ==========  By Pass   ==========
	byPass HandlerByPass
	// ==========    Error   ==========
	mappingStandardError MappingStandardError
	// ========== Life Cycle ==========
	onPreStartList          HandlerCycle
	middlewareStartList     []HandlerCycle
	onStart                 HandlerCycle
	parseLanguage           HandlerCycle
	onPreAuth               HandlerCycle
	validateAuth            ValidateAuth
	validateRole            ValidateRole
	onPostAuth              HandlerCycle
	validateHeader          HandlerCycle
	validateParam           ValidateParam
	validateQuery           ValidateQuery
	parseRequest            ParseRequest
	validateRequest         ValidateRequest
	onPreHandler            OnPreHandler
	handlerLogic            HandlerLogic
	onPostHandler           OnPostHandler
	mappingResponse         MappingResponse
	validateResponse        ValidateResponse
	mappingStandardResponse MappingStandardResponse
	onPreEndList            OnPreResponse
	middlewareEndList       []OnPreResponse
	onPreResponse           OnPreResponse
	// ==================================
	sendResponse SendResponse
}

type APIResponse struct {
	ContentType string
	Code        int
	Response    interface{}
}

func (apiLifeCycle *APILifeCycle) CheckerLifeCycle(apiMiddleware *APILifeCycle, mappingStandardResponse MappingStandardResponse, mappingStandardError MappingStandardError) {

	if mappingStandardResponse == nil {
		mappingStandardResponse = handlerDefaultMappingStandardResponse()
	}

	if mappingStandardError == nil {
		mappingStandardError = handlerDefaultMappingStandardError()
	}

	apiLifeCycle.MappingStandardError(checkerLifeCycleHandler("MappingStandardError", apiLifeCycle.GetMappingStandardError(), apiMiddleware.GetMappingStandardError(), mappingStandardError).(MappingStandardError))
	// ========================================================================

	apiLifeCycle.OnPreStartList(checkerLifeCycleHandler("OnPreStartList", apiLifeCycle.GetOnPreStartList(), apiMiddleware.GetOnPreStartList(), handlerDefaultCycle()).(HandlerCycle))
	// MiddlewareStartList
	apiLifeCycle.OnStart(checkerLifeCycleHandler("OnStart", apiLifeCycle.GetOnStart(), apiMiddleware.GetOnStart(), handlerDefaultCycle()).(HandlerCycle))
	apiLifeCycle.ParseLanguage(checkerLifeCycleHandler("ParseLanguage", apiLifeCycle.GetParseLanguage(), apiMiddleware.GetParseLanguage(), handlerDefaultParseLanguage()).(HandlerCycle))

	apiLifeCycle.OnPreAuth(checkerLifeCycleHandler("OnPreAuth", apiLifeCycle.GetOnPreAuth(), apiMiddleware.GetOnPreAuth(), handlerDefaultCycle()).(HandlerCycle))
	apiLifeCycle.ValidateAuth(checkerLifeCycleHandler("ValidateAuth", apiLifeCycle.GetValidateAuth(), apiMiddleware.GetValidateAuth(), handlerDefaultValidateAuth()).(ValidateAuth))
	apiLifeCycle.ValidateRole(checkerLifeCycleHandler("ValidateRole", apiLifeCycle.GetValidateRole(), apiMiddleware.GetValidateRole(), handlerDefaultValidateRole()).(ValidateRole))
	apiLifeCycle.OnPostAuth(checkerLifeCycleHandler("OnPostAuth", apiLifeCycle.GetOnPostAuth(), apiMiddleware.GetOnPostAuth(), handlerDefaultCycle()).(HandlerCycle))

	apiLifeCycle.ValidateHeader(checkerLifeCycleHandler("ValidateHeader", apiLifeCycle.GetValidateHeader(), apiMiddleware.GetValidateHeader(), handlerDefaultCycle()).(HandlerCycle))

	apiLifeCycle.ValidateParam(checkerLifeCycleHandler("ValidateParam", apiLifeCycle.GetValidateParam(), apiMiddleware.GetValidateParam(), handlerDefaultValidateParam()).(ValidateParam))
	apiLifeCycle.ValidateQuery(checkerLifeCycleHandler("ValidateQuery", apiLifeCycle.GetValidateQuery(), apiMiddleware.GetValidateQuery(), handlerDefaultValidateQuery()).(ValidateQuery))
	apiLifeCycle.ParseRequest(checkerLifeCycleHandler("ParseRequest", apiLifeCycle.GetParseRequest(), apiMiddleware.GetParseRequest(), handlerDefaultParseRequest()).(ParseRequest))
	apiLifeCycle.ValidateRequest(checkerLifeCycleHandler("ValidateRequest", apiLifeCycle.GetValidateRequest(), apiMiddleware.GetValidateRequest(), handlerDefaultValidateRequest()).(ValidateRequest))

	apiLifeCycle.OnPreHandler(checkerLifeCycleHandler("OnPreHandler", apiLifeCycle.GetOnPreHandler(), apiMiddleware.GetOnPreHandler(), handlerDefaultOnPreHandler()).(OnPreHandler))
	apiLifeCycle.HandlerLogic(checkerLifeCycleHandler("HandlerLogic", apiLifeCycle.GetHandlerLogic(), apiMiddleware.GetHandlerLogic(), handlerDefaultHandlerLogic()).(HandlerLogic))
	apiLifeCycle.OnPostHandler(checkerLifeCycleHandler("OnPostHandler", apiLifeCycle.GetOnPostHandler(), apiMiddleware.GetOnPostHandler(), handlerDefaultOnPostHandler()).(OnPostHandler))

	apiLifeCycle.MappingResponse(checkerLifeCycleHandler("MappingResponse", apiLifeCycle.GetMappingResponse(), apiMiddleware.GetMappingResponse(), handlerDefaultMappingResponse()).(MappingResponse))
	apiLifeCycle.ValidateResponse(checkerLifeCycleHandler("ValidateResponse", apiLifeCycle.GetValidateResponse(), apiMiddleware.GetValidateResponse(), handlerDefaultValidateResponse()).(ValidateResponse))
	apiLifeCycle.MappingStandardResponse(checkerLifeCycleHandler("MappingStandardResponse", apiLifeCycle.GetMappingStandardResponse(), apiMiddleware.GetMappingStandardResponse(), mappingStandardResponse).(MappingStandardResponse))

	apiLifeCycle.OnPreEndList(checkerLifeCycleHandler("OnPreEndList", apiLifeCycle.GetOnPreEndList(), apiMiddleware.GetOnPreEndList(), handlerDefaultOnPreResponse()).(OnPreResponse))
	// MiddlewareEndList
	apiLifeCycle.OnPreResponse(checkerLifeCycleHandler("OnPreResponse", apiLifeCycle.GetOnPreResponse(), apiMiddleware.GetOnPreResponse(), handlerDefaultOnPreResponse()).(OnPreResponse))

	apiLifeCycle.SendResponse(checkerLifeCycleHandler("SendResponse", apiLifeCycle.GetSendResponse(), apiMiddleware.GetSendResponse(), handlerDefaultSendResponse()).(SendResponse))
}

// =====================================================================
// EncodeError Func
type EncodeError = func(context *godd.Context, goddErr *godd.Error) (interface{}, error)

// HandlerCycle Type
type HandlerCycle = func(context *godd.Context) (goddErr *godd.Error)

// Handler Default
func handlerDefaultCycle() HandlerCycle {
	return func(context *godd.Context) (goddErr *godd.Error) {
		return nil
	}
}

// =====================================================================

func checkerLifeCycleHandler(name string, api interface{}, mdw interface{}, setDefault interface{}) interface{} {

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
