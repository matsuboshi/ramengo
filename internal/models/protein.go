package models

import (
	"errors"
	"sort"
)

type protein struct {
	ID            string  `json:"id"`
	ImageInactive string  `json:"imageInactive"`
	ImageActive   string  `json:"imageActive"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
}

var proteinMap = map[int]protein{
	1: {
		ID:            "1",
		ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
		Name:          "CHASU",
		Description:   "You couldn't choose any better.",
		Price:         10,
	},
	2: {
		ID:            "2",
		ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
		Name:          "TOFU",
		Description:   "It's fine! We know you are broke.",
		Price:         6,
	},
}

func GetProteins() ([]protein, error) {
	if len(proteinMap) == 0 {
		return nil, errors.New("could not get proteins")
	}

	proteins := []protein{}

	for _, protein := range proteinMap {
		proteins = append(proteins, protein)
	}

	sort.Slice(proteins, func(i, j int) bool {
		return proteins[i].Price < proteins[j].Price
	})

	return proteins, nil
}
