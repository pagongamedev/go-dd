package godd

// APILifeCycle Struct
type APILifeCycle struct {
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
	mappingResponseStandard MappingResponseStandard
	onPreResponse           OnPreResponse
	sendResponse            SendResponse
}

// HandlerByPass Type
type HandlerByPass = func(service interface{}, serviceOptionList map[string]interface{}) Handler

// HandlerCycle Type
type HandlerCycle = func(context InterfaceContext) (err *Error)

// ValidateParam Type
type ValidateParam = func(context InterfaceContext) (requestValidatedParam interface{}, err *Error)

// ValidateQuery Type
type ValidateQuery = func(context InterfaceContext) (requestValidatedQuery interface{}, err *Error)

// ParseRequest Type
type ParseRequest = func(context InterfaceContext) (requestMappingBody interface{}, err *Error)

// ValidateAuth Type
type ValidateAuth = func(context InterfaceContext) (roleData interface{}, err *Error)

// ValidateRole Type
type ValidateRole = func(context InterfaceContext, roleData interface{}) (err *Error)

// ValidateRequest Type
type ValidateRequest = func(context InterfaceContext, requestMappingBody interface{}) (requestValidatedBody interface{}, err *Error)

// OnPreHandler Type
type OnPreHandler = func(context InterfaceContext, requestValidatedBody interface{}, requestValidatedParam interface{}, requestValidatedQuery interface{}) (requestValidatedBodyOut interface{}, requestValidatedParamOut interface{}, requestValidatedQueryOut interface{}, err *Error)

// HandlerLogic Type
type HandlerLogic = func(context InterfaceContext, requestValidatedBody interface{}, requestValidatedParam interface{}, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *ResponsePagination, err *Error)

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

// OnStart func
func (api *APILifeCycle) OnStart(handler HandlerCycle) {
	api.onStart = handler
}

// ParseLanguage func
func (api *APILifeCycle) ParseLanguage(handler HandlerCycle) {
	api.parseLanguage = handler
}

// OnPreAuth func
func (api *APILifeCycle) OnPreAuth(handler HandlerCycle) {
	api.onPreAuth = handler
}

// ValidateAuth func
func (api *APILifeCycle) ValidateAuth(handler ValidateAuth) {
	api.validateAuth = handler
}

// ValidateRole func
func (api *APILifeCycle) ValidateRole(handler ValidateRole) {
	api.validateRole = handler
}

// OnPostAuth func
func (api *APILifeCycle) OnPostAuth(handler HandlerCycle) {
	api.onPostAuth = handler
}

// ValidateHeader func
func (api *APILifeCycle) ValidateHeader(handler HandlerCycle) {
	api.validateHeader = handler
}

// ValidateParam func
func (api *APILifeCycle) ValidateParam(handler ValidateParam) {
	api.validateParam = handler
}

// ValidateQuery func
func (api *APILifeCycle) ValidateQuery(handler ValidateQuery) {
	api.validateQuery = handler
}

// ParseRequest func
func (api *APILifeCycle) ParseRequest(handler ParseRequest) {
	api.parseRequest = handler
}

// ValidateRequest func
func (api *APILifeCycle) ValidateRequest(handler ValidateRequest) {
	api.validateRequest = handler
}

// OnPreHandler func
func (api *APILifeCycle) OnPreHandler(handler OnPreHandler) {
	api.onPreHandler = handler
}

// HandlerLogic func
func (api *APILifeCycle) HandlerLogic(handler HandlerLogic) {
	api.handlerLogic = handler
}

// OnPostHandler func
func (api *APILifeCycle) OnPostHandler(handler OnPostHandler) {
	api.onPostHandler = handler
}

// MappingResponse func
func (api *APILifeCycle) MappingResponse(handler MappingResponse) {
	api.mappingResponse = handler
}

// ValidateResponse func
func (api *APILifeCycle) ValidateResponse(handler ValidateResponse) {
	api.validateResponse = handler
}

// MappingResponseStandard func
func (api *APILifeCycle) MappingResponseStandard(handler MappingResponseStandard) {
	api.mappingResponseStandard = handler
}

// OnPreResponse func
func (api *APILifeCycle) OnPreResponse(handler OnPreResponse) {
	api.onPreResponse = handler
}

// SendResponse func
func (api *APILifeCycle) SendResponse(handler SendResponse) {
	api.sendResponse = handler
}

//========================================

// GetOnStart func
func (api *APILifeCycle) GetOnStart() HandlerCycle {
	return api.onStart
}

// GetParseLanguage func
func (api *APILifeCycle) GetParseLanguage() HandlerCycle {
	return api.parseLanguage
}

// GetOnPreAuth func
func (api *APILifeCycle) GetOnPreAuth() HandlerCycle {
	return api.onPreAuth
}

// GetValidateAuth func
func (api *APILifeCycle) GetValidateAuth() ValidateAuth {
	return api.validateAuth
}

// GetValidateRole func
func (api *APILifeCycle) GetValidateRole() ValidateRole {
	return api.validateRole
}

// GetOnPostAuth func
func (api *APILifeCycle) GetOnPostAuth() HandlerCycle {
	return api.onPostAuth
}

// GetValidateHeader func
func (api *APILifeCycle) GetValidateHeader() HandlerCycle {
	return api.validateHeader
}

// GetValidateParam func
func (api *APILifeCycle) GetValidateParam() ValidateParam {
	return api.validateParam
}

// GetValidateQuery func
func (api *APILifeCycle) GetValidateQuery() ValidateQuery {
	return api.validateQuery
}

// GetParseRequest func
func (api *APILifeCycle) GetParseRequest() ParseRequest {
	return api.parseRequest
}

// GetValidateRequest func
func (api *APILifeCycle) GetValidateRequest() ValidateRequest {
	return api.validateRequest
}

// GetOnPreHandler func
func (api *APILifeCycle) GetOnPreHandler() OnPreHandler {
	return api.onPreHandler
}

// GetHandlerLogic func
func (api *APILifeCycle) GetHandlerLogic() HandlerLogic {
	return api.handlerLogic
}

// GetOnPostHandler func
func (api *APILifeCycle) GetOnPostHandler() OnPostHandler {
	return api.onPostHandler
}

// GetMappingResponse func
func (api *APILifeCycle) GetMappingResponse() MappingResponse {
	return api.mappingResponse
}

// GetValidateResponse func
func (api *APILifeCycle) GetValidateResponse() ValidateResponse {
	return api.validateResponse
}

// GetMappingResponseStandard func
func (api *APILifeCycle) GetMappingResponseStandard() MappingResponseStandard {
	return api.mappingResponseStandard
}

// GetOnPreResponse func
func (api *APILifeCycle) GetOnPreResponse() OnPreResponse {
	return api.onPreResponse
}

// GetSendResponse func
func (api *APILifeCycle) GetSendResponse() SendResponse {
	return api.sendResponse
}

//========================================
