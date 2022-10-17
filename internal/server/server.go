package server

import (
	"github.com/gin-gonic/gin"
	"web-service-gin/internal/controllers"
	"web-service-gin/internal/models"
	"web-service-gin/internal/models/sql"
)

// Server represents all the services and controllers.
type Server struct {
	AlbumService models.AlbumService
	Gin          *gin.Engine
}

// NewServer creates a new server using environment variables to
// configure DB connection.
func NewServer() (*Server, error) {
	db, err := sql.NewSQL()
	if err != nil {
		return nil, err
	}

	albumService, err := sql.NewAlbumService(db)
	if err != nil {
		return nil, err
	}

	r := gin.Default()
	{
		route := r.Group("/albums")
		ctrl := controllers.NewAlbum(albumService)

		route.GET("", ctrl.GetAll)
		route.POST("", ctrl.Post)
		route.PUT("/:id", ctrl.Put)
		route.GET("/:id", ctrl.Get)
		route.DELETE("/:id", ctrl.Delete)
	}

	return &Server{
		AlbumService: albumService,
		Gin:          r,
	}, nil
}
