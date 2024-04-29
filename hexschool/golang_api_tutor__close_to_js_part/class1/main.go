package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	PORT = ":8888"
)

func main() {
	var err error

	repository := NewRepository()
	controller := NewController(repository)

	server := gin.Default()
	server = setupRouter(server, controller)
	err = server.Run(PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter(router *gin.Engine, controller Controller) *gin.Engine {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/v1/create", controller.CreatePostController)
	router.POST("/api/v1/posts", controller.ReadPostsController)

	return router
}
