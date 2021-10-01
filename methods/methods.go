package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/IsaGodoy/go-eCommerce-API/types"
)

func GetProduct(SKU string) (types.Variant, bool, error) {

	baseURL := os.Getenv("API_BASE_URL")
	apiKey := os.Getenv("API_KEY")
	apiPassword := os.Getenv("API_PASSWORD")
	client := &http.Client{}
	urlRequest := ""
	var sinceID int64 = 0

	for sinceID >= 0 {
		//Set url for request to Shopify
		urlRequest = fmt.Sprint(baseURL, "/products.json?fields=id,variants&since_id=", sinceID)

		//Create internal request object
		req, _ := http.NewRequest(http.MethodGet, urlRequest, nil)

		//Set headers
		req.SetBasicAuth(apiKey, apiPassword)

		//Execute request
		resp, err := client.Do(req)

		if err != nil {
			return types.Variant{}, false, err
		}

		//Decoding response into struct type
		var _productRequest types.ProductRequest
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&_productRequest)

		if err != nil {
			return types.Variant{}, false, err
		}

		sinceID = -1

		for _, product := range _productRequest.Products {
			sinceID = product.ID

			for _, variant := range product.Variants {
				if variant.SKU == SKU {
					return variant, true, nil
				}
			}
		}
	}

	return types.Variant{}, true, nil
}

func UpdatePrices(product types.ProductPut) (bool, error) {
	baseURL := os.Getenv("API_BASE_URL")
	apiKey := os.Getenv("API_KEY")
	apiPassword := os.Getenv("API_PASSWORD")
	client := &http.Client{}

	//Set url for request to Shopify
	urlRequest := fmt.Sprint(baseURL, "/variants/", product.Variant.ID, ".json")

	//Convert object to byte array for pass to request body.
	payload, _ := product.ToJson()

	//Create internal request object
	req, _ := http.NewRequest(http.MethodPut, urlRequest, bytes.NewBuffer(payload))

	//Set headers
	req.SetBasicAuth(apiKey, apiPassword)
	req.Header.Set("Content-Type", "application/json")

	//Execute request
	_, err := client.Do(req)

	if err != nil {
		return false, err
	}

	return true, nil
}
