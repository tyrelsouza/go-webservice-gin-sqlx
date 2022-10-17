package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

// NewSQL creates and SQL connection using environment variables
// to configure.
func NewSQL() (*sqlx.DB, error) {
	//host := strings.TrimSpace(os.Getenv("MYSQL_HOST"))
	//port := strings.TrimSpace(os.Getenv("MYSQL_PORT"))
	//user := strings.TrimSpace(os.Getenv("MYSQL_USER"))
	//password := strings.TrimSpace(os.Getenv("MYSQL_PASSWORD"))
	//db := strings.TrimSpace(os.Getenv("MYSQL_DB"))
	//
	//info := fmt.Sprintf(
	//	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	host,
	//	port,
	//	user,
	//	password,
	//	db,
	//)
	return sqlx.Connect(
		"mysql",
		"mysql:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local",
	)
}

// validID checks if the given string is a valid id.
func validID(id string) bool {
	_, err := strconv.Atoi(id)
	return err == nil
}
