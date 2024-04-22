package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"vid-outbox-demo-go/internal/application"
	"vid-outbox-demo-go/internal/persistence"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}
	conn, err := persistence.OpenSQLConnection()
	if err != nil {
		panic(err)
	}
	server := application.New(conn)
	server.RegisterAPIRoutes()
	server.Logger.Fatal(server.Start(":1323"))
}
