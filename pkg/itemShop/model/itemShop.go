package model

type (
	Item struct {
		ID          uint64  `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Picture     string  `json:"picture"`
		Price       float64 `json:"price"`
	}

	ItemFilter struct {
		Name        string `query:"name" validate:"omitempty,max=64"`
		Description string `query:"description" validate:"omitempty,max=128"`
	}
)
