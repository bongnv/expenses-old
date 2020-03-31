package main

import (
	"log"

	"github.com/bongnv/expenses/internal/expenseservice"
	"github.com/bongnv/expenses/internal/handlers"
	"github.com/bongnv/expenses/internal/service"
	gokitServer "github.com/bongnv/gokit/util/server"
)

func main() {
	opts := []gokitServer.Option{
		gokitServer.WithHTTPAddress(":8080"),
	}

	mainHandler := handlers.New()
	opts = append(opts, service.GetOptions(mainHandler)...)
	opts = append(opts, expenseservice.GetOptions(mainHandler)...)

	err := gokitServer.Serve(
		opts...,
	)

	log.Println("Service stopped with:", err)
}
