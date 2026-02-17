# AMSVault - Guia para Agentes de IA

## Visão Geral do Projeto

AMSVault é uma API REST desenvolvida em **Go (Golang)** para gerenciar séries, animes e mangás. O sistema permite que usuários criem contas, adicionem histórias (stories) ao banco de dados e criem bookmarks para acompanhar seu progresso de visualização/leitura.

## Arquitetura e Padrões

### Arquitetura em Camadas

O projeto segue uma arquitetura limpa (Clean Architecture) organizada nas seguintes camadas:

```
┌─────────────────────────────────────────────┐
│            HTTP Layer (Server)              │
│        Echo Framework + Middlewares          │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│          Controller Layer                    │
│     Handlers + Request/Response ViewModels   │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│           Service Layer                      │
│       Business Logic + Validations           │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│         Data/Database Layer                  │
│      Repositories + Database Operations      │
└─────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────┐
│            Entity Layer                      │
│      Domain Models + Business Rules          │
└─────────────────────────────────────────────┘
```

### Estrutura de Diretórios

```
AMSVault/
├── main.go                     # Ponto de entrada da aplicação
├── go.mod                      # Dependências do projeto
├── config/                     # Configurações e variáveis de ambiente
│   └── config.go
├── constants/                  # Constantes da aplicação
│   └── constants.go
├── entity/                     # Entidades de domínio
│   ├── user.go                 # Entidade de usuário
│   ├── story.go                # Entidade de história
│   ├── bookmarks.go            # Entidade de bookmark
│   └── token.go                # Entidade de token JWT
├── controller/                 # Camada de controladores
│   ├── controller.go           # Struct principal do controller
│   ├── user.go                 # Handlers de usuário
│   ├── story.go                # Handlers de story
│   ├── bookmarks.go            # Handlers de bookmarks
│   ├── token.go                # Handlers de autenticação
│   └── viewmodel/              # DTOs de request/response
│       ├── request/
│       └── response/
├── service/                    # Camada de serviços (regras de negócio)
│   ├── service.go
│   ├── user.go
│   ├── story.go
│   ├── bookmarks.go
│   └── token.go
├── data/                       # Camada de acesso a dados
│   ├── data.go                 # Struct principal de dados
│   ├── connection.go           # Gerenciamento de conexões
│   ├── model/                  # Models de banco de dados
│   ├── mysql/                  # Implementação MySQL
│   └── mongo/                  # Implementação MongoDB
├── database/                   # Interfaces de repositórios
│   ├── user_db.go
│   └── story_db.go
├── server/                     # Configuração do servidor HTTP
│   ├── server.go
│   ├── middleware/             # Middlewares (autenticação, etc)
│   └── router/                 # Definição de rotas
├── integration/                # Integrações externas
│   └── my_anime_list.go        # Integração com MyAnimeList API
├── utils/                      # Utilitários
└── pkg/                        # Pacotes compartilhados
    ├── entity/
    └── errors/                 # Sistema de erros customizados
```

## Tecnologias e Dependências

### Stack Principal

- **Go**: 1.24.2
- **Web Framework**: Echo v4
- **Banco de Dados Relacional**: MySQL (via go-sql-driver/mysql)
- **Banco de Dados NoSQL**: MongoDB (via mongo-driver)
- **ORM**: GORM
- **Autenticação**: JWT (go-chi/jwtauth + golang-jwt/jwt/v4)
- **Criptografia**: bcrypt (golang.org/x/crypto)
- **Configuração**: Viper
- **UUID**: google/uuid

### Bancos de Dados

O projeto utiliza dois bancos de dados:

1. **MySQL**: Armazena usuários e stories
2. **MongoDB**: Armazena bookmarks (devido à natureza flexível dos dados)

## Fluxo de Inicialização

A aplicação segue este fluxo de inicialização (ver `main.go`):

