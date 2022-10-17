package controllers

import (
	"cloudhumans/internal/models"
	"cloudhumans/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

const OK_ALLOCATE = "Hello from cloudhumans."
const ERR_BINDING = "ERR. Binding parameters."

type ProjectsController struct {
	service *services.ProjectsService
}

func (controller *ProjectsController) LoadRoutes(gr *echo.Group) {
	projectsGroup := gr.Group("/projects")
	projectsGroup.POST("/allocate", controller.Allocate)
}

func (controller *ProjectsController) Allocate(c echo.Context) error {
	req := &models.Request{}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, ERR_BINDING)
	}

	err := controller.service.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}

	//score := controller.service.CalculateEligibilityScore(req)

	return nil
}
