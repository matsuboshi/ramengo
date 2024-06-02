package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type IncomingID struct {
	OrderId string `json:"orderId"`
	Message string `json:"message"`
}

func GenerateOrderID(secretKey string) (string, error) {
	url := "https://api.tech.redventures.com.br/orders/generate-id"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return "", errors.New(fmt.Sprint("Error creating request: ", err))
	}

	req.Header.Set("x-api-key", secretKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(fmt.Sprint("could not fetch data from id generator: ", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("could not read response body from id generator")
	}

	var data IncomingID
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", errors.New("could not decode json")
	}

	if data.Message != "" {
		return "", errors.New(fmt.Sprint("Error generating new id: ", data.Message))
	}

	if data.OrderId == "" {
		return "", errors.New("id generator failed to return a new id")
	}

	return data.OrderId, nil
}
