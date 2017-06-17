package router

import (
	"github.com/labstack/echo"
	"net/http"
	"echo_note/controller"
)

func init() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
	})

	e.GET("/hello", controller.Hello)
}