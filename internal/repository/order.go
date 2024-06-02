package repository

import (
	"fmt"
	"sync"

	"github.com/matsuboshi/ramengo/internal/client"
	customError "github.com/matsuboshi/ramengo/internal/error"
	"github.com/matsuboshi/ramengo/internal/model"
)

func CreateOrder(secretKey, brothOption, proteinOption string) (model.Order, error) {
	brothName := ""
	proteinName := ""
	newOrderId := ""
	sharedError := &customError.SharedError{}

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
		newId, err := client.GenerateOrderID(secretKey)
		newOrderId = newId
		if err != nil {
			sharedError.Update(err)
		}
	}()

	wg.Wait()

	err := sharedError.Read()
	if err != nil {
		return model.Order{}, err
	}

	description := fmt.Sprintf("%s and %s Ramen", brothName, proteinName)

	order := model.Order{
		newOrderId,
		description,
		"https://tech.redventures.com.br/icons/ramen/ramenChasu.png",
	}

	fmt.Printf("New order %s: %s\n", order.ID, order.Description)

	return order, nil
}
