# AMSVault

API REST em Go para gerenciamento de histórias (animes, mangás, séries) com autenticação JWT e sistema de bookmarks para acompanhamento de progresso.

## 📖 Visão Geral

O AMSVault permite:
- criar e autenticar usuários
- cadastrar e consultar histórias
- criar, atualizar e remover bookmarks por usuário
- integrar dados de histórias via MyAnimeList

A aplicação segue arquitetura em camadas (`server -> controller -> service -> data/database -> entity`) para manter separação de responsabilidades e facilitar evolução.

## ✨ Funcionalidades

- Autenticação com JWT (`POST /login`)
- CRUD de usuários
- CRUD básico de stories (create + busca por id/nome)
- CRUD de bookmarks com persistência no MongoDB
- Integração com MyAnimeList para enriquecimento de dados
- Tratamento de erros com tipos customizados (`pkg/errors`)

## 🛠️ Tecnologias

- Go `1.24.2`
- Echo v4
- GORM
- MySQL (usuários e stories)
- MongoDB (bookmarks)
- JWT (`go-chi/jwtauth` + `golang-jwt/jwt/v4`)
- Viper (configuração por ambiente)

## 📋 Pré-requisitos

- Go `1.24+`
- Docker e Docker Compose (recomendado para bancos)
- MySQL e MongoDB (caso rode sem Docker)

## 🚀 Instalação e Execução

### 1) Subir bancos com Docker (recomendado)

```bash
cd build
docker-compose up -d
```

### 2) Configurar variáveis de ambiente

Crie o arquivo `.env` na raiz do projeto com os campos abaixo:

```env
# MySQL
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=amsvault
DB_PASSWORD=amsvaultPwd
DB_NAME=amsvault

# MongoDB
MONGO_URI=mongodb://localhost:27017
MONGO_DB=amsvault
MONGO_COLLECTION=bookmarks

# Servidor
WEB_SERVER_PORT=8080

# JWT
JWT_SECRET=seu_secret_aqui
JWT_EXPIRATION_TIME=86400

# MyAnimeList API
MAL_API_URL=https://api.myanimelist.net/v2
MAL_API_AUTH_URL=https://myanimelist.net/v1/oauth2/token
MAL_CLIENT_ID=seu_client_id
MAL_CLIENT_SECRET=seu_client_secret
MAL_GRANT_TYPE=refresh_token
MAL_REFRESH=seu_refresh_token
MAL_TOKEN=seu_access_token
```

### 3) Instalar dependências e executar API

```bash
go mod download
go run main.go
```

A API ficará disponível em: `http://localhost:8080`

## ⚙️ Configuração

A aplicação lê variáveis de ambiente via `viper` durante a inicialização (`config.LoadConfig`).

Fluxo de startup atual:
1. carregar config
2. abrir conexões de banco
3. inicializar integrações externas
4. inicializar services e controllers
5. subir servidor HTTP

## 🔐 Autenticação

- Endpoint de login: `POST /login`
- Envie credenciais (`email`, `password`) e receba token JWT
- Para rotas protegidas, use:

```http
Authorization: Bearer <token>
```

## 📚 Endpoints Principais

### Públicos
- `POST /login`
- `POST /user`

### Protegidos (JWT)
- `GET /user`
- `GET /user/:id`
- `PUT /user`
- `DELETE /user/:id`
- `POST /story`
- `GET /story/:id`
- `GET /story?name=<nome>`
- `POST /bookmarks`
- `GET /bookmarks/:id`
- `GET /bookmarks/user/:user_id`
- `PUT /bookmarks/:id`
- `DELETE /bookmarks/:id`

Documentação detalhada em: [API_DOCUMENTATION.md](API_DOCUMENTATION.md)

## 🧪 Testes

No momento, a base não possui suíte completa de testes automatizados.

Quando houver testes implementados, execute:

```bash
go test ./...
```

## 🧰 Comandos Úteis

```bash
# Build
go build -o amsvault

# Formatação
go fmt ./...

# Análise estática básica
go vet ./...

# Atualizar dependências
go mod tidy
```

## 📁 Estrutura do Projeto

```text
AMSVault/
├── main.go
├── config/
├── controller/
├── service/
├── data/
├── database/
├── entity/
├── integration/
├── server/
├── pkg/
└── build/
```

## 🤝 Contribuindo

1. Crie uma branch para sua alteração
2. Mantenha o padrão de arquitetura em camadas
3. Evite acoplamento entre HTTP e regras de negócio
4. Atualize documentação ao alterar contrato de API

## 📄 Licença

Este projeto não possui licença definida no repositório até o momento.
