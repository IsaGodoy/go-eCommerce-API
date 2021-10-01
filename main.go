package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IsaGodoy/go-eCommerce-API/handlers"
	"github.com/IsaGodoy/go-eCommerce-API/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	server := server.NewServer(fmt.Sprint(":", port))
	server.AddRoute(http.MethodGet, "/", handlers.HandleHome)
	server.AddRoute(http.MethodPut, "/product", handlers.ChangePrices)
	server.Listen()
}
