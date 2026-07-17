package seed

import (
	"database/sql"
	"fmt"

	"github.com/Ludimila-Araujo/lab-supply-api/internal/domain"
	"github.com/Ludimila-Araujo/lab-supply-api/internal/repository"
)

// Queries SQL

const productExistsQuery = `
SELECT EXISTS(
	SELECT 1
	FROM products
	WHERE name = $1
)
`

//Estrutura utilizada pelo SEED:

type seedProduct struct {
	Name        string
	Description string
	Brand       string
	Price       float64
	Stock       int
}

//Dados do SEED:

var products = []seedProduct{

	// Micropipetas

	{
		Name:        "Micropipeta Monocanal 10 µL",
		Description: "Micropipeta de alta precisão para pequenos volumes",
		Brand:       "Kasvi",
		Price:       1290.90,
		Stock:       8,
	},
	{
		Name:        "Micropipeta Monocanal 100 µL",
		Description: "Micropipeta de alta precisão",
		Brand:       "Kasvi",
		Price:       1385.50,
		Stock:       10,
	},
	{
		Name:        "Micropipeta Monocanal 1000 µL",
		Description: "Micropipeta para grandes volumes",
		Brand:       "Eppendorf",
		Price:       2490.00,
		Stock:       6,
	},
	{
		Name:        "Micropipeta Multicanal 300 µL",
		Description: "Micropipeta multicanal para placas",
		Brand:       "Eppendorf",
		Price:       4890.00,
		Stock:       4,
	},

	// Consumíveis

	{
		Name:        "Tubo Falcon 15 mL",
		Description: "Pacote com 50 unidades",
		Brand:       "TPP",
		Price:       58.90,
		Stock:       200,
	},
	{
		Name:        "Tubo Falcon 50 mL",
		Description: "Pacote com 25 unidades",
		Brand:       "TPP",
		Price:       74.50,
		Stock:       150,
	},
	{
		Name:        "Microtubo 1,5 mL",
		Description: "Pacote com 500 unidades",
		Brand:       "Axygen",
		Price:       89.90,
		Stock:       120,
	},
	{
		Name:        "Criotubo 2 mL",
		Description: "Pacote com 100 unidades",
		Brand:       "Kasvi",
		Price:       120.00,
		Stock:       80,
	},
	{
		Name:        "Ponteira 200 µL",
		Description: "Caixa com 1000 unidades",
		Brand:       "Kasvi",
		Price:       95.90,
		Stock:       250,
	},
	{
		Name:        "Ponteira 1000 µL",
		Description: "Caixa com 500 unidades",
		Brand:       "Kasvi",
		Price:       115.00,
		Stock:       180,
	},

	// Vidrarias

	{
		Name:        "Béquer 250 mL",
		Description: "Vidraria em borossilicato",
		Brand:       "Global Glass",
		Price:       32.90,
		Stock:       40,
	},
	{
		Name:        "Béquer 500 mL",
		Description: "Vidraria em borossilicato",
		Brand:       "Global Glass",
		Price:       48.90,
		Stock:       35,
	},
	{
		Name:        "Erlenmeyer 250 mL",
		Description: "Frasco Erlenmeyer graduado",
		Brand:       "Global Glass",
		Price:       42.00,
		Stock:       30,
	},
	{
		Name:        "Proveta Graduada 100 mL",
		Description: "Proveta em vidro graduado",
		Brand:       "Global Glass",
		Price:       39.50,
		Stock:       25,
	},

	// Equipamentos

	{
		Name:        "Centrífuga Digital 12 Tubos",
		Description: "Centrífuga digital para laboratório",
		Brand:       "Kasvi",
		Price:       8950.00,
		Stock:       3,
	},
	{
		Name:        "Vórtex Mixer",
		Description: "Misturador tipo vórtex",
		Brand:       "Biomixer",
		Price:       1490.00,
		Stock:       5,
	},
	{
		Name:        "Banho Maria Digital",
		Description: "Banho Maria com controle digital",
		Brand:       "SolidSteel",
		Price:       2890.00,
		Stock:       2,
	},
	{
		Name:        "Agitador Magnético",
		Description: "Agitador magnético com aquecimento",
		Brand:       "Fisatom",
		Price:       2190.00,
		Stock:       4,
	},

	// Reagentes

	{
		Name:        "PBS 10X 500 mL",
		Description: "Solução tampão fosfato",
		Brand:       "Sigma",
		Price:       185.00,
		Stock:       40,
	},
	{
		Name:        "Agarose Molecular Biology",
		Description: "Agarose para eletroforese",
		Brand:       "Invitrogen",
		Price:       420.00,
		Stock:       25,
	},
	{
		Name:        "Master Mix PCR 2X",
		Description: "Master Mix para PCR",
		Brand:       "Thermo Fisher",
		Price:       980.00,
		Stock:       18,
	},
	{
		Name:        "Taq DNA Polymerase",
		Description: "Enzima para PCR",
		Brand:       "Thermo Fisher",
		Price:       720.00,
		Stock:       15,
	},
	{
		Name:        "Kit ELISA IgG",
		Description: "Kit para ensaio ELISA",
		Brand:       "Bio-Rad",
		Price:       1850.00,
		Stock:       12,
	},
	{
		Name:        "Azul de Bromofenol",
		Description: "Corante para eletroforese",
		Brand:       "Sigma",
		Price:       95.00,
		Stock:       30,
	},
}

//Função do SEED:

func SeedProducts(db *sql.DB) error {

	created := 0
	skipped := 0

	productRepository := repository.NewPostgresProductRepository(db)

	for _, p := range products {

		var exists bool //verificação

		err := db.QueryRow(
			productExistsQuery,
			p.Name,
		).Scan(&exists)

		if err != nil {
			return err
		}

		if exists {
			skipped++
			continue
		}

		product, err := domain.NewProduct(
			p.Name,
			p.Description,
			p.Brand,
			p.Price,
			p.Stock,
		)

		if err != nil {
			return err
		}

		if err := productRepository.Create(product); err != nil {
			return err
		}

		created++
	}

	fmt.Printf(
		"Products: %d created | %d skipped\n",
		created,
		skipped,
	)

	return nil
}
