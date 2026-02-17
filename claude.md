# AMSVault - Contexto para Claude

Este documento fornece contexto específico sobre o projeto AMSVault para auxiliar o Claude (e outros LLMs) a entender profundamente a aplicação e realizar modificações com precisão.

## O Que É Este Projeto?

AMSVault é uma **API REST em Go** para gerenciamento de histórias de entretenimento (animes, mangás, séries) com sistema de bookmarks para rastreamento de progresso. Pense nisso como um sistema de biblioteca pessoal com marcadores de leitura/visualização.

## Contexto de Arquitetura

### Filosofia do Projeto

- **Clean Architecture**: Separação clara entre camadas
- **Dependency Injection**: Dependências injetadas via construtores
- **Interface-based Design**: Abstrações para repositórios e serviços
- **Explicit Error Handling**: Go idiomático com propagação de erros

### Fluxo de Dados

```
Cliente HTTP → Echo Router → Middleware (Auth) → Controller → Service → Repository → Database
                                                      ↓
                                                   Entity
```

**Princípio chave**: Dados fluem de fora para dentro (HTTP → Entity), e validações ocorrem em múltiplas camadas.

## Regras de Validação

### Camada Controller
- Valida formato de request (JSON parsing)
- Verifica campos obrigatórios de ViewModels
- Valida tipos de dados básicos

### Camada Entity
- Valida regras de domínio (ex: valores negativos)
- Aplica business rules (ex: hash de senha)
- Garante consistência da entidade

### Camada Service
- Valida regras de negócio complexas
- Verifica existência de dependências
- Coordena operações entre múltiplas entidades

## Padrões de Implementação

### 1. Criação de Entidades

**SEMPRE** usar construtor `NewX()`:

```go
// ✅ CORRETO
user, err := entity.NewUser(name, email, password)
if err != nil {
    return nil, err
}

// ❌ INCORRETO
user := &entity.User{
    Name:     name,
    Email:    email,
    Password: password, // Sem hash!
}
```

### 2. Tratamento de Erros

**Usar erros customizados** do pacote `pkg/errors`:

```go
// ✅ CORRETO
if user == nil {
    return errors.NewNotFoundError("user not found")
}

// ❌ INCORRETO
if user == nil {
    return errors.New("user not found")
}
```

### 3. Controllers

**Estrutura padrão** de um handler:

```go
func (c *Controller) HandlerName(ctx echo.Context) error {
    // 1. Parse request
    var req RequestViewModel
    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(400, ErrorResponse{...})
    }
    
    // 2. Validar request
    if err := req.Validate(); err != nil {
        return ctx.JSON(400, ErrorResponse{...})
    }
    
    // 3. Chamar service
    result, err := c.service.SomeMethod(req)
    if err != nil {
        return handleError(ctx, err)
    }
    
    // 4. Retornar resposta
    return ctx.JSON(200, ResponseViewModel{...})
}
```

### 4. Services

**Responsabilidades**:
- Orquestrar operações
- Validar regras de negócio
- Não conhecer detalhes HTTP (não usar `echo.Context`)

```go
func (s *Service) CreateUser(name, email, password string) (*entity.User, error) {
    // 1. Verificar duplicação
    existing, _ := s.userDB.FindByEmail(email)
    if existing != nil {
        return nil, errors.NewAlreadyExistsError("user already exists")
    }
    
    // 2. Criar entidade
    user, err := entity.NewUser(name, email, password)
    if err != nil {
        return nil, err
    }
    
    // 3. Persistir
    return s.userDB.Create(user)
}
```

### 5. Repositories (Database Layer)

**Interfaces** em `database/`, **implementações** em `data/mysql/` ou `data/mongo/`:

