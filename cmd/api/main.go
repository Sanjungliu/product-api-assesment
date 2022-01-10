package main

import (
	"database/sql"
	"log"

	"github.com/Sanjungliu/product-api-assesment/config"
	"github.com/Sanjungliu/product-api-assesment/internal/httpserver"
	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

func main() {
	config := config.Init()
	db, err := sql.Open("postgres", config.DBConnectionString())
	if err != nil {
		log.Fatal("failed connect to database")
	}

	redis := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr(),
	})
	_, err = redis.Ping().Result()
	if err != nil {
		log.Println("FAIL: connection to redis failed.")
		panic(err)
	}

	app := buildInternalService(db, redis)

	server := httpserver.NewServer(app)
	server.Serve()
}
