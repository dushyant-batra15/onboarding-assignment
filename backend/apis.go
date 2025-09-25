package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getCryptosFromCache(c *gin.Context) ([]cryptoCurrencies, error) {
	providerName := c.DefaultQuery("provider", os.Getenv("PROVIDER_NAME"))
	provider := providers[providerName]
	if provider == nil {
		return nil, errors.New("Provider not found")
	}
	ttl, err := strconv.Atoi(os.Getenv("TTL_SECONDS"))
	if err != nil {
		return nil, errors.New("ENVIRONMENT ISSUE")
	}
	return cacheData.getOrFetchCryptoCurrencies(provider, ttl)
}

func getAllCurrencies(c *gin.Context) {
	cryptos, err := getCryptosFromCache(c)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	response := make(map[string]map[string]map[string]float64)
	data := make(map[string]map[string]float64)

	for _, crypto := range cryptos {
		prices := make(map[string]float64)
		for k, v := range crypto.Quotes {
			prices[k] = v.Price
		}
		data[crypto.Name] = prices
	}

	response["data"] = data
	c.JSON(http.StatusOK, response)
}

func getCurrencyByName(c *gin.Context) {
	name := c.DefaultQuery("name", "all")
	if name == "all" {
		getAllCurrencies(c)
		return
	}
	cryptos, err := getCryptosFromCache(c)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	for _, crypto := range cryptos {
		if crypto.Name == name {
			prices := make(map[string]float64)
			for k, v := range crypto.Quotes {
				prices[k] = v.Price
			}
			c.JSON(http.StatusFound, map[string]map[string]float64{crypto.Name: prices})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Currency with this name not found"})

}
