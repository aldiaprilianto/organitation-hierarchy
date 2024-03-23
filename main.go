package main

import (
	"github.com/aldiaprilianto/takana/config"
	"github.com/aldiaprilianto/takana/router"
	"github.com/aldiaprilianto/takana/utility"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	/* Database Connection */
	err := config.InitializeDatabase()
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}
	defer config.CloseDatabase()

	/* Inject Repository */
	db, err := config.GetDB()
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}
	routing := router.InjectRepository(db)

	/* Create Context */
	r := gin.Default()

	/* Middleware CORS */
	r.Use(utility.CORS())

	/* Routes */
	routing.SetupRoutes(r)

	/* Main.go Running test */
	err = r.Run()
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

}
