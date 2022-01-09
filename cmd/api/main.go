package main

import (
	"database/sql"
	"log"

	"github.com/Sanjungliu/product-api-assesment/config"
	"github.com/Sanjungliu/product-api-assesment/internal/httpserver"
	_ "github.com/lib/pq"
)

func main() {
	config := config.Init()
	db, err := sql.Open("postgres", config.DBConnectionString())
	if err != nil {
		log.Fatal("failed connect to database")
	}

	app := buildInternalService(db)

	server := httpserver.NewServer(app)
	server.Serve()
}
