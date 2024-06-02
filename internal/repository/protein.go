package repository

import (
	"errors"
	"fmt"
	"sort"

	"github.com/matsuboshi/ramengo/internal/model"
)

var proteinMap = map[string]model.Protein{
	"1": {
		ID:            "1",
		ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
		Name:          "CHASHU",
		Description:   "You couldn't choose any better.",
		Price:         13,
	},
	"2": {
		ID:            "2",
		ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
		Name:          "KAMABOKO",
		Description:   "It's fish cake. You'll love it!",
		Price:         10,
	},
	"3": {
		ID:            "3",
		ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
		ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
		Name:          "TOFU",
		Description:   "It's fine! We know you are broke.",
		Price:         6,
	},
}

func AllProteins() ([]model.Protein, error) {
	if len(proteinMap) == 0 {
		return nil, errors.New("could not get proteins")
	}

	proteins := []model.Protein{}

	for _, protein := range proteinMap {
		proteins = append(proteins, protein)
	}

	sort.Slice(proteins, func(i, j int) bool {
		return proteins[i].Price < proteins[j].Price
	})

	return proteins, nil
}

func ProteinNameById(proteinId string) (string, error) {
	protein, ok := proteinMap[proteinId]
	if !ok {
		return "", errors.New(fmt.Sprint("could not find protein id ", proteinId))
	}

	return protein.Name, nil
}
