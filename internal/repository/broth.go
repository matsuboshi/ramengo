package repository

import (
	"errors"
	"fmt"
	"sort"

	"github.com/matsuboshi/ramengo/internal/model"
)

var brothMap = map[string]model.Broth{
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

func AllBroths() ([]model.Broth, error) {
	if len(brothMap) == 0 {
		return nil, errors.New("could not get broths")
	}

	broths := make([]model.Broth, 0, len(brothMap))

	for _, broth := range brothMap {
		broths = append(broths, broth)
	}

	sort.Slice(broths, func(i, j int) bool {
		return broths[i].Price < broths[j].Price
	})

	return broths, nil
}

func BrothNameById(brothId string) (string, error) {
	broth, ok := brothMap[brothId]
	if !ok {
		return "", errors.New(fmt.Sprint("could not find broth id ", brothId))
	}

	return broth.Name, nil
}
