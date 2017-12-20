package main

import (
	"net/http"
	"github.com/labstack/echo"
	"echo_note/controller"
	"time"
	"fmt"
	"github.com/labstack/echo/middleware"
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
	//实例化echo
	e := echo.New()

	//注册路由
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
	})

	//注册到控制器
	e.GET("/hello", controller.Hello)

	//注册函数
	e.GET("/wcookie", wcookie)
	e.GET("/rcookie", rcookie)
	e.GET("/rallcookie", rallcookie)

	//注册组路由
	g := e.Group("/admin")

	//中间件
	g.Use(middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))
	g.GET("/login", controller.Hello)

	//监听端口
	s := &http.Server{
		Handler: e,
		Addr: ":8080",
	}

	//运行服务
	e.Logger.Fatal(s.ListenAndServe())
}
