package sql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"web-service-gin/internal/forms"
	models2 "web-service-gin/internal/models"
)

// AlbumService is the implementation of the album data mapping layer
// using SQL.
type AlbumService struct {
	conn *sqlx.DB
}

// NewAlbumService creates the album service using the given
// connection pool to a mysql DB.
func NewAlbumService(conn *sqlx.DB) (*AlbumService, error) {
	// TODO: It would be better to use a DB management tool
	// to make migrations painless.

	//	_, err := conn.Exec(`
	//CREATE TABLE albums (
	//                          id bigint NOT NULL AUTO_INCREMENT,
	//                          title longtext,
	//                          artist longtext,
	//                          price double DEFAULT NULL,
	//                          PRIMARY KEY (id),
	//                          KEY idx_albums_deleted_at (deleted_at)
	//) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
	//`)
	//	if err != nil {
	//		return nil, err
	//	}
	return &AlbumService{conn: conn}, nil
}

// Create will try to add the album to the DB.
func (s *AlbumService) Create(form *forms.CreateAlbum) (*models2.Album, error) {
	q := `INSERT INTO albums(title, artist, price) VALUES (?, ?, ?);`

	result, err := s.conn.Exec(
		q,
		*form.Title,
		*form.Artist,
		*form.Price,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	album, err := s.GetByID(strconv.FormatUint(uint64(id), 10))
	if err != nil {
		return nil, err
	}
	return album, nil
}

// Update will replace the values of the give album with those provided.
func (s *AlbumService) Update(p *models2.Album) error {
	q := `
UPDATE albums
SET updated_at = NOW(),
    title = ?,
    artist = ?,
    price = ?
WHERE id = ?;
`

	_, err := s.conn.Exec(
		q,
		p.Title,
		p.Artist,
		p.Price,
		p.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetByID fetches the album with the given id.
func (s *AlbumService) GetByID(id string) (*models2.Album, error) {
	if !validID(id) {
		return nil, models2.ErrNotFound
	}

	q := `
SELECT *
FROM albums
WHERE id = ?;`

	var output models2.Album
	err := s.conn.Get(
		&output,
		q,
		id,
	)
	// Replace the SQL error with our own error type.
	if err == sql.ErrNoRows {
		return nil, models2.ErrNotFound
	} else if err != nil {
		return nil, err
	} else {
		return &output, nil
	}
}

// GetAll fetches all albums.
func (s *AlbumService) GetAll() (*[]models2.Album, error) {
	q := `SELECT * FROM albums;`

	var output []models2.Album
	err := s.conn.Select(&output, q)
	// Replace the SQL error with our own error type.
	if err == sql.ErrNoRows {
		return nil, models2.ErrNotFound
	} else if err != nil {
		return nil, err
	} else {
		return &output, nil
	}
}

// Delete removes the album with the given id from the DB.
// TODO: this should just mark the object as deleted,
// not actually get rid of the data.
func (s *AlbumService) Delete(id string) error {
	if !validID(id) {
		return models2.ErrNotFound
	}

	q := `
DELETE FROM albums
WHERE id = ?;
`

	_, err := s.conn.Exec(
		q,
		id,
	)
	return err
}

// Check it implements the interface
var _ models2.AlbumService = &AlbumService{}
