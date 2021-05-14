package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// ParseLanguage Set
func (api *APILifeCycle) ParseLanguage(handler HandlerCycle) {
	api.parseLanguage = handler
}

// GetParseLanguage Get
func (api *APILifeCycle) GetParseLanguage() HandlerCycle {
	return api.parseLanguage
}

// Handler Default
func handlerDefaultParseLanguage() HandlerCycle {
	return func(context godd.InterfaceContext) (goddErr *godd.Error) {

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
