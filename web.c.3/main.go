package main

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   string `json:"age" validate:"gte=0,lte=80"`
}

type CostumValidator struct {
	validator *validator.Validate
}

func (cv *CostumValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CostumValidator{validator: validator.New()}

	e.POST("/users", func(c echo.Context) error {
		u := new(User)

		if err := c.Bind(u); err != nil {
			return err
		}

		if err := c.Validate(u); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":9001"))
}
