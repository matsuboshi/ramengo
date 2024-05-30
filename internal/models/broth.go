package models

type Broth struct {
	ID            string  `json:"id"`
	ImageInactive string  `json:"imageInactive"`
	ImageActive   string  `json:"imageActive"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
}

func GetBroths() ([]Broth, error) {
	broths := []Broth{
		{
			ID:            "1",
			ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/salt/active.svg",
			Name:          "Salt",
			Description:   "Simple like the seawater, nothing more",
			Price:         10,
		},
		{
			ID:            "2",
			ImageInactive: "https://tech.redventures.com.br/icons/salt/inactive.svg",
			ImageActive:   "https://tech.redventures.com.br/icons/salt/active.svg",
			Name:          "Shoyu",
			Description:   "Go for it if you like an extra flavor",
			Price:         12,
		},
	}

	return broths, nil
}
