package models

import (
	"github.com/guregu/null"
	"web-service-gin/internal/forms"
)

type Album struct {
	ID        int64       `db:"id"`
	CreatedAt null.String `db:"created_at"`
	UpdatedAt null.String `db:"updated_at"`
	DeletedAt null.String `db:"deleted_at"`
	Title     string      `db:"title"`
	Artist    string      `db:"artist"`
	Price     float64     `db:"price"`
}

func (a *Album) ApplyForm(form *forms.CreateAlbum) {
	a.ID = *form.ID
	a.Title = *form.Title
	a.Artist = *form.Artist
	a.Price = *form.Price
	a.CreatedAt = null.StringFromPtr(form.CreatedAt)
	a.UpdatedAt = null.StringFromPtr(form.UpdatedAt)
	a.DeletedAt = null.StringFromPtr(form.DeletedAt)
}

type AlbumService interface {
	Create(*forms.CreateAlbum) (*Album, error)
	GetByID(string) (*Album, error)
	GetAll() (*[]Album, error)
	Update(*Album) error
	Delete(string) error
}
