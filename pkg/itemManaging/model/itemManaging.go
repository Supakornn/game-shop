package model

type (
	ItemCreatingReq struct {
		AdminID     string
		Name        string  `json:"name" validate:"required,max=64"`
		Description string  `json:"description" valdiate:"required,max=128"`
		Picture     string  `json:"picture" validate:"required"`
		Price       float64 `json:"price" validate:"required"`
	}

	ItemEditingReq struct {
		AdminId     string
		Name        string  `json:"name" validate:"omitempty,max=64"`
		Description string  `json:"description" validate:"omitempty,max=128"`
		Picture     string  `json:"picture" validate:"omitempty"`
		Price       float64 `json:"price" validate:"omitempty"`
	}
)
