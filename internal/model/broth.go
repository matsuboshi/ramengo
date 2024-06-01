package model

import (
	"errors"
	"fmt"
	"sort"
)

type broth struct {
	ID            string  `json:"id"`
	ImageInactive string  `json:"imageInactive"`
	ImageActive   string  `json:"imageActive"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
}

var brothMap = map[string]broth{
	"1": {
		ID:            "1",
		ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/salt/active.svg",
		Name:          "SHIO",
		Description:   "It's just salt, are you sure???",
		Price:         9,
	},
	"2": {
		ID:            "2",
		ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/salt/active.svg",
		Name:          "SHOYU",
		Description:   "Shoyu is the most popular broth in Japan.",
		Price:         13,
	},
	"3": {
		ID:            "3",
		ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/salt/active.svg",
		Name:          "TONKOTSU",
		Description:   "Rich and creamy pork broth. You'll love it!",
		Price:         15,
	},
}

func GetBroths() ([]broth, error) {
	if len(brothMap) == 0 {
		return nil, errors.New("could not get broths")
	}

	broths := []broth{}

	for _, broth := range brothMap {
		broths = append(broths, broth)
	}

	sort.Slice(broths, func(i, j int) bool {
		return broths[i].Price < broths[j].Price
	})

	return broths, nil
}

func GetBrothName(brothId string) (string, error) {
	broth, ok := brothMap[brothId]
	if !ok {
		return "", errors.New(fmt.Sprint("could not find broth id ", brothId))
	}

	return broth.Name, nil
}