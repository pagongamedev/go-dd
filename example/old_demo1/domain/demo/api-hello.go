package demo

import (
	"net/http"

	godd "github.com/pagongamedev/go-dd"
	goddAPI "github.com/pagongamedev/go-dd/api"
)

// Job struct
type Job struct {
	Type   string `json:"type"   default:"none" validate:"required,min=3,max=32"`
	Salary int    `json:"salary" default:"200"  validate:"required,number"`
}

//User struct
type User struct {
	Name     string `json:"name"      default:"john" validate:"required,min=3,max=32"`
	IsActive bool   `json:"is_active" default:"true" validate:"required,eq=True|eq=False"`
	Email    string `json:"email"                    validate:"required,email,min=6,max=32"`
	Job      Job    `json:"job"                      validate:"dive"`
}

// HandlerHello API
func HandlerHello() *goddAPI.HTTP {
	api := goddAPI.NewAPIHTTP()

	api.LifeCycle.ValidateRequest(func(context *godd.Context, requestMappingBody interface{}) (requestValidatedBody interface{}, goddErr *godd.Error) {
		user := new(User)

		context.BodyParser(user)
		// SetDefault Request
		context.SetDefaultStruct(user)
		// Validate Request
		errors := context.ValidateStruct(user, godd.Map{"User": &User{}, "Job": &Job{}})
		if errors != nil {
			return http.StatusBadRequest, nil
		}
		return user, nil
	})
	api.LifeCycle.HandlerLogic(func(context *godd.Context, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, goddErr *godd.Error) {
		user := requestValidatedBody.(*User)
		return http.StatusOK, user, nil, nil
	})
	return api
}
