# 🧪 Lab Supply API

API REST desenvolvida em Go para gerenciamento de uma distribuidora de produtos laboratoriais. O sistema contempla o gerenciamento de produtos, clientes e pedidos, utilizando arquitetura em camadas, PostgreSQL e regras de negócio implementadas no domínio da aplicação

Este projeto foi desenvolvido como desafio final dos módulos 1 e 2 da disciplina de Go, com foco na aplicação dos principais conceitos da linguagem e de desenvolvimento Back-end.

---

## Tecnologias utilizadas

- Go
- PostgreSQL
- UUID
- bcrypt (criptografia de senhas)
- net/http
- database/sql

---

## Arquitetura

O projeto foi estruturado utilizando Layered Architecture, separando as responsabilidades entre apresentação, regras de negócio, persistência e domínio.

```
HTTP Request
      │
      ▼
Controller
      │
      ▼
DTO
      │
      ▼
Service
      │
      ▼
Repository
      │
      ▼
PostgreSQL
```

### Responsabilidades

**Controllers**

- Recebem as requisições HTTP
- Validam entrada
- Convertem DTOs
- Retornam respostas HTTP

**Services**

- Implementam as regras de negócio
- Coordenam os casos de uso da aplicação

**Repositories**

- Responsáveis pela persistência dos dados
- Encapsulam o acesso ao banco de dados

**Domain**

- Entidades
- Validações
- Estados
- Regras de negócio
- Erros centralizados

---

## Estrutura do projeto

```
cmd/
    app/
    seed/

internal/
    config/
    controllers/
        dto/
    database/
    domain/
    repository/
    routes/
    server/
    service/

migrations/
```

Os testes unitários permanecem junto aos respectivos pacotes seguindo a convenção da linguagem Go.

---

## Funcionalidades implementadas

### Produtos

- Cadastro de produtos
- Consulta de produtos por ID
- Listagem de produtos
- Atualização de produtos
- Remoção de produtos

---

### Clientes

- Cadastro de clientes
- Consulta de clientes por ID
- Listagem de clientes
- Atualização de clientes
- Remoção de clientes
- Criptografia de senha utilizando bcrypt

---

### Pedidos

- Criação de pedidos
- Consulta de pedido por ID
- Listagem paginada de pedidos
- Pagamento de pedidos
- Cancelamento de pedidos
- Atualização automática de estoque na criação do pedido
- Restauração automática de estoque na criação do pedido
- Persistência dos itens do pedido

---

## Regras de negócio implementadas

### Produtos

- Nome obrigatório
- Descrição obrigatória
- Marca obrigatória
- Preço maior que zero
- Estoque não pode ser negativo

---

### Clientes

- Nome obrigatório
- CPF obrigatório
- CPF deve possuir 11 dígitos
- Idade entre 18 e 120 anos
- Endereço obrigatório
- Email obrigatório
- Telefone obrigatório
- CPF único
- Senha armazenada utilizando hash bcrypt

---

### Pedidos

- Cliente obrigatório
- Pedido deve possuir ao menos um item
- Produto obrigatório
- Quantidade maior que zero
- Não permite estoque insuficiente
- Apenas pedidos com status **PENDING** podem ser modificados
- Apenas pedidos pendentes podem ser pagos
- Apenas pedidos pendentes podem ser cancelados
- Cancelamento restaura automaticamente o estoque

---

## Endpoints

### Produtos

| Método | Endpoint |
|---------|----------|
| POST | `/products` |
| GET | `/products` |
| GET | `/products/{id}` |
| PUT | `/products/{id}` |
| DELETE | `/products/{id}` |

---

## Clientes

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
| POST | `/orders/{id}/cancel` | 

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

## Testes

O projeto possui testes unitários para as principais regras de negócio implementadas.

Atualmente o projeto possui testes unitários para:

- Domain
- ProductService
- CustomerService
- OrderService

Os testes contemplam cenários de:

- sucesso
- validações
- erros
- regras de negócio
- transições de estado dos pedidos

---

## Próximas evoluções (Versão 3.0)

A estrutura do projeto foi planejada para permitir evolução contínua.

Algumas melhorias previstas para uma versão futura:

- Autenticação JWT
- Middleware de autorização
- Swagger/OpenAPI
- Docker e Docker Compose
- Ampliação da cobertura dos testes
- Testes de integração
- Testes End-to-End
- Logging estruturado
- GitHub Actions (CI/CD)
- Migrations automatizadas
- Paginação genérica
- Filtros nas consultas
- Soft Delete

---

## Conceitos aplicados

- Go
- APIs REST
- Arquitetura em camadas
- Repository Pattern
- Injeção de Dependências
- DTOs
- Modelagem de domínio
- Validação de entidades
- Tratamento centralizado de erros
- UUID
- PostgreSQL
- bcrypt
- Testes unitários
- Cobertura de código

---

## Autora

**Ludimila Araújo**

Desenvolvido como projeto final dos módulos 1 e 2 da formação em Go, com foco em arquitetura de software, APIs REST e testes unitários.