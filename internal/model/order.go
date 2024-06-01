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

type SharedError struct {
	mu  sync.Mutex
	err error
}

func (e *SharedError) Update(err error) {
	e.mu.Lock()
	e.err = err
	e.mu.Unlock()
}

func (e *SharedError) Read() error {
	e.mu.Lock()
	err := e.err
	e.mu.Unlock()
	return err
}

func CreateOrder(secretKey, brothOption, proteinOption string) (Order, error) {
	brothName := ""
	proteinName := ""
	newOrderId := ""
	sharedError := &SharedError{}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		name, err := BrothNameById(brothOption)
		brothName = name
		if err != nil {
			sharedError.Update(err)
		}
	}()

	go func() {
		defer wg.Done()
		name, err := ProteinNameById(proteinOption)
		proteinName = name
		if err != nil {
			sharedError.Update(err)
		}
	}()

	go func() {
		defer wg.Done()
		newId, err := GenerateOrderID(secretKey)
		newOrderId = newId
		if err != nil {
			sharedError.Update(err)
		}
	}()

	wg.Wait()

	err := sharedError.Read()
	if err != nil {
		return Order{}, err
	}

	description := fmt.Sprintf("%s and %s Ramen", brothName, proteinName)

	order := Order{
		newOrderId,
		description,
		"https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
	}

	fmt.Printf("New order %s: %s\n", order.ID, order.Description)

	return order, nil
}
