package models

import "fmt"

type Order struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func CreateOrder(secretKey, brothOption, proteinOption string) (Order, error) {
	brothName, err := GetBrothName(brothOption)
	if err != nil {
		return Order{}, err
	}

	proteinName, err := GetProteinName(proteinOption)
	if err != nil {
		return Order{}, err
	}

	newOrderId, err := GenerateOrderID(secretKey)
	if err != nil {
		return Order{}, err
	}

	description := fmt.Sprintf("%s and %s Ramen", brothName, proteinName)

	order := Order{
		ID:          newOrderId,
		Description: description,
		Image:       "https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
	}

	return order, nil
}