```go
1. LoadConfig()      → Carrega variáveis de ambiente
2. data.New()        → Inicializa conexões com bancos de dados
3. NewIntegration()  → Configura integrações externas (MAL API)
4. NewService()      → Inicializa camada de serviços
5. NewController()   → Inicializa camada de controllers
6. server.Start()    → Inicia servidor HTTP
```

## Padrões de Código

### 1. Entidades (Entity Layer)

- Localização: `entity/`
- Contêm regras de negócio do domínio
- Construtores: Usam padrão `NewX()` que retorna `(*Entity, error)`
- Validação: Método `Validate() error`
- Exemplo:

```go
func NewUser(name, email, password string) (*User, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, errors.NewInternalError("NewUser", err)
    }
    return &User{
        Name:     name,
        Email:    email,
        Password: string(hash),
    }, nil
}
```

### 2. Controllers

- Localização: `controller/`
- Responsabilidades:
  - Receber requisições HTTP
  - Validar dados de entrada
  - Chamar serviços apropriados
  - Retornar respostas HTTP formatadas
- Usam ViewModels para request/response
- Exemplo de assinatura:

```go
func (c *Controller) CreateUser(ctx echo.Context) error
```

### 3. Services

- Localização: `service/`
- Contêm lógica de negócio
- Orquestram operações entre diferentes repositórios
- Validam regras de negócio complexas
- Retornam erros customizados

### 4. Data/Database Layer

- Localização: `data/` e `database/`
- `database/`: Define interfaces de repositórios
- `data/`: Implementações concretas (MySQL, MongoDB)
- Padrão Repository para abstrair acesso a dados

### 5. Sistema de Erros

- Localização: `pkg/errors/`
- Erros customizados com tipos:
  - `VALIDATION_ERROR`: Erros de validação
  - `NOT_FOUND`: Recurso não encontrado
  - `ALREADY_EXISTS`: Conflito de duplicação
  - `INTERNAL_ERROR`: Erros internos
  - `DATABASE_ERROR`: Erros de banco de dados
  - `EXTERNAL_SERVICE_ERROR`: Falhas em APIs externas

## Autenticação e Segurança

### JWT Token

- Gerado no login via `POST /login`
- Token contém claims com ID do usuário
- Expira após tempo configurado em `JWT_EXPIRATION_TIME`
- Middleware de autenticação em `server/middleware/auth.go`

### Proteção de Rotas

- Rotas públicas: `/login`, `POST /user`
- Rotas privadas: Requerem header `Authorization: Bearer {token}`

### Hash de Senhas

- Utiliza bcrypt com custo padrão
- Hash aplicado na criação do usuário (`NewUser`)
- Validação via `ValidatePassword()`

## Modelos de Dados

### User (MySQL)

