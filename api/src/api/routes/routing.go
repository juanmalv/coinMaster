package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/juanmalv/coinMaster/src/api/controller"
)

var (
	//Router exposes the endpoint of the application.
	Router *gin.Engine
)

// Register API routes
func Register() {
	Router.GET("/ping", controller.Ping)
}

// Configure the routing behavior
func Configure() {
	Router = gin.New()
}
