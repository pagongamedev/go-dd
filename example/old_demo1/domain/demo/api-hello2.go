package demo

import (
	"log"
	"net/http"

	godd "github.com/pagongamedev/go-dd"
	goddAPI "github.com/pagongamedev/go-dd/api"
)

// HandlerHello2 API
func HandlerHello2() *goddAPI.HTTP {
	api := goddAPI.NewAPIHTTP()

	api.LifeCycle.ValidateAuth(func(context *godd.Context) (roleData interface{}, goddErr *godd.Error) {
		log.Println("API")
		return nil, nil
	})
	api.LifeCycle.HandlerLogic(func(context *godd.Context, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, goddErr *godd.Error) {
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
