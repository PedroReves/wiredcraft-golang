package routes

import (
	"log"

	"github.com/PedroReves/wiredcraft-golang/controllers"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()

	r.GET("/users", controllers.getUsers)
	r.GET("/users/:id", controllers.getUser)
	r.POST("/users", controllers.createUser)
	r.PUT("/users/:id", controllers.updateUser)
	r.DELETE("/users", controllers.deleteUser)

	if err := r.Run(); err != nil {
		log.Fatalf("Unable to Start Server %v", err)
	}

}
