package utils

import (
	"context"
	"errors"
)

//||------------------------------------------------------------------------------------------------||
//|| Convert to USD (simple number only)
//||------------------------------------------------------------------------------------------------||

func ConvertToUSD(amount int64, currency string) (float64, error) {
	if amount < 0 {
		return 0, errors.New("invalid amount")
	}

	rate, _, err := getUSDRate(context.Background(), currency)
	if err != nil {
		return 0, err
	}

	dollars := float64(amount) * 100.0 // convert cents to dollars

	return dollars / rate, nil
}
