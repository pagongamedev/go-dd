package demo

import (
	"net/http"

	godd "github.com/pagongamedev/go-dd"
)

// Job struct
type Job struct {
	Type   string `json:"type" validate:"required,min=3,max=32"`
	Salary int    `json:"salary" validate:"required,number"`
}

//User struct
type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive bool   `json:"is_active" validate:"required,eq=True|eq=False"`
	Email    string `json:"email"   validate:"required,email,min=6,max=32"`
	Job      Job    `json:"job"  validate:"dive"`
}

// HandlerHello API
func HandlerHello() *godd.APIHTTP {
	api := godd.NewAPIHTTP()

	api.HandlerLogic(func(context godd.InterfaceContext, requestValidated interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, err *godd.Error) {
		user := new(User)

		if err := context.BodyParser(user); err != nil {
			// c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			// 	message: err.Error(),
			// })
			// return
		}

		// field, _ := reflect.TypeOf(user).Elem().FieldByName("IsActive")
		// tag := string(field.Tag)
		// log.Println("Tag : ", tag)

		//errors := ValidateStruct(*user)
		x := map[string]interface{}{
			"User": &User{},
			"Job":  &Job{},
		}

		errors := godd.ValidateStruct(User{
			Name:     "T",
			IsActive: false,
			Email:    "a",
			Job: Job{
				Type:   "E",
				Salary: 20,
			},
		}, x)

		if errors != nil {
			// c.JSON(errors)
			return http.StatusBadRequest, nil, nil, errors
		}

		return http.StatusOK, godd.Map{"message": "Helllllo"}, nil, nil
	})
	return api
}
