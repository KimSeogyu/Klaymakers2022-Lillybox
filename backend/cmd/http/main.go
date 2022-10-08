package main

import (
	_ "lillybox-backend/docs"
	"lillybox-backend/internal/http"
	"log"
)

// Bootstrap
// @title          Lillybox Backend
// @version        0.1
// @description    This is API Documentation Lillybox Backend
// @termsOfService http://swagger.io/terms/
// @contact.name   Backend-Tech
// @contact.email  asap0208@gmail.com
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @host           3.36.115.152:8080
// @BasePath       /
func main() {
	app := http.Server{}
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic")
		}
	}()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
