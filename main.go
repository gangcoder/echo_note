package main

import (
	"net/http"
	"github.com/labstack/echo"
	"echo_note/controller"
	"time"
	"fmt"
)

func wcookie(c echo.Context) error {
	cookie := new (http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24*time.Hour)
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "write a cookie")
}

func rcookie (c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}

	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)

	return c.String(http.StatusOK, "read a cookie")
}

func rallcookie(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "read all cookie")
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
	})

	e.GET("/hello", controller.Hello)

	e.GET("/wcookie", wcookie)
	e.GET("/rcookie", rcookie)
	e.GET("/rallcookie", rallcookie)

	s := &http.Server{
		Handler: e,
		Addr: ":8080",
	}
	e.Logger.Fatal(s.ListenAndServe())
}
