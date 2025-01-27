package main

import (
	"fmt"
	"go/crypto/api"
	"sync"
)

func main() {
	currencies := []string {"BTC", "ETH", "USDC"}

	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func() {
			getCurrencyRate(currency)
			wg.Done()
		}()
	}
	wg.Wait()
}

func getCurrencyRate(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}