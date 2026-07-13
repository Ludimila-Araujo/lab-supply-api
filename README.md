# 🧪 Lab Supply API

API REST desenvolvida em Go para gerenciamento de uma distribuidora de produtos laboratoriais fictícia. O sistema permite o cadastro de produtos laboratoriais, clientes e pedidos, utilizando arquitetura em camadas, PostgreSQL e regras de negócio implementadas no domínio da aplicação.

Este projeto foi desenvolvido como desafio final do módulo 1 da disciplina de Go, com foco na aplicação dos principais conceitos da linguagem e de desenvolvimento Back-end.

---

# 🚀 Tecnologias utilizadas

- Go
- PostgreSQL
- UUID
- bcrypt (criptografia de senhas)
- net/http
- database/sql

---

# 🏗 Arquitetura

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

# 📂 Estrutura do projeto

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

# ✅ Funcionalidades implementadas

## Produtos

- Cadastro de produtos
- Consulta de produtos
- Atualização de estoque
- Persistência em PostgreSQL

---

## Clientes

- Cadastro de clientes
- Consulta de clientes
- Criptografia de senha utilizando bcrypt
- Persistência em PostgreSQL

---

## Pedidos

- Criação de pedidos
- Associação entre clientes e produtos
- Persistência do pedido
- Persistência dos itens do pedido
- Atualização automática do estoque
- Transações utilizando PostgreSQL (BEGIN / COMMIT / ROLLBACK)

---

# 🔄 Fluxo de criação de pedidos

```text
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

---

# 🛡 Regras de negócio implementadas

- Não permite criação de pedidos sem cliente.
- Não permite itens sem produto.
- Não permite quantidade menor ou igual a zero.
- Não permite pedidos com estoque insuficiente.
- Geração automática de UUID.
- Controle de datas de criação e atualização.
- Senhas armazenadas utilizando hash bcrypt.
- Atualização automática do estoque após criação do pedido.

---

# 📌 Endpoints implementados

## Produtos

- POST /products
- GET /products

---

## Clientes

- POST /customers
- GET /customers

---

## Pedidos

- POST /orders

---

# 📈 Próximas evoluções (Versão 2.0)

O projeto foi estruturado para permitir evolução contínua. Algumas funcionalidades planejadas para versões futuras incluem:

- Implementação completa do CRUD de pedidos.
- Paginação nas consultas.
- Endpoint para pagamento de pedidos.
- Endpoint para cancelamento de pedidos.
- Listagem detalhada de pedidos por cliente.
- Atualização parcial utilizando PATCH.
- Testes de integração.
- Documentação automática com Swagger/OpenAPI.
- Docker e Docker Compose.
- Migrations automatizadas.
- Pipeline de CI/CD utilizando GitHub Actions.

---

# 🎯 Objetivos de aprendizagem

Durante o desenvolvimento deste projeto foram aplicados conceitos como:

- Organização de projetos em Go.
- Arquitetura em camadas.
- Repository Pattern.
- Injeção de Dependências.
- Manipulação de HTTP com net/http.
- PostgreSQL.
- Transações.
- Criptografia de senhas.
- DTOs.
- Tratamento de erros.
- Modelagem de domínio.
- Desenvolvimento de APIs REST.

---

# 👩‍💻 Desenvolvido por

**Ludimila Araújo**


