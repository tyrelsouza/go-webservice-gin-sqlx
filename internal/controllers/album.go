package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"web-service-gin/internal/forms"
	models2 "web-service-gin/internal/models"
)

// Album provides the handlers for the album entity.
type Album struct {
	albumService models2.AlbumService
}

// NewAlbum creates the controller using the given data mapper for
// albums.
func NewAlbum(albumService models2.AlbumService) *Album {
	return &Album{
		albumService: albumService,
	}
}

// Post will create a new album from the given data, if the form is valid.
func (p *Album) Post(c *gin.Context) {
	var form forms.CreateAlbum
	if c.ShouldBindWith(&form, binding.JSON) != nil {
		// TODO: Give a better error message.
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{"message": "invalid data."},
		)
		c.Abort()
		return
	}

	album, err := p.albumService.Create(&form)
	if err != nil {
		// TODO: An error middleware should log the error,
		// and email admin.
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error"},
		)
		c.Abort()
		return
	}

	// TODO: use a view if part of the album data should not be
	// returned to the client.
	c.JSON(
		http.StatusCreated,
		album,
	)
}

// Put will perform an update of a album.
func (p *Album) Put(c *gin.Context) {
	var form forms.CreateAlbum
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		// TODO: Give a better error message.
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": "invalid data.",
				"form":    form,
				"error":   err.Error(),
			},
		)
		c.Abort()
		return
	}
	id := c.Param("id")

	album, err := p.albumService.GetByID(id)
	if err == models2.ErrNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": "user not found"},
		)
		c.Abort()
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	album.ApplyForm(&form)
	err = p.albumService.Update(album)
	if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": "updated"},
	)
}

// Get will fetch an album by ID.
func (p *Album) Get(c *gin.Context) {
	id := c.Param("id")
	album, err := p.albumService.GetByID(id)
	if err == models2.ErrNotFound {
		c.JSON(
			http.StatusNotFound,
			gin.H{"message": "user not found"},
		)
		c.Abort()
		return
	} else if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		album,
	)
}

// GetAll will fetch all Albums.
// TODO: Pagination
func (p *Album) GetAll(c *gin.Context) {
	albums, err := p.albumService.GetAll()
	if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
		return
	}
	c.JSON(
		http.StatusOK,
		albums,
	)
}

// Delete will remove a album from the DB.
func (p *Album) Delete(c *gin.Context) {
	id := c.Param("id")

	err := p.albumService.Delete(id)

	if err != nil {
		c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "internal error."},
		)
		c.Abort()
	}

	c.JSON(
		http.StatusOK,
		gin.H{"message": "deleted"},
	)
}
