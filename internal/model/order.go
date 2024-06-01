package model

import (
	"fmt"
	"sync"
)

type Order struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func CreateOrder(secretKey, brothOption, proteinOption string) (Order, error) {
	brothChan := make(chan string, 1)
	proteinChan := make(chan string, 1)
	orderIDChan := make(chan string, 1)
	errChan := make(chan error, 3)
	defer close(brothChan)
	defer close(proteinChan)
	defer close(orderIDChan)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		name, err := BrothNameById(brothOption)
		errChan <- err
		brothChan <- name
	}()

	go func() {
		defer wg.Done()
		name, err := ProteinNameById(proteinOption)
		errChan <- err
		proteinChan <- name
	}()

	go func() {
		defer wg.Done()
		newId, err := GenerateOrderID(secretKey)
		errChan <- err
		orderIDChan <- newId
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return Order{}, err
		}
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
