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
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&input)

	if err != nil {
		fmt.Fprintf(writer, "Error when decoding input: %v", err)
		return
	}

	variant, ok, err := methods.GetProduct(input.SKU)

	if !ok {
		fmt.Fprintf(writer, "Error when searching product: %v", err)
		return
	}

	if variant.ID == 0 {
		fmt.Fprintf(writer, "Error: Product does not exist")
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
		fmt.Fprintf(writer, "Error when updating prices: %v", err)
		return
	}

	fmt.Fprintf(writer, "OK")
}
