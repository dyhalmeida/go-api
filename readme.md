# Go Products API

Este projeto é uma API (CRUD de Produtos) desenvolvida em **Go** utilizando **Chi Router**, **JWT para autenticação** e **banco de dados relacional**.

## 🚀 Tecnologias Utilizadas

- **Golang**
- **Chi Router** (roteamento)
- **JWT** (autenticação)
- **Banco de Dados Relacional** (SQLite)

## 📂 Estrutura do Projeto

```
    ├── api/
    │   └── .gitkeep
    ├── cmd/
    │   ├── main.go
    │   ├── .env.example
    │   └── .gitkeep
    ├── configs/
    │   ├── config.go
    │   └── .gitkeep
    ├── docs/
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── internal/
    │   ├── .gitkeep
    │   ├── dto/
    │   │   └── dto.go
    │   ├── entity/
    │   │   ├── product.go
    │   │   ├── product_test.go
    │   │   ├── user.go
    │   │   └── user_test.go
    │   └── infra/
    │       ├── database/
    │       │   ├── product_db.go
    │       │   ├── product_db_test.go
    │       │   ├── user_db.go
    │       │   └── user_db_test.go
    │       └── websever/
    │           └── handlers/
    │               ├── product_handlers.go
    │               └── user_handlers.go
    ├── pkg/
    │   ├── .gitkeep
    │   └── entity/
    │       └── id.go
    └── test/
        ├── product.http
        ├── token.http
        ├── user.http
        └── .gitkeep

```  

## 🔧 Configuração do Ambiente

1. Clone o repositório:
   ```bash
   git clone https://github.com/dyhalmeida/go-apis.git
   ```
2. Entre no diretório do projeto:
   ```bash
   cd go-apis
   ```
3. Instale as dependências:
   ```bash
   go mod tidy
   ```

4. Configure o arquivo `.env` com as variáveis de ambiente necessárias:
   ```bash
  cp cmd/.env.example cmd/.env
   ```
   ```env
   DB_HOST=localhost
   DB_USER=root
   DB_PASSWORD=senha
   DB_NAME=goapis
   JWT_SECRET=sua-chave-secreta
   JWT_EXPIRES_IN=3600
   ```
4. Execute o servidor:
   ```bash
   cd cmd 
   go run main.go
   ```

## 🛠 Endpoints Disponíveis

### Usuários
- **Criar Usuário:** `POST /users`
- **Obter Token:** `POST /users/token`

### Produtos
- **Criar Produto:** `POST /products`
- **Buscar Produto por ID:** `GET /products/{id}`
- **Atualizar Produto:** `PUT /products/{id}`
- **Deletar Produto:** `DELETE /products/{id}`
- **Listar Produtos:** `GET /products?page=1&limit=3&sort=asc`

## 🔥 Testando a API

### Via Arquivos `.http` (VS Code REST Client)

Utilize os arquivos de testes localizados na pasta `test/`:
```bash
code test/user.http
code test/token.http
code test/product.http
```

Ou execute as requisições via **cURL**:
```bash
curl -X POST http://localhost:3333/users -H "Content-Type: application/json" -d '{"name":"Diego", "email":"diego@mail.com", "password":"123456"}'
```

### ✨ Autor: [Diego Almeida](https://github.com/dyhalmeida)