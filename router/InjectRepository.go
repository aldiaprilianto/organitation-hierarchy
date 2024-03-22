package router

import (
	"github.com/aldiaprilianto/takana/controller"
	"github.com/aldiaprilianto/takana/repository"
	"github.com/aldiaprilianto/takana/service"
	"gorm.io/gorm"
)

func InjectRepository(db *gorm.DB) Routes {

	//Inject Repository
	userRepository := repository.NewUserRepository(db)

	//Inject Service
	userService := service.NewUserService(userRepository)

	//Inject Controller
	userController := controller.NewUserController(userService)

	//Inject Routes
	router := NewRoutes(userController)

	return router
}
