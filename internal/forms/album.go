package forms

type CreateAlbum struct {
	ID     *int64   `form:"id" json:"id" binding:"required"`
	Title  *string  `form:"title" json:"title" binding:"required"`
	Artist *string  `form:"artist" json:"artist" binding:"required"`
	Price  *float64 `form:"price" json:"price" binding:"required"`

	CreatedAt *string `form:"created_at" json:"created_at"`
	UpdatedAt *string `form:"updated_at" json:"updated_at"`
	DeletedAt *string `form:"deleted_at" json:"deleted_at"`
}
