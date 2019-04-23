package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	r := echo.New()

	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)

		if err = c.Bind(u); err != nil {
			return c.String(http.StatusOK, err.Error())
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	fmt.Println("server started at :9001")
	r.Start(":9001")
}
