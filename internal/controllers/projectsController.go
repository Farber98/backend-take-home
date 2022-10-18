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
	Service *services.ProjectsService
}

func (controller *ProjectsController) LoadRoutes(gr *echo.Group) {
	projectsGroup := gr.Group("/projects")
	projectsGroup.POST("/allocate", controller.Allocate)
}

func (controller *ProjectsController) Allocate(c echo.Context) error {
	req := &models.Request{}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewMsgResponse(ERR_BINDING))
	}

	err := controller.Service.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewMsgResponse(err.Error()))

	}

	score := controller.Service.CalculateEligibilityScore(req)

	selected, eligible, notEligible := controller.Service.Projects(score)

	response := &models.Response{
		Score:              score,
		SelectedProject:    selected,
		EligibleProjects:   eligible,
		InelegibleProjects: notEligible,
	}

	return c.JSON(http.StatusOK, response)
}
