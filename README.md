# 🧪 Lab Supply API

API REST desenvolvida em Go para gerenciamento de uma distribuidora de produtos laboratoriais fictícia. O sistema permite o cadastro de produtos laboratoriais, clientes e pedidos, utilizando arquitetura em camadas, PostgreSQL e regras de negócio implementadas no domínio da aplicação.

Este projeto foi desenvolvido como desafio final do módulo 1 da disciplina de Go, com foco na aplicação dos principais conceitos da linguagem e de desenvolvimento Back-end.

---

# Tecnologias utilizadas

- Go
- PostgreSQL
- UUID
- bcrypt (criptografia de senhas)
- net/http
- database/sql

---

# Arquitetura

O projeto foi organizado utilizando arquitetura em camadas (Layered Architecture):

```
Controller
      ↓
Service
      ↓
Repository
      ↓
PostgreSQL
```

Cada camada possui responsabilidade única:

- **Controllers**: recebem e respondem às requisições HTTP.
- **Services**: implementam as regras de negócio.
- **Repositories**: realizam a comunicação com o banco de dados.
- **Domain**: contém as entidades e regras de domínio.

---

# Estrutura do projeto

```
cmd/
    app/

internal/
    config/
    controllers/
    database/
    domain/
    repository/
    routes/
    server/
    service/

migrations/
```

---

# Funcionalidades implementadas

## Produtos

- Cadastro de produtos
- Consulta de produtos por ID
- Listagem de produtos
- Atualização de produtos
- Remoção de produtos

---

## Clientes

- Cadastro de clientes
- Consulta de clientes por ID
- Listagem de clientes
- Atualização de clientes
- Remoção de clientes
- Criptografia de senha utilizando bcrypt

---

## Pedidos

- Criação de pedidos
- Consulta de pedido por ID
- Listagem paginada de pedidos
- Pagamento de pedidos
- Cancelamento de Pedidos
- Atualizçaão automática de estoque na criação do pedido
- Restauração automática de estoque na criação do pedido
- Persistência dos itens do pedido

---

# Regras de negócio implementadas

### Produtos

- Controle de estoque.
- Não permite estoque negativo.

### Clientes

- CPF único.
- Senhas armazenadas utilizando hash bcrypt.

### Pedidos

- Um pedido deve possuir um cliente.
- Um pedido deve possuir ao menos um item.
- Não permite quantidade menor ou igual a zero.
- Não permite pedidos com estoque insuficiente.
- Apenas pedidos com status **PENDING** podem ser pagos.
- Apenas pedidos com status **PENDING** podem ser cancelados.
- Cancelamento devolve automaticamente os produtos ao estoque.

---

# Endpoints

| Método | Endpoint |
|---------|----------|
| POST | `/products` |
| GET | `/products` |
| GET | `/products/{id}` |
| PUT | `/products/{id}` |
| DELETE | `/products/{id}` |

---

### Clientes

| Método | Endpoint |
|---------|----------|
| POST | `/customers` |
| GET | `/customers` |
| GET | `/customers/{id}` |
| PUT | `/customers/{id}` |
| DELETE | `/customers/{id}` |

---

### Pedidos

| Método | Endpoint |
|---------|----------|
| POST | `/orders` |
| GET | `/orders` |
| GET | `/orders/{id}` |
| POST | `/orders/{id}/pay` |
| POST | `/orders/{id}/cancel` 

---

## Fluxo de criação de pedidos

```
Cliente
      │
      ▼
POST /orders
      │
      ▼
OrderController
      │
      ▼
OrderService
      │
      ├── Busca Cliente
      ├── Busca Produtos
      ├── Valida Estoque
      ├── Cria Pedido
      ▼
OrderRepository
      │
      ├── INSERT orders
      ├── INSERT order_items
      ├── UPDATE products
      ▼
COMMIT
```

# Próximas evoluções (Versão 2.0)

O projeto foi estruturado para permitir evolução contínua. Algumas funcionalidades planejadas para versões futuras incluem:

- Autenticação JWT.
- Middleware de autorização.
- Swagger/OpenAPI.
- Docker e Docker Compose.
- Testes unitários.
- Testes de integração.
- Logging estruturado.
- GitHub Actions (CI/CD).
- Migrations automatizadas.
- Paginação genérica.
- Filtros nas consultas.
- Soft Delete.
---

## Conceitos aplicados

Durante o desenvolvimento foram utilizados conceitos de:

- Arquitetura em camadas
- Repository Pattern
- Injeção de Dependências
- DTOs
- APIs REST
- UUID
- PostgreSQL
- Transações
- Criptografia de senhas
- Tratamento de erros
- Modelagem de domínio
---

## Autora

**Ludimila Araújo**