```go
// database/user_db.go
type UserDBInterface interface {
    Create(user *entity.User) (*entity.User, error)
    FindByEmail(email string) (*entity.User, error)
    FindByID(id int64) (*entity.User, error)
    Update(user *entity.User) (*entity.User, error)
    Delete(id int64) error
}

// data/mysql/user.go
type UserDB struct {
    db *gorm.DB
}

func (r *UserDB) Create(user *entity.User) (*entity.User, error) {
    // Implementação com GORM
}
```

## Diferenças Entre MySQL e MongoDB

### MySQL (Users e Stories)
- IDs: `int64` auto-incrementados
- GORM como ORM
- Relacionamentos: Foreign keys
- Localização: `data/mysql/`

### MongoDB (Bookmarks)
- IDs: ObjectID (string de 24 chars hex)
- Driver nativo MongoDB
- Documentos flexíveis
- Localização: `data/mongo/`

**Por que MongoDB para Bookmarks?**
- Estrutura flexível (nem todo bookmark tem todos os campos)
- Alta frequência de updates
- Não requer relacionamentos complexos

## Sistema de Autenticação

### Fluxo JWT

```
1. Cliente faz POST /login com email/password
2. Service valida credenciais
3. Token JWT gerado com claims: { user_id: X }
4. Token retornado ao cliente
5. Cliente inclui token em requests: Authorization: Bearer {token}
6. Middleware valida token e injeta user_id no contexto
7. Controller acessa user_id via context
```

### Middleware de Autenticação

Localizado em `server/middleware/auth.go`:

```go
// Extrai e valida JWT
// Injeta user_id no echo.Context
// Bloqueia requisições não autenticadas
```

### Acesso ao Usuário Autenticado

```go
// No controller
func (c *Controller) SomeHandler(ctx echo.Context) error {
    userID := ctx.Get("user_id").(int64)
    // Use userID...
}
```

## ViewModels (DTOs)

### Request ViewModels

Localização: `controller/viewmodel/request/`

**Propósito**: Definir estrutura de dados esperada do cliente

```go
type CreateUserRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
    if r.Name == "" || r.Email == "" || r.Password == "" {
        return errors.New("all fields are required")
    }
    return nil
}
```

### Response ViewModels

Localização: `controller/viewmodel/response/`

**Propósito**: Formatar dados de retorno ao cliente

```go
type UserResponse struct {
    ID    int64  `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    // Note: Password NÃO incluída
}
```

## Constantes

Localização: `constants/constants.go`

**Usar constantes** para strings repetidas:

```go
const (
    ERROR_NAME_REQUIRED = "name is required"
    ERROR_EMAIL_INVALID = "invalid email format"
    ERROR_USER_NOT_FOUND = "user not found"
)
```

## Integração MyAnimeList

Localização: `integration/my_anime_list.go`

**Funcionalidades**:
- Buscar dados de animes/mangás por ID
- Buscar por nome
- OAuth2 com refresh token

**Uso no Service**:

```go
animeData, err := s.integration.GetAnimeByID(malID)
if err != nil {
    // API externa falhou, lidar apropriadamente
}
```

## Testes (Guidelines)

Embora não implementados ainda, a estrutura suporta:

### Unit Tests
- Testar entities isoladamente
- Mockar repositories para testar services
- Testar validações

### Integration Tests
- Testar handlers com servidor HTTP de teste
- Usar bancos de dados de teste
- Testar fluxos completos

## Comandos de Desenvolvimento

```bash
# Executar aplicação
go run main.go

# Build otimizado
go build -o amsvault -ldflags="-s -w"

# Testes (quando implementados)
go test ./...

# Verificar código
go vet ./...
go fmt ./...

# Atualizar dependências
go mod tidy
go mod download

