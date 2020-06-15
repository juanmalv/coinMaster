package webService

import (
	"github.com/gin-gonic/gin"
)

var (
	//Router exposes the endpoint of the application.
	Router *gin.Engine
)

// Register API routes
func Register() {
	Router.GET("/ping", Ping)
	Router.POST("/income", NewIncome)
}

// Configure the routing behavior
func Configure() {
	Router = gin.New()
}