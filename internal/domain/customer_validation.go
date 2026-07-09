package domain

import (
	"strings"
	"time"
)

// funções de validação para cadastro de clientes

func validateCPF(cpf string) error {
	cpf = strings.TrimSpace(cpf)

	if cpf == "" {
		return ErrCustomerCpfRequired
	}

	if len(cpf) != 11 {
		return ErrCustomerCpfInvalid
	}

	return nil
}

func validateBirthDate(birthDate time.Time) error {

	now := time.Now()

	if birthDate.IsZero() {
		return ErrCustomerBirthDateRequired
	}

	if birthDate.After(now) {
		return ErrCustomerBirthDateInvalid
	}

	age := calculateAge(birthDate)

	if age < 18 {
		return ErrCustomerUnderAge
	}

	if age > 120 {
		return ErrCustomerOverAge
	}

	return nil
}

func calculateAge(birthDate time.Time) int {

	now := time.Now()

	age := now.Year() - birthDate.Year()

	if now.Month() < birthDate.Month() ||
		(now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}

	return age
}
