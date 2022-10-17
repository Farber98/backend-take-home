package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const OK_HELLO = "Hello from cloudhumans."

type HelloController struct{}

func (controller *HelloController) LoadRoutes(gr *echo.Group) {
	gr.GET("/hello", controller.Hello)
}

func (controller *HelloController) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, OK_HELLO)
}
