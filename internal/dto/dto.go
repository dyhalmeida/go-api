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

type CredentialsInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CredentialsOutputDTO struct {
	AccessToken string `json:"access_token"`
}
