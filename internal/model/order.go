package model

import "fmt"

type Order struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func CreateOrder(secretKey, brothOption, proteinOption string) (Order, error) {
	brothChan := make(chan string)
	proteinChan := make(chan string)
	orderIDChan := make(chan string)
	errChan := make(chan error)
	defer close(errChan)

	go func() {
		defer close(brothChan)
		name, err := BrothNameById(brothOption)
		errChan <- err
		brothChan <- name
	}()

	go func() {
		defer close(proteinChan)
		name, err := ProteinNameById(proteinOption)
		errChan <- err
		proteinChan <- name
	}()

	go func() {
		defer close(orderIDChan)
		newId, err := GenerateOrderID(secretKey)
		errChan <- err
		orderIDChan <- newId
	}()

	var err error = nil
	for i := 0; i < 3; i++ {
		currentError := <-errChan
		if currentError != nil {
			err = currentError
		}
	}
	if err != nil {
		return Order{}, err
	}

	brothName := <-brothChan
	proteinName := <-proteinChan
	newOrderId := <-orderIDChan

	description := fmt.Sprintf("%s and %s Ramen", brothName, proteinName)

	order := Order{
		ID:          newOrderId,
		Description: description,
		Image:       "https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
	}

	fmt.Printf("New order %s: %s\n", order.ID, order.Description)

	return order, nil
}
