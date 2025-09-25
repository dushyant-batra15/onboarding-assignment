package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var cacheData cache

var providers = map[string]Provider{
	"coinPaprika": coinPaprika{name: "coinPaprika", baseUrl: "https://api.coinpaprika.com/v1/tickers?quotes=USD,EUR,INR"},
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
		os.Exit(1)
	}

	defaultProvider := providers[os.Getenv("PROVIDER_NAME")]
	cryptos, err := fetchCryptoCurrencies(defaultProvider)
	if err != nil {
		fmt.Println("Issue in fetching data", err)
		os.Exit(1)
	}
	ttl, err := strconv.Atoi(os.Getenv("TTL_SECONDS"))
	if err != nil {
		fmt.Println("TTL in env is not Integer")
		os.Exit(1)
	}

	cacheData = cache{currencies: cryptos, ttl: time.Now().Add(time.Duration(ttl) * time.Second), provider: defaultProvider}
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(Logger())

	r.GET("/getAllCurrencies", getAllCurrencies)
	r.GET("/getCurrency", getCurrencyByName)
	r.Run(":9090")
}
