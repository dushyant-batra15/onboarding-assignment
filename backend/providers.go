package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Provider interface {
	fetchCryptoData() ([]map[string]interface{}, error)
	getName() string
}

type coinPaprika struct {
	name    string
	baseUrl string
}

func (cp coinPaprika) fetchCryptoData() ([]map[string]interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, cp.baseUrl, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	httpClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return cp.renameFields(data), nil
}

func (cp coinPaprika) getName() string {
	return cp.name
}

func (cp coinPaprika) renameFields(respData []map[string]interface{}) []map[string]interface{} {
	//ADD THIS LOGIC IN ALL PROVIDERS NEEDED
	// result := make([]map[string]interface{}, 0, len(respData))

	// for _, item := range respData {
	// 	newItem := make(map[string]interface{})

	// 	for key, value := range item {
	// 		switch key {
	// 		case "name":
	// 			newItem["coin_name"] = value // rename "name" → "coin_name"
	// 		case "prices":
	// 			newItem["price_data"] = value // rename "prices" → "price_data"
	// 		default:
	// 			newItem[key] = value // keep other fields as is
	// 		}
	// 	}

	// 	result = append(result, newItem)
	// }

	// return result
	return respData
}