# Ver dependências
go mod graph
```

## Docker e Database

### Iniciar bancos de dados

```bash
cd build
docker-compose up -d
```

### Executar migrations

SQL inicial em `build/database/initial.sql`

## Debugging

### Logs

Adicionar logs nos pontos chave:

```go
fmt.Printf("Debug: user_id=%d, story_id=%d\n", userID, storyID)
```

### Erros Comuns

1. **"invalid memory address"**: Ponteiro nil não checado
2. **"cannot bind"**: JSON de request incompatível com struct
3. **"unauthorized"**: Token ausente/inválido/expirado
4. **"user not found"**: Busca no DB falhou

## Modificações Comuns

### Adicionar Campo a Entity

```go
// 1. Adicionar campo à struct
type User struct {
    // ...campos existentes...
    Phone string `json:"phone"`
}

// 2. Atualizar construtor NewUser()
func NewUser(name, email, password, phone string) (*User, error) {
    // ...
    return &User{
        // ...
        Phone: phone,
    }, nil
}

// 3. Atualizar Validate() se necessário
func (u *User) Validate() error {
    // Adicionar validação de phone
}

// 4. Atualizar ViewModels
// 5. Atualizar migrations de DB
// 6. Atualizar documentação API
```

### Adicionar Novo Endpoint

```go
// 1. Criar ViewModels (request/response)
// 2. Adicionar método no Service
// 3. Adicionar handler no Controller
// 4. Registrar rota em server/router/
// 5. Adicionar middleware se necessário
// 6. Documentar em API_DOCUMENTATION.md
```

### Modificar Validação

```go
// Entity validation (regras de domínio)
func (s *Story) Validate() error {
    if s.TotalEpisode < 0 {
        return errors.New(constants.ERROR_EPISODE_INVALID)
    }
    // Nova validação aqui
}

// Service validation (regras de negócio)
func (s *Service) CreateStory(req StoryRequest) error {
    // Verificar duplicação
    // Validar dependências
    // etc.
}
```

## Perguntas Frequentes

### Q: Por que dois bancos de dados?

**A**: MySQL para dados estruturados e relacionais (users, stories). MongoDB para dados flexíveis e de alta frequência de update (bookmarks).

### Q: Por que GORM só no MySQL?

**A**: GORM é específico para SQL. MongoDB usa driver nativo com interface diferente.

### Q: Como adicionar novo campo opcional?

**A**: Use ponteiro na struct e omit no JSON: `Phone *string `json:"phone,omitempty"``

### Q: Como fazer soft delete?

**A**: Adicione campo `DeletedAt time.Time` e filtre por ele nas queries. GORM suporta isso nativamente.

### Q: Como paginar resultados?

**A**: Adicione parâmetros `limit` e `offset` nos repositories e passe via query params.

## Checklist de Modificação

Ao modificar código, verificar:

- [ ] Validação adicionada nas camadas apropriadas
- [ ] Erros customizados usados
- [ ] Construtores `NewX()` usados
- [ ] Interfaces de repositório atualizadas
- [ ] Testes adicionados (quando implementados)
- [ ] Documentação atualizada (API_DOCUMENTATION.md)
- [ ] Constantes usadas para strings repetidas
- [ ] Logs adicionados para debugging
- [ ] Status HTTP apropriados retornados

## Filosofia de Código

1. **Explícito sobre Implícito**: Preferir código claro a "mágico"
2. **Fail Fast**: Validar cedo, falhar cedo
3. **Single Responsibility**: Uma função, uma responsabilidade
4. **Composition over Inheritance**: Go não tem herança, usar composição
5. **Errors are Values**: Tratar erros explicitamente, não panic

## Próximos Passos (Roadmap)

- Implementar testes unitários e de integração
- Adicionar logs estruturados (ex: zerolog)
- Implementar paginação em listagens
- Adicionar cache (Redis) para dados frequentes
- Melhorar documentação com Swagger/OpenAPI
- Implementar rate limiting
- Adicionar métricas e observabilidade

---

**Para Claude**: Ao trabalhar neste projeto, sempre considere a arquitetura em camadas, use os padrões estabelecidos, valide em múltiplas camadas e mantenha a separação de responsabilidades. Quando em dúvida sobre estrutura, consulte código existente similar.

**Última Atualização**: Fevereiro 2026
