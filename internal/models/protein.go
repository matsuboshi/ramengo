package models

type Protein struct {
	ID            string  `json:"id"`
	ImageInactive string  `json:"imageInactive"`
	ImageActive   string  `json:"imageActive"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
}

func GetProteins() ([]Protein, error) {
	proteins := []Protein{
		{
			ID:            "1",
			ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
			Name:          "Chasu",
			Description:   "A sliced flavourful pork meat with a selection of season vegetables.",
			Price:         10,
		},
		{
			ID:            "2",
			ImageInactive: "https://tech.redventures.com.br/icons/pork/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/pork/active.svg",
			Name:          "Tofu",
			Description:   "We know you are vegetarian, or you might be broke. We got you covered.",
			Price:         6,
		},
	}

	return proteins, nil
}
