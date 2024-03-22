package controller

import (
	"github.com/aldiaprilianto/takana/dto/response"
	"github.com/aldiaprilianto/takana/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GenerateOrganization(ctx *gin.Context) *response.OrgHierarchyResponse
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (n userController) GenerateOrganization(ctx *gin.Context) *response.OrgHierarchyResponse {

	data, errs := n.userService.GenerateOrganization(ctx)
	if errs != nil {
		return nil
	}

	return data
}
