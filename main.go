package main

// https://go.dev/doc/tutorial/web-service-gin
// https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/

import (
	"net/http"
	"web-service-gin/controllers"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router = controllers.ApplyAlbumRouter(router)

	return router
}

func main() {
	router := SetupRouter()

	models.ConnectDatabase()

	s := &http.Server{
		Addr:    ":8123",
		Handler: router,
	}
	s.ListenAndServe()
}
