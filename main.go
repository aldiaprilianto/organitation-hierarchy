package main

import (
	"github.com/aldiaprilianto/takana/config"
	"github.com/aldiaprilianto/takana/router"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
	//r.Use(utility.CORS())
	// Middleware untuk menangani CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	/* Routes */
	routing.SetupRoutes(r)

	/* Main.go Running test */
	err = r.Run()
	if err != nil {
		logrus.Fatal(err.Error())
		return
	}

}
