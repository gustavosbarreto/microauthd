package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

func login(c echo.Context) error {
	return nil
}

func signup(c echo.Context) error {
	return nil
}

func getAccount(c echo.Context) error {
	return nil
}

func main() {
	e := echo.New()

	e.POST("/login", login)
	e.POST("/signup", signup)

	e.GET("/account", getAccount)

	logrus.Fatal(e.Start(":8080"))
}
