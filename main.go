package main

import (
	"net/http"

	"github.com/IsaGodoy/go-eCommerce-API/handlers"
	"github.com/IsaGodoy/go-eCommerce-API/server"
)

func main() {
	server := server.NewServer(":3000")
	server.AddRoute(http.MethodGet, "/", handlers.HandleHome)
	server.Listen()
}
