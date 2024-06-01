package model

import "fmt"

type Order struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func CreateOrder(secretKey, brothOption, proteinOption string) (Order, error) {
	brothName, err := BrothNameById(brothOption)
	if err != nil {
		return Order{}, err
	}

	proteinName, err := ProteinNameById(proteinOption)
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

	fmt.Printf("New order %s: %s\n", order.ID, order.Description)

	return order, nil
}
