package main

import (
	"fmt"
	"go/crypto/api"
)

func main() {}

func getCurrencyRate(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}