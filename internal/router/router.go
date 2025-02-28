package router

import (
	"cloudhumans/internal/controllers"
	"cloudhumans/internal/interfaces"
	"cloudhumans/internal/services"
	"log"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var onceEcho sync.Once
var echoInstance *echo.Echo

//InitRoutes Initializes API routes.
func Init() *echo.Echo {
	onceEcho.Do(func() {
		e := echo.New()
		e.Use(middleware.CORS())
		e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogMethod:  true,
			LogURI:     true,
			LogStatus:  true,
			LogLatency: true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				log.Printf("%v %v | Status: %v | Latency: %v\n", v.Method, v.URI, v.Status, v.Latency)
				return nil
			},
		}))

		arrayControllers := make([]interfaces.IController, 0)
		arrayControllers = append(arrayControllers, &controllers.HelloController{})

		projectService := &services.ProjectsService{}
		arrayControllers = append(arrayControllers, &controllers.ProjectsController{Service: projectService})

		e.Static("/", "internal/docs")

		group := e.Group("")
		for _, c := range arrayControllers {
			c.LoadRoutes(group)
		}

		echoInstance = e
	})
	return echoInstance
}
