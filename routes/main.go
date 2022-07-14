package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jardelkuhnen/video-api/middlewares"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	// router = gin.Default()
	router = gin.New()
)

func Run() {
	middlewares.SetupLogOutput()
	router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	configureDebugMode()

	getRoutes()
	port := getPortConfig()
	router.Run(":" + port)
}

func getPortConfig() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "8080"
	}

	return port
}

func configureDebugMode() {
	envMode := os.Getenv("ENVIRONMENT")

	if envMode == "debug" {
		fmt.Println("DEBUG is enable. Activating gindump")
		router.Use(gindump.Dump())
	}
}

func getRoutes() {
	v1 := router.Group("/v1")
	addPingRoutes(v1)
	addVideosRoutes(v1)
	addPingAuthorizedRoutes(v1)

	v2 := router.Group("/v2")
	addPingRoutes(v2)
}
