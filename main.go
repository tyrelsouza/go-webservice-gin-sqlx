package main

import "web-service-gin/internal/server"

// Initial based on:
// https://go.dev/doc/tutorial/web-service-gin
// https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
// Then moved to Sqlx
// https://github.com/wetterj/gin-sqlx-crud

func main() {
	server, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	server.Gin.Run()
}
