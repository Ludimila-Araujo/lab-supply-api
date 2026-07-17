package seed

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

const customerExistsQuery = `
SELECT EXISTS(
	SELECT 1
	FROM customers
	WHERE cpf = $1
)
`

func hashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

type seedCustomer struct {
	Name      string
	CPF       string
	BirthDate string
	Address   string
	Email     string
	Phone     string
	Password  string
}

var customers = []seedCustomer{
	{
		Name:      "Ana Carolina Lima",
		CPF:       "52998224725",
		BirthDate: "1992-04-10",
		Address:   "Rua das Flores, 120",
		Email:     "ana.lima@email.com",
		Phone:     "83999990001",
		Password:  "123456",
	},
	{
		Name:      "Bruno Henrique Alves",
		CPF:       "12345678909",
		BirthDate: "1989-09-15",
		Address:   "Av. Epitácio Pessoa, 300",
		Email:     "bruno.alves@email.com",
		Phone:     "83999990002",
		Password:  "123456",
	},
	{
		Name:      "Camila Fernandes Rocha",
		CPF:       "39053344705",
		BirthDate: "1994-02-18",
		Address:   "Rua João Pessoa, 88",
		Email:     "camila.rocha@email.com",
		Phone:     "83999990003",
		Password:  "123456",
	},
	{
		Name:      "Daniel Oliveira Costa",
		CPF:       "16899535009",
		BirthDate: "1988-12-02",
		Address:   "Rua das Acácias, 55",
		Email:     "daniel.costa@email.com",
		Phone:     "83999990004",
		Password:  "123456",
	},
	{
		Name:      "Fernanda Ribeiro Santos",
		CPF:       "11144477735",
		BirthDate: "1991-06-08",
		Address:   "Rua dos Ipês, 77",
		Email:     "fernanda@email.com",
		Phone:     "83999990005",
		Password:  "123456",
	},
	{
		Name:      "Gabriel Martins Silva",
		CPF:       "86288366757",
		BirthDate: "1990-01-20",
		Address:   "Rua Amazonas, 41",
		Email:     "gabriel@email.com",
		Phone:     "83999990006",
		Password:  "123456",
	},
	{
		Name:      "Juliana Pereira Rocha",
		CPF:       "15350946056",
		BirthDate: "1996-08-14",
		Address:   "Rua Ceará, 17",
		Email:     "juliana@email.com",
		Phone:     "83999990007",
		Password:  "123456",
	},
	{
		Name:      "Lucas Almeida Gomes",
		CPF:       "98765432100",
		BirthDate: "1993-11-09",
		Address:   "Rua Bahia, 205",
		Email:     "lucas@email.com",
		Phone:     "83999990008",
		Password:  "123456",
	},
	{
		Name:      "Mariana Carvalho Dias",
		CPF:       "71460238001",
		BirthDate: "1995-07-01",
		Address:   "Rua Pernambuco, 61",
		Email:     "mariana@email.com",
		Phone:     "83999990009",
		Password:  "123456",
	},
	{
		Name:      "Rafael Barbosa Melo",
		CPF:       "74697131401",
		BirthDate: "1987-05-30",
		Address:   "Rua Paraíba, 99",
		Email:     "rafael@email.com",
		Phone:     "83999990010",
		Password:  "123456",
	},
}

func SeedCustomers(db *sql.DB) error {

	customerRepository := repository.NewPostgresCustomerRepository(db)

	created := 0
	skipped := 0

	for _, c := range customers {

		var exists bool

		err := db.QueryRow(
			customerExistsQuery,
			c.CPF,
		).Scan(&exists)

		if err != nil {
			return err
		}

		if exists {
			skipped++
			continue
		}

		birthDate, err := time.Parse(
			"2006-01-02",
			c.BirthDate,
		)

		if err != nil {
			return err
		}

		passwordHash, err := hashPassword(c.Password)
		if err != nil {
			return err
		}

		customer, err := domain.NewCustomer(
			c.Name,
			c.CPF,
			birthDate,
			c.Address,
			c.Email,
			c.Phone,
			passwordHash,
		)

		if err != nil {
			return err
		}

		if err := customerRepository.Create(customer); err != nil {
			return err
		}

		created++
	}

	fmt.Printf(
		"Customers: %d created | %d skipped\n",
		created,
		skipped,
	)

	return nil
}
