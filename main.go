package main

import (
	"net/http"

	"github.com/IsaGodoy/go-eCommerce-API/handlers"
	"github.com/IsaGodoy/go-eCommerce-API/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server := server.NewServer(":3000")
	server.AddRoute(http.MethodGet, "/", handlers.HandleHome)
	server.AddRoute(http.MethodPut, "/product", handlers.ChangePrices)
	server.Listen()
}