```go
type User struct {
    ID        int64
    Name      string
    Email     string     // único
    Password  string     // bcrypt hash
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Story (MySQL)

```go
type Story struct {
    ID           int64
    MALID        int64      // MyAnimeList ID
    Name         string
    Source       string     // anime, manga, novel, etc
    Description  string
    TotalSeason  int64
    TotalEpisode int64
    TotalVolume  int64
    TotalChapter int64
    Status       string     // ongoing, completed, etc
    MainPicture  MainPicture
}
```

### Bookmarks (MongoDB)

```go
type Bookmarks struct {
    ID             string    // MongoDB ObjectID
    UserID         int64
    StoryID        int64
    Status         string    // watching, reading, completed, dropped
    CurrentSeason  int64
    CurrentEpisode int64
    CurrentVolume  int64
    CurrentChapter int64
    CreatedAt      time.Time
    UpdatedAt      time.Time
    DeletedAt      time.Time // soft delete
}
```

## Rotas Principais

### Autenticação
- `POST /login` - Login e geração de token

### Usuários
- `POST /user` - Criar usuário
- `GET /user` - Buscar usuário autenticado
- `GET /user/:id` - Buscar usuário por ID
- `PUT /user` - Atualizar usuário
- `DELETE /user/:id` - Deletar usuário

### Stories
- `POST /story` - Criar história
- `GET /story/:id` - Buscar história por ID
- `GET /story?name=X` - Buscar histórias por nome

### Bookmarks
- `POST /bookmarks` - Criar bookmark
- `GET /bookmarks/:id` - Buscar bookmark por ID
- `GET /bookmarks/user/:user_id` - Listar bookmarks do usuário
- `PUT /bookmarks/:id` - Atualizar bookmark
- `DELETE /bookmarks/:id` - Deletar bookmark

## Configuração (.env)

Variáveis necessárias:

```env
# MySQL
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
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
```

## Convenções de Código

### Nomenclatura

- **Packages**: lowercase, singular (ex: `entity`, `service`, `controller`)
- **Interfaces**: Sufixo `Interface` (ex: `UserDBInterface`)
- **Construtores**: Prefixo `New` (ex: `NewUser`, `NewService`)
- **Getters**: Sem prefixo `Get` (ex: `user.Name()` não `user.GetName()`)
- **Validação**: Método `Validate() error`

### Tratamento de Erros

- Sempre propagar erros para camadas superiores
- Usar erros customizados do pacote `pkg/errors`
- Log de erros na camada apropriada
- Retornar status HTTP adequados

### Organização de Imports

```go
import (
    // Bibliotecas padrão
    "fmt"
    "time"
    
    // Bibliotecas de terceiros
    "github.com/labstack/echo/v4"
    
    // Pacotes internos do projeto
    "github.com.br/GregoryLacerda/AMSVault/entity"
)
```

## Integração com MyAnimeList

- Localização: `integration/my_anime_list.go`
- Funcionalidades:
  - Buscar animes por ID
  - Buscar animes por nome
  - Autenticação OAuth2
  - Refresh de tokens

## Testes

(A ser implementado)

## Docker

- Localização: `build/docker-compose.yaml`
- Contém configuração para:
  - MySQL
  - MongoDB
  - (Potencialmente) API container

## Comandos Úteis

```bash
# Executar aplicação
go run main.go

# Build
go build -o amsvault

# Instalar dependências
go mod download

# Atualizar dependências
go mod tidy

# Executar com Docker
cd build
docker-compose up -d
```

## Quando Modificar Código

### Adicionar Nova Entidade

1. Criar struct em `entity/`
2. Implementar construtor `NewX()` e `Validate()`
3. Criar interface de repositório em `database/`
4. Implementar repositório em `data/mysql/` ou `data/mongo/`
5. Adicionar métodos no service (`service/`)
6. Criar controllers (`controller/`)
7. Adicionar rotas (`server/router/`)
8. Atualizar documentação

### Adicionar Nova Rota

1. Definir ViewModels de request/response em `controller/viewmodel/`
2. Criar handler no controller apropriado
3. Registrar rota em `server/router/`
4. Aplicar middleware de autenticação se necessário
5. Documentar em `API_DOCUMENTATION.md`

### Modificar Validações

1. Entidades: Atualizar método `Validate()` em `entity/`
2. Controllers: Atualizar validação de ViewModels
3. Services: Adicionar validações de regras de negócio
4. Atualizar constantes de erro em `constants/`

## Notas Importantes

1. **Soft Delete**: Bookmarks usam soft delete (campo `DeletedAt`)
2. **IDs**: Users e Stories usam `int64`, Bookmarks usam MongoDB ObjectID (string)
3. **Thread Safety**: Services devem ser thread-safe
4. **Transações**: Implementar quando necessário operações atômicas
5. **Logging**: Adicionar logs adequados para debugging
6. **Validação**: Sempre validar dados na entrada (controller) e domínio (entity)

## Referências

- [Documentação da API](API_DOCUMENTATION.md)
- [Guia de Implementação](implementation_guide.md)
- [Echo Framework](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
- [MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/)

---

**Última Atualização**: Fevereiro 2026
