package webService

import (
	"log"
)

func RunApi() {
	Configure()
	Register()

	if err := run("8080"); err != nil {
		log.Println("error running server", err)
	}
}

func run(port string) error {
	return Router.Run(":" + port)
}
