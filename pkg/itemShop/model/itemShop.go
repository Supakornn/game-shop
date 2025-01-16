package model

type (
	Item struct {
		ID          uint64  `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Picture     string  `json:"picture"`
		Price       float64 `json:"price"`
	}
)
