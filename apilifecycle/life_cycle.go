package apilifecycle

import (
	"log"

	godd "github.com/pagongamedev/go-dd"
)

// ================== By Pass =======================

func (api *APILifeCycle) HandlerLifeCycle(context *godd.Context) (*APIResponse, error) {
	var err error

	if api.byPass != nil {
		return nil, api.byPass(context)
	}

	code, response, goddErr := api.coreLifeCycle(context)

	if goddErr != nil {
		code, response, goddErr = api.GetMappingStandardError()(goddErr)
		if goddErr != nil && goddErr.Error != nil {
			err = goddErr.Error
			log.Println(err)
		}
	}

	return &APIResponse{
		Code:     code,
		Response: response,
	}, err
}

func (api *APILifeCycle) coreLifeCycle(context *godd.Context) (int, interface{}, *godd.Error) {
	var goddErr *godd.Error

	if context != nil {
		context.ClearState()
	}

	// ================== Pre Start List =======================

	goddErr = api.GetOnPreStartList()(context)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	// ================== Start List =======================

	if !godd.IsInterfaceIsNil(api.middlewareStartList) {
		for _, startList := range api.middlewareStartList {
			goddErr = startList(context)
			if goddErr != nil {
				return 0, nil, goddErr
			}
		}
	}

	// ================== On Start =======================

	goddErr = api.GetOnStart()(context)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	goddErr = api.GetParseLanguage()(context)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	// ================== Auth =======================

	goddErr = api.GetOnPreAuth()(context)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	roleData, goddErr := api.GetValidateAuth()(context)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	goddErr = api.GetValidateRole()(context, roleData)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	goddErr = api.GetOnPostAuth()(context)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	// ================== Validate Request =======================

	goddErr = api.GetValidateHeader()(context)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	requestValidatedParam, goddErr := api.GetValidateParam()(context)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	requestValidatedQuery, goddErr := api.GetValidateQuery()(context)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	requestMappingBody, goddErr := api.GetParseRequest()(context)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	requestValidatedBody, goddErr := api.GetValidateRequest()(context, requestMappingBody)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	// ================== Handler =======================

	requestValidatedBody, requestValidatedParam, requestValidatedQuery, goddErr = api.GetOnPreHandler()(context, requestValidatedBody, requestValidatedParam, requestValidatedQuery)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	code, responseRaw, responsePagination, goddErr := api.GetHandlerLogic()(context, requestValidatedBody, requestValidatedParam, requestValidatedQuery)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	code, responseRaw, responsePagination, goddErr = api.GetOnPostHandler()(context, code, responseRaw, responsePagination)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	// ================== Validate Response =======================

	code, responseMapping, responsePagination, goddErr := api.GetMappingResponse()(context, code, responseRaw, responsePagination)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	code, responseValidated, responsePagination, goddErr := api.GetValidateResponse()(context, code, responseMapping, responsePagination)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	code, responseStandard, goddErr := api.GetMappingStandardResponse()(context, code, responseValidated, responsePagination)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	// ================== Pre End List =======================

	code, responseStandard, goddErr = api.GetOnPreEndList()(context, code, responseStandard)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	// ================== End List =======================

	if !godd.IsInterfaceIsNil(api.middlewareEndList) {
		for _, endList := range api.middlewareEndList {
			code, responseStandard, goddErr = endList(context, code, responseStandard)
			if goddErr != nil {
				return 0, nil, goddErr
			}
		}
	}

	// ====================================================

	code, responseStandard, goddErr = api.GetOnPreResponse()(context, code, responseStandard)
	if goddErr != nil {
		return 0, nil, goddErr
	}

	// ========================================================================

	// goddErr = api.GetSendResponse()(context, code, responseStandard)
	// if goddErr != nil {
	// 	return 0, nil, goddErr
	// }

	return code, responseStandard, goddErr
}
