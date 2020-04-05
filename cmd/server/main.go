package main

import (
	"log"

	"github.com/bongnv/expenses/internal/expenseservice"
	"github.com/bongnv/expenses/internal/handlers"
	"github.com/bongnv/expenses/internal/service"
	gokitServer "github.com/bongnv/gokit/util/server"
	"github.com/rs/cors"
)

func main() {
	opts := []gokitServer.Option{
		gokitServer.WithHTTPAddress(":8080"),
		gokitServer.WithCORS(cors.Options{
			AllowedOrigins: []string{"http://localhost:8081"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		}),
	}

	mainHandler := handlers.New()
	opts = append(opts, service.GetOptions(mainHandler)...)
	opts = append(opts, expenseservice.GetOptions(mainHandler)...)

	err := gokitServer.Serve(
		opts...,
	)

	log.Println("Service stopped with:", err)
}
