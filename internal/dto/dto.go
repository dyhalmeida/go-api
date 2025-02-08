package dto

type ProductInputDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UserInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
