package notexample

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
	goddAPI "github.com/pagongamedev/go-dd/api"
	goddMicroService "github.com/pagongamedev/go-dd/microservice"
)

//================================================================================
//            	  Just Resoure Please See Example in Main.go
//================================================================================

// ============ Microservice ===============

// Router Func
func Router(app *fiber.App, path string) *goddMicroService.MicroService {
	ms := goddMicroService.New(app, path, nil, nil, nil)
	ms.Get("/hello", handlerHello())
	return ms
}

// ============ API ===============

// @Router /hello/v1/hello [get]
// @Summary Hello Go DD
// @Description Hello Go DD
// @Tags hello
// @Accept  json
// @Produce  json
// @Success 200 {object} ResponseDataList
func handlerHello() *goddAPI.HTTP {
	api := goddAPI.NewAPIHTTP()

	api.LifeCycle.HandlerLogic(func(context godd.InterfaceContext, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, goddErr *godd.Error) {
		return 200, godd.ConvertToArray(godd.Map{"Message": "Hello Go-DD"}), nil, nil
	})

	return api
}

// ============ Test Close ===============

// DummyDatabase func
func DummyDatabase() (interface{}, error) {
	return &DummyClose{}, nil

}

// DummyClose struct
type DummyClose struct {
}

// Close func
func (dc *DummyClose) Close() error {
	log.Println("Defer Close")
	return nil
}

//===========================================

// ResponseDataList for Send Response Message to Encode Response
type ResponseDataList struct {
	Success            bool                `json:"success"              swaggertype:"boolean"`
	Message            string              `json:"message"              swaggertype:"string"`
	Data               interface{}         `json:"data,omitempty"       swaggertype:"object"`
	ResponsePagination *ResponsePagination `json:"pagination,omitempty" swaggertype:"object"`
	ResponseError      *ResponseError      `json:"error,omitempty"      swaggertype:"object"`
}

// ResponsePagination Struct
type ResponsePagination struct {
	Page       int `json:"page"         swaggertype:"integer"`
	PageSize   int `json:"page_size"    swaggertype:"integer"`
	PageCount  int `json:"page_count"   swaggertype:"integer"`
	ItemCount  int `json:"item_count"   swaggertype:"integer"`
	TotalCount int `json:"total_count"  swaggertype:"integer"`
}

// ResponseError Struct
type ResponseError struct {
	Message  string                            `json:"message" swaggertype:"string"`
	Validate *map[string]ResponseErrorValidate `json:"validate,omitempty" swaggertype:"object"`
}

// ResponseErrorValidate struct
type ResponseErrorValidate struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Param   string `json:"param,omitempty"`
}

// ==============================================================
