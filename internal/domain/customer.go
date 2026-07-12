package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// Customer representa um cliente da distribuidora de materiais laboratoriais.

type Customer struct {
	ID           uuid.UUID
	Name         string
	CPF          string
	BirthDate    time.Time
	Address      string
	Email        string
	Phone        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewCustomer(
	name string,
	cpf string,
	birthDate time.Time,
	address string,
	email string,
	phone string,
	passwordHash string,
) (*Customer, error) {

	if strings.TrimSpace(name) == "" {
		return nil, ErrCustomerNameRequired
	}

	if strings.TrimSpace(address) == "" {
		return nil, ErrCustomerAddressRequired
	}

	if strings.TrimSpace(email) == "" {
		return nil, ErrCustomerEmailRequired
	}

	if strings.TrimSpace(phone) == "" {
		return nil, ErrCustomerPhoneRequired
	}

	if strings.TrimSpace(passwordHash) == "" {
		return nil, ErrCustomerPasswordHashRequired
	}

	if err := validateCPF(cpf); err != nil {
		return nil, err

	}

	if err := validateBirthDate(birthDate); err != nil {
		return nil, err
	}

	now := time.Now()

	return &Customer{
		ID:           uuid.New(),
		Name:         name,
		CPF:          cpf,
		Address:      address,
		BirthDate:    birthDate,
		Email:        email,
		Phone:        phone,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}
