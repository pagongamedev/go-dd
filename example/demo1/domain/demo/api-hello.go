package demo

import (
	"net/http"

	godd "github.com/pagongamedev/go-dd"
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
func HandlerHello() *godd.APIHTTP {
	api := godd.NewAPIHTTP()

	api.HandlerLogic(func(context godd.InterfaceContext, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, err *godd.Error) {
		user := new(User)

		if err := context.BodyParser(user); err != nil {
			// c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			// 	message: err.Error(),
			// })
			// return
		}

		context.SetDefaultStruct(user)

		// x := godd.Map{
		// 	"User": &User{},
		// 	"Job":  &Job{},
		// }

		// errors := context.ValidateStruct(User{
		// 	Name:     "TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT",
		// 	IsActive: false,
		// 	Email:    "a",
		// 	Job: Job{
		// 		Type:   "E",
		// 		Salary: 200,
		// 	},
		// }, x)

		// if errors != nil {
		// 	// c.JSON(errors)
		// 	return http.StatusBadRequest, nil, nil, errors
		// }

		return http.StatusOK, user, nil, nil
	})
	return api
}
