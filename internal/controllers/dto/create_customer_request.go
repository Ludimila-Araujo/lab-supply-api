package dto

type CreateCustomerRequest struct {
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}
