package handlers

import (
	"fmt"
	"net/http"
)

//Funciones
func HandleHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to my API!")
}
