package main

import (
	"fossil/api"
	"fossil/database"
	"fossil/repo"
	"fossil/router"
	"fossil/service"
	"log"
	"net/http"
)

func main() {

	// setup database connection
	err := database.Connect("root:root@tcp(127.0.0.1:3306)/fossil_link?parseTime=true")
	if err != nil {
		// oops failed connecting
		panic(err)
	}
	database.Conn.Migrate()

	// bootstrap request validator singleton
	api.BootstrapRequestValidator()
	repo.BootstrapRepositories()
	service.BootstrapServices()

	handler := router.BootstrapRouter()

	server := http.Server{Addr: ":9090", Handler: handler}
	log.Fatal(server.ListenAndServe())
}
