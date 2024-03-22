package router

import (
	"github.com/aldiaprilianto/takana/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Routes interface {
	SetupRoutes(r *gin.Engine)
}

type routes struct {
	userController controller.UserController
}

func NewRoutes(
	userController controller.UserController) Routes {
	return &routes{
		userController: userController,
	}
}

func (ro *routes) SetupRoutes(r *gin.Engine) {

	r.POST("/GenerateJSONStructure/:org_id", ro.generateOrganization)
}

func (ro *routes) generateOrganization(c *gin.Context) {
	response := ro.userController.GenerateOrganization(c)
	c.JSON(http.StatusOK, response)
}
