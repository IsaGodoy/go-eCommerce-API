package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IsaGodoy/go-eCommerce-API/methods"
	"github.com/IsaGodoy/go-eCommerce-API/types"
)

//Funciones
func HandleHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to my API!")
}

func ChangePrices(writer http.ResponseWriter, request *http.Request) {
	var input types.VariantPut
	var message string
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&input)
	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		message = fmt.Sprintf(`{"message": "Error when decoding input: %v"}`, err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(message))
		return
	}

	variant, ok, err := methods.GetProduct(input.SKU)

	if !ok {
		message = fmt.Sprintf(`{"message": "Error when searching product: %v"}`, err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(message))
		return
	}

	if variant.ID == 0 {
		message = `{"message": "Product does not exist"}`
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(message))
		return
	}

	var product types.ProductPut
	product.Variant = types.VariantPut{
		ID:               variant.ID,
		SKU:              variant.SKU,
		Price:            input.Price,
		Compare_at_price: input.Compare_at_price,
	}

	ok, err = methods.UpdatePrices(product)

	if !ok {
		message = fmt.Sprintf(`{"message": "Error when updating prices: %v"}`, err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(message))
		return
	}

	message = `{"message": "OK"}`
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(message))
}
