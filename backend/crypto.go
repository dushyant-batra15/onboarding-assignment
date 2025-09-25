package main

import (
	"encoding/json"
	"time"
)

type Quote struct {
	Price float64 `json:"price"`
}

type cryptoCurrencies struct {
	Name   string           `json:"name"`
	Quotes map[string]Quote `json:"quotes"`
}

type cache struct {
	currencies []cryptoCurrencies
	ttl        time.Time
	provider   Provider
}

func (c *cache) getOrFetchCryptoCurrencies(provider Provider, cacheTtl int) ([]cryptoCurrencies, error) {
	if provider.getName() == c.provider.getName() && time.Now().Before(c.ttl) {
		return c.currencies, nil
	}

	cryptos, err := fetchCryptoCurrencies(provider)
	if err != nil {
		return nil, err
	}
	(*c).currencies = cryptos
	(*c).ttl = time.Now().Add(time.Duration(cacheTtl) * time.Second)
	(*c).provider = provider

	return cryptos, nil
}

func fetchCryptoCurrencies(p Provider) ([]cryptoCurrencies, error) {
	/*
		This data always expects a array in this format.
		[
			{
				"name": "BitCoin",
				"quotes": {
					"USD": {"price": 1.3},
					"INR": {"price": 1.2}
				}
			}
		]
	*/
	data, err := p.fetchCryptoData()

	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Step 2: Unmarshal into typed struct
	var cryptos []cryptoCurrencies
	if err := json.Unmarshal(jsonBytes, &cryptos); err != nil {
		return nil, err
	}
	return cryptos, nil
}
