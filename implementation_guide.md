# Guia de Implementação de Melhorias - AMSVault

Este documento fornece orientações detalhadas para IAs implementarem melhorias na aplicação AMSVault. Cada seção indica claramente o **status atual** e os **passos necessários**.

> **Importante**: Sempre consulte `agents.md` e `claude.md` para entender a arquitetura antes de implementar.

---

## 📋 Índice

1. [Como Usar Este Guia](#como-usar-este-guia)
2. [Legenda de Status](#legenda-de-status)
3. [Arquitetura e Design](#arquitetura-e-design)
   - Middleware de Autenticação/Autorização
   - Substituir Panics
   - Versionamento da API
   - Graceful Shutdown
4. [Banco de Dados](#banco-de-dados)
   - Sistema de Migrações
   - Índices de Performance
   - Transações
   - Soft Delete Consistente
5. [Segurança](#segurança)
   - Rate Limiting
   - Validação de Entrada
   - Headers de Segurança
   - CORS
6. [Funcionalidades](#funcionalidades)
   - Paginação
   - Busca Avançada
   - Cache
7. [Testes](#testes)
   - Testes Unitários
   - Testes de Integração
8. [Documentação](#documentação)
   - Swagger/OpenAPI
   - Diagrama de BD
   - README
9. [DevOps e Infraestrutura](#devops-e-infraestrutura)
   - CI/CD
   - Docker
   - Logs Estruturados
10. [Qualidade de Código](#qualidade-de-código)
    - Linting e Formatação
11. [Checklist de Implementação](#checklist-de-implementação)
12. [Priorização Sugerida](#priorização-sugerida)

---

## Como Usar Este Guia

Este guia é otimizado para IAs que precisam implementar melhorias no projeto. Para cada item:

1. **Verifique o status** - Não reimplemente o que já existe
2. **Leia o contexto** - Entenda por que a melhoria é necessária
3. **Siga os passos** - Implementação passo a passo com código real
4. **Teste** - Sempre teste após implementar
5. **Documente** - Atualize documentação quando relevante

---

## Legenda de Status

- ✅ **JÁ IMPLEMENTADO**: Funcionalidade já existe, não precisa ser implementada
- ⚠️ **IMPLEMENTAÇÃO PARCIAL**: Existe parcialmente, necessita melhorias
- 🔴 **PRECISA IMPLEMENTAR**: Não implementado, necessita desenvolvimento completo

---

## Arquitetura e Design

### ✅ 1. Tratamento de Erros Centralizado (JÁ IMPLEMENTADO)

**Status**: Sistema de erros customizados já existe em `pkg/errors/`

**O que já existe**:
- Tipos de erro: `VALIDATION_ERROR`, `NOT_FOUND`, `ALREADY_EXISTS`, `INTERNAL_ERROR`, `DATABASE_ERROR`, `EXTERNAL_SERVICE_ERROR`
- Funções construtoras: `NewValidationError()`, `NewNotFoundError()`, etc.
- Usado em todo o projeto

**Não fazer**: Recriar sistema de erros  
**Pode fazer**: Adicionar novos tipos de erro se necessário seguindo padrão existente

---

### ⚠️ 2. Middleware de Autenticação (IMPLEMENTAÇÃO PARCIAL)

**Status**: Autenticação JWT existe em `server/middleware/auth.go`, mas falta autorização baseada em roles

**O que já existe**:
- Middleware JWT que valida tokens
- Extração de `user_id` do token
- Proteção de rotas privadas

**O que falta implementar**:

**Passo 1**: Adicionar campo `role` à entidade User
```go
// entity/user.go
type User struct {
    ID        int64
    Name      string
    Email     string
    Password  string
    Role      string    // Adicionar: "admin", "user", "moderator"
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

**Passo 2**: Criar middleware de autorização baseada em roles
```go
// server/middleware/authorization.go
func RequireRole(allowedRoles ...string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            userRole := c.Get("user_role").(string)
            // Verificar se userRole está em allowedRoles
            // Retornar 403 se não autorizado
        }
    }
}
```

**Passo 3**: Aplicar em rotas específicas
```go
// server/router/
adminRoutes.DELETE("/user/:id", ctrl.DeleteUser, middleware.RequireRole("admin"))
```

---

### ⚠️ 3. Substituir Panic por Tratamento de Erros

**Status**: Alguns `panic` ainda existem no código

**Locais principais onde há panic**:
- `main.go`: Falha ao carregar config ou conectar banco
- `config/config.go`: Falha ao ler .env

**Como implementar**:

**Passo 1**: Identificar todos os panics
```bash
grep -r "panic(" --include="*.go" .
```

**Passo 2**: Substituir em `main.go`
```go
// ❌ ANTES
if err := viper.ReadInConfig(); err != nil {
    panic(err)
}

// ✅ DEPOIS
if err := viper.ReadInConfig(); err != nil {
    log.Fatalf("Failed to load config: %v", err)
    return // ou os.Exit(1)
}
```

**Passo 3**: Substituir em inicializações
- Retornar erro em vez de panic
- Tratar erro no chamador (main.go)
- Adicionar logs descritivos

---

### ✅ 4. Interfaces para Dependências (JÁ IMPLEMENTADO)

**Status**: Interfaces já existem para repositórios

**O que já existe**:
- `database/user_db.go`: Interface `UserDBInterface`
- `database/story_db.go`: Interface `StoryDBInterface`
- Implementações em `data/mysql/` e `data/mongo/`
- Injeção de dependência via construtores

**Não fazer**: Recriar interfaces  
**Pode fazer**: Adicionar novas interfaces seguindo o padrão existente

---

### ✅ 5. Clean Architecture (JÁ IMPLEMENTADO)

**Status**: Arquitetura em camadas já está implementada

**Estrutura atual**:
```
HTTP (server) → Controller → Service → Repository (database/) → Entity
```

**Não fazer**: Reorganizar estrutura de pastas  
**Observação**: A arquitetura atual segue Clean Architecture, apenas com nomenclatura diferente

---

### 🔴 6. Versionamento da API

**Status**: API não possui versionamento

**Como implementar**:

**Passo 1**: Criar estrutura de versões no router
```go
// server/router/router.go
func NewRouter(e *echo.Echo, ctrl *controller.Controller) {
    // Grupo v1
    v1 := e.Group("/api/v1")
    
    // Rotas públicas
    v1.POST("/login", ctrl.Login)
    v1.POST("/user", ctrl.CreateUser)
    
    // Rotas privadas
    auth := v1.Group("", middleware.JWTAuth(ctrl.Config.TokenAuth))
    auth.GET("/user", ctrl.GetUser)
    // ... outras rotas
}
```

**Passo 2**: Adicionar header de versão nas respostas
```go
// Middleware para adicionar versão
func AddVersionHeader(version string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            c.Response().Header().Set("API-Version", version)
            return next(c)
        }
    }
}
```

**Passo 3**: Atualizar documentação da API

---

### ✅ 7. Separação de Entidades e DTOs (JÁ IMPLEMENTADO)

**Status**: DTOs já existem como ViewModels

**O que já existe**:
- `controller/viewmodel/request/`: DTOs de requisição
- `controller/viewmodel/response/`: DTOs de resposta
- Transformação entre DTOs e Entities nos controllers

**Não fazer**: Recriar ViewModels

---

### 🔴 8. Graceful Shutdown

**Status**: Não implementado

**Como implementar**:

**Passo 1**: Capturar sinais do SO em `main.go`
```go
// main.go
func main() {
    // ... inicializações existentes ...
    
    // Canal para sinais do SO
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    
    // Inicia servidor em goroutine
    go func() {
        if err := srv.Start(cfg, ctrl, nil); err != nil && err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()
    
    // Aguarda sinal
    <-quit
    log.Println("Shutting down server...")
    
    // Context com timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Shutdown graceful
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal(err)
    }
    
    // Fecha conexões de BD
    data.Close()
    log.Println("Server exited")
}
```

**Passo 2**: Implementar Shutdown no servidor
```go
// server/server.go
func (s *Server) Shutdown(ctx context.Context) error {
    return s.echo.Shutdown(ctx)
}
```

---

## Banco de Dados

### 🔴 1. Implementar Sistema de Migrações

**Status**: Apenas existe `build/database/initial.sql`, sem sistema de migrações

**Por que é importante**: Facilita controle de versão do schema e deploys

**Como implementar**:

**Passo 1**: Instalar ferramenta de migração
```bash
go get -u github.com/golang-migrate/migrate/v4
go get -u github.com/golang-migrate/migrate/v4/database/mysql
go get -u github.com/golang-migrate/migrate/v4/source/file
```

**Passo 2**: Criar estrutura de migrações
```bash
mkdir -p migrations
```

**Passo 3**: Criar migração inicial baseada no schema atual
```bash
# migrations/000001_initial_schema.up.sql
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

# migrations/000001_initial_schema.down.sql
DROP TABLE IF EXISTS users;
```

**Passo 4**: Executar migrações na inicialização
```go
// data/migrations.go
func RunMigrations(cfg *config.Config) error {
    dbURL := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s",
        cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
    
    m, err := migrate.New("file://migrations", dbURL)
    if err != nil {
        return err
    }
    
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return err
    }
    
    return nil
}
```

---

### 🔴 2. Adicionar Índices de Performance

**Status**: Não existem índices além das PKs/FKs

**Campos que precisam de índices**:

**Stories**:
- `name`: Usado em buscas LIKE
- `mal_id`: Usado para lookups diretos
- `status`: Usado em filtragens

**Users**:
- `email`: Já é UNIQUE, mas adicionar índice explícito ajuda

**Como implementar**:

**Passo 1**: Criar migração
```sql
-- migrations/000002_add_indexes.up.sql
CREATE INDEX idx_stories_name ON stories(name);
CREATE INDEX idx_stories_mal_id ON stories(mal_id);
CREATE INDEX idx_stories_status ON stories(status);
CREATE INDEX idx_stories_source ON stories(source);
CREATE INDEX idx_users_email ON users(email);

-- migrations/000002_add_indexes.down.sql
DROP INDEX idx_stories_name ON stories;
DROP INDEX idx_stories_mal_id ON stories;
DROP INDEX idx_stories_status ON stories;
DROP INDEX idx_stories_source ON stories;
DROP INDEX idx_users_email ON users;
```

**Passo 2**: Executar migração
```bash
migrate -path migrations -database "mysql://..." up
```

**Passo 3**: Medir performance com EXPLAIN
```sql
EXPLAIN SELECT * FROM stories WHERE name LIKE '%anime%';
```

---

### ⚠️ 3. Implementar Transações

**Status**: Operações individuais não usam transações

**Cenários que precisam de transações**:
1. Deletar usuário + seus bookmarks
2. Criar story + validar duplicação

**Como implementar**:

**Passo 1**: Adicionar método de transação no repositório
```go
// database/transaction.go
type TransactionFunc func(*gorm.DB) error

type DatabaseInterface interface {
    WithTransaction(fn TransactionFunc) error
}

// data/data.go
func (d *Data) WithTransaction(fn TransactionFunc) error {
    tx := d.mysqlDB.Begin()
    if tx.Error != nil {
        return tx.Error
    }
    
    if err := fn(tx); err != nil {
        tx.Rollback()
        return err
    }
    
    return tx.Commit().Error
}
```

**Passo 2**: Usar no service
```go
// service/user.go
func (s *Service) DeleteUser(userID int64) error {
    return s.data.WithTransaction(func(tx *gorm.DB) error {
        // 1. Deletar bookmarks do usuário
        if err := s.bookmarksDB.DeleteByUserID(tx, userID); err != nil {
            return err
        }
        
        // 2. Deletar usuário
        if err := s.userDB.Delete(tx, userID); err != nil {
            return err
        }
        
        return nil
    })
}
```

---

### ⚠️ 4. Soft Delete Consistente

**Status**: Bookmarks usam soft delete (MongoDB), mas Users e Stories não

**Como implementar**:

**Passo 1**: Adicionar campo deleted_at
```sql
-- migrations/000003_add_soft_delete.up.sql
ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL;
ALTER TABLE stories ADD COLUMN deleted_at TIMESTAMP NULL DEFAULT NULL;

-- migrations/000003_add_soft_delete.down.sql
ALTER TABLE users DROP COLUMN deleted_at;
ALTER TABLE stories DROP COLUMN deleted_at;
```

**Passo 2**: Usar GORM soft delete
```go
// entity/user.go
type User struct {
    ID        int64
    Name      string
    Email     string
    Password  string
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"` // Adicionar
}
```

**Passo 3**: GORM automaticamente filtra registros deletados nas queries

---

## Segurança

### 🔴 1. Rate Limiting

**Status**: Não implementado

**Por que é importante**: Prevenir abuso da API e ataques DDoS

**Como implementar**:

**Passo 1**: Instalar biblioteca
```bash
go get github.com/ulule/limiter/v3
go get github.com/ulule/limiter/v3/drivers/store/memory
```

**Passo 2**: Criar middleware
```go
// server/middleware/rate_limit.go
package middleware

import (
    "github.com/labstack/echo/v4"
    "github.com/ulule/limiter/v3"
    "github.com/ulule/limiter/v3/drivers/middleware/stdlib"
    "github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimit() echo.MiddlewareFunc {
    rate := limiter.Rate{
        Period: 1 * time.Minute,
        Limit:  60, // 60 requests por minuto
    }
    
    store := memory.NewStore()
    instance := limiter.New(store, rate)
    
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            context, err := instance.Get(c.Request().Context(), c.RealIP())
            if err != nil {
                return echo.NewHTTPError(500, "rate limiter error")
            }
            
            c.Response().Header().Set("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
            c.Response().Header().Set("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
            c.Response().Header().Set("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))
            
            if context.Reached {
                return echo.NewHTTPError(429, "rate limit exceeded")
            }
            
            return next(c)
        }
    }
}
```

**Passo 3**: Aplicar no router
```go
// server/router/router.go
e.Use(middleware.RateLimit())
```

---

### 🔴 2. Validação de Entrada Robusta

**Status**: Validação básica existe, mas pode melhorar

**O que falta**:
- Sanitização de strings
- Validação de formato de email
- Validação de comprimento de campos
- Proteção contra SQL injection (GORM já protege, mas validar entrada)

**Como implementar**:

**Passo 1**: Instalar biblioteca de validação
```bash
go get github.com/go-playground/validator/v10
```

**Passo 2**: Adicionar tags de validação nos ViewModels
```go
// controller/viewmodel/user.go
type CreateUserRequest struct {
    Name     string `json:"name" validate:"required,min=3,max=100"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8,max=100"`
}
```

**Passo 3**: Criar função de validação
```go
// utils/validator.go
var validate = validator.New()

func ValidateStruct(s interface{}) error {
    return validate.Struct(s)
}
```

**Passo 4**: Usar nos controllers
```go
func (c *Controller) CreateUser(ctx echo.Context) error {
    var req viewmodel.CreateUserRequest
    if err := ctx.Bind(&req); err != nil {
        return ctx.JSON(400, ErrorResponse{Message: "invalid request"})
    }
    
    if err := utils.ValidateStruct(&req); err != nil {
        return ctx.JSON(400, ErrorResponse{Message: err.Error()})
    }
    
    // ... continuar processamento
}
```

---

### 🔴 3. Headers de Segurança

**Status**: Não implementados

**Como implementar**:

**Passo 1**: Criar middleware de segurança
```go
// server/middleware/security_headers.go
func SecurityHeaders() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // Prevenir XSS
            c.Response().Header().Set("X-XSS-Protection", "1; mode=block")
            
            // Prevenir clickjacking
            c.Response().Header().Set("X-Frame-Options", "DENY")
            
            // Prevenir MIME type sniffing
            c.Response().Header().Set("X-Content-Type-Options", "nosniff")
            
            // Content Security Policy
            c.Response().Header().Set("Content-Security-Policy", "default-src 'self'")
            
            // HSTS (apenas se usar HTTPS)
            // c.Response().Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
            
            // Referrer Policy
            c.Response().Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
            
            return next(c)
        }
    }
}
```

**Passo 2**: Aplicar globalmente
```go
// server/server.go
e.Use(middleware.SecurityHeaders())
```

---

### 🔴 4. CORS Configurável

**Status**: Pode não estar configurado ou estar muito permissivo

**Como implementar**:

**Passo 1**: Configurar CORS apropriadamente
```go
// server/server.go
e.Use(middlewareEcho.CORSWithConfig(middlewareEcho.CORSConfig{
    AllowOrigins:     []string{"https://seu-frontend.com"}, // Específico!
    AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
    AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
    AllowCredentials: true,
    MaxAge:           86400,
}))
```

**Passo 2**: Adicionar origins permitidas no .env
```env
ALLOWED_ORIGINS=https://frontend.com,https://app.frontend.com
```

---

## Funcionalidades

### 🔴 1. Paginação

**Status**: Endpoints de listagem não possuem paginação

**Endpoints que precisam**:
- `GET /story?name=X` - Buscar histórias
- `GET /bookmarks/user/:user_id` - Listar bookmarks

**Como implementar**:

**Passo 1**: Criar struct de paginação
```go
// utils/pagination.go
type PaginationParams struct {
    Page     int `query:"page"`
    PageSize int `query:"page_size"`
}

type PaginatedResponse struct {
    Data       interface{} `json:"data"`
    Page       int         `json:"page"`
    PageSize   int         `json:"page_size"`
    TotalItems int64       `json:"total_items"`
    TotalPages int         `json:"total_pages"`
}

func NewPaginationParams(page, pageSize int) PaginationParams {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 20 // Default
    }
    return PaginationParams{Page: page, PageSize: pageSize}
}

func (p PaginationParams) GetOffset() int {
    return (p.Page - 1) * p.PageSize
}
```

**Passo 2**: Adicionar paginação no repositório
```go
// database/story_db.go
type StoryDBInterface interface {
    // ... métodos existentes
    FindByNamePaginated(name string, offset, limit int) ([]*entity.Story, int64, error)
}

// data/mysql/story.go
func (r *StoryDB) FindByNamePaginated(name string, offset, limit int) ([]*entity.Story, int64, error) {
    var stories []*entity.Story
    var total int64
    
    query := r.db.Model(&model.Story{}).Where("name LIKE ?", "%"+name+"%")
    
    // Contar total
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }
    
    // Buscar com paginação
    if err := query.Offset(offset).Limit(limit).Find(&stories).Error; err != nil {
        return nil, 0, err
    }
    
    return stories, total, nil
}
```

**Passo 3**: Usar no controller
```go
// controller/story.go
func (c *Controller) FindStoryByName(ctx echo.Context) error {
    name := ctx.QueryParam("name")
    pagination := utils.NewPaginationParams(
        ctx.QueryParam("page"),
        ctx.QueryParam("page_size"),
    )
    
    stories, total, err := c.service.FindStoryByNamePaginated(name, pagination)
    if err != nil {
        return ctx.JSON(500, err)
    }
    
    response := utils.PaginatedResponse{
        Data:       stories,
        Page:       pagination.Page,
        PageSize:   pagination.PageSize,
        TotalItems: total,
        TotalPages: int(math.Ceil(float64(total) / float64(pagination.PageSize))),
    }
    
    return ctx.JSON(200, response)
}
```

---

### 🔴 2. Busca Avançada de Stories

**Status**: Apenas busca por nome existe

**Filtros úteis**:
- Status (ongoing, completed)
- Source (anime, manga, novel)
- Ordenação (por nome, data, popularidade)

**Como implementar**:

**Passo 1**: Criar ViewModel de filtros
```go
// controller/viewmodel/request/story_filter.go
type StoryFilterRequest struct {
    Name     string `query:"name"`
    Status   string `query:"status"`
    Source   string `query:"source"`
    OrderBy  string `query:"order_by"`  // name, created_at, mal_id
    Order    string `query:"order"`     // asc, desc
    Page     int    `query:"page"`
    PageSize int    `query:"page_size"`
}
```

**Passo 2**: Implementar busca com filtros no repositório
```go
// data/mysql/story.go
func (r *StoryDB) SearchWithFilters(filter StoryFilterRequest) ([]*entity.Story, int64, error) {
    query := r.db.Model(&model.Story{})
    
    // Aplicar filtros
    if filter.Name != "" {
        query = query.Where("name LIKE ?", "%"+filter.Name+"%")
    }
    if filter.Status != "" {
        query = query.Where("status = ?", filter.Status)
    }
    if filter.Source != "" {
        query = query.Where("source = ?", filter.Source)
    }
    
    // Ordenação
    orderBy := "name" // default
    if filter.OrderBy != "" {
        orderBy = filter.OrderBy
    }
    order := "ASC"
    if strings.ToUpper(filter.Order) == "DESC" {
        order = "DESC"
    }
    query = query.Order(fmt.Sprintf("%s %s", orderBy, order))
    
    // Contar e paginar
    var total int64
    query.Count(&total)
    
    var stories []*entity.Story
    offset := (filter.Page - 1) * filter.PageSize
    query.Offset(offset).Limit(filter.PageSize).Find(&stories)
    
    return stories, total, nil
}
```

---

### 🔴 3. Cache de Dados

**Status**: Sem cache implementado

**Dados que se beneficiam de cache**:
- Stories (raramente mudam)
- Informações de usuários (sessão)

**Como implementar**:

**Passo 1**: Instalar biblioteca de cache
```bash
go get github.com/patrickmn/go-cache
```

**Passo 2**: Adicionar cache na camada de data
```go
// data/cache.go
package data

import (
    "time"
    gocache "github.com/patrickmn/go-cache"
)

type Cache struct {
    cache *gocache.Cache
}

func NewCache() *Cache {
    return &Cache{
        cache: gocache.New(5*time.Minute, 10*time.Minute),
    }
}

func (c *Cache) Get(key string) (interface{}, bool) {
    return c.cache.Get(key)
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
    c.cache.Set(key, value, duration)
}

func (c *Cache) Delete(key string) {
    c.cache.Delete(key)
}
```

**Passo 3**: Usar cache no service
```go
// service/story.go
func (s *Service) FindStoryByID(id int64) (*entity.Story, error) {
    cacheKey := fmt.Sprintf("story:%d", id)
    
    // Tentar buscar do cache
    if cached, found := s.cache.Get(cacheKey); found {
        return cached.(*entity.Story), nil
    }
    
    // Buscar do banco
    story, err := s.storyDB.FindByID(id)
    if err != nil {
        return nil, err
    }
    
    // Armazenar no cache
    s.cache.Set(cacheKey, story, 10*time.Minute)
    
    return story, nil
}
```

**Passo 4**: Invalidar cache ao atualizar
```go
func (s *Service) UpdateStory(story *entity.Story) error {
    if err := s.storyDB.Update(story); err != nil {
        return err
    }
    
    // Invalidar cache
    cacheKey := fmt.Sprintf("story:%d", story.ID)
    s.cache.Delete(cacheKey)
    
    return nil
}
```

---

## Testes

### 🔴 1. Testes Unitários

**Status**: Não implementados

**Prioridade**: Começar por entities e services

**Como implementar**:

**Passo 1**: Criar estrutura de testes
```bash
mkdir -p entity/test
mkdir -p service/test
```

**Passo 2**: Exemplo de teste de entity
```go
// entity/user_test.go
package entity_test

import (
    "testing"
    "github.com.br/GregoryLacerda/AMSVault/entity"
)

func TestNewUser(t *testing.T) {
    t.Run("should create user with hashed password", func(t *testing.T) {
        user, err := entity.NewUser("John Doe", "john@example.com", "password123")
        
        if err != nil {
            t.Errorf("Expected no error, got %v", err)
        }
        
        if user.Name != "John Doe" {
            t.Errorf("Expected name 'John Doe', got %s", user.Name)
        }
        
        if user.Password == "password123" {
            t.Error("Password should be hashed, but is plain text")
        }
    })
    
    t.Run("should validate password correctly", func(t *testing.T) {
        user, _ := entity.NewUser("John", "john@example.com", "password123")
        
        if !user.ValidatePassword("password123") {
            t.Error("Password validation failed for correct password")
        }
        
        if user.ValidatePassword("wrongpassword") {
            t.Error("Password validation passed for incorrect password")
        }
    })
}
```

**Passo 3**: Exemplo de teste de service com mock
```go
// service/user_test.go
package service_test

import (
    "testing"
    "github.com.br/GregoryLacerda/AMSVault/entity"
    "github.com.br/GregoryLacerda/AMSVault/service"
)

// Mock do repositório
type MockUserDB struct {
    users map[string]*entity.User
}

func (m *MockUserDB) FindByEmail(email string) (*entity.User, error) {
    if user, exists := m.users[email]; exists {
        return user, nil
    }
    return nil, errors.NewNotFoundError("user not found")
}

func (m *MockUserDB) Create(user *entity.User) (*entity.User, error) {
    m.users[user.Email] = user
    return user, nil
}

func TestCreateUser(t *testing.T) {
    mockDB := &MockUserDB{users: make(map[string]*entity.User)}
    svc := service.NewService(nil, mockDB, nil, nil)
    
    t.Run("should create new user", func(t *testing.T) {
        user, err := svc.CreateUser("John", "john@example.com", "password123")
        
        if err != nil {
            t.Errorf("Expected no error, got %v", err)
        }
        
        if user.Email != "john@example.com" {
            t.Errorf("Expected email 'john@example.com', got %s", user.Email)
        }
    })
    
    t.Run("should fail for duplicate email", func(t *testing.T) {
        _, err := svc.CreateUser("Jane", "john@example.com", "password456")
        
        if err == nil {
            t.Error("Expected error for duplicate email, got nil")
        }
    })
}
```

**Passo 4**: Executar testes
```bash
go test ./entity/...
go test ./service/...
go test ./... -v
```

---

### 🔴 2. Testes de Integração

**Status**: Não implementados

**Como implementar**:

**Passo 1**: Criar testes de integração para controllers
```go
// controller/user_integration_test.go
// +build integration

package controller_test

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "github.com/labstack/echo/v4"
)

func TestCreateUserIntegration(t *testing.T) {
    // Setup
    e := echo.New()
    ctrl := setupTestController() // Função helper
    
    // Criar request
    userJSON := `{"name":"John Doe","email":"john@example.com","password":"password123"}`
    req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(userJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    
    // Executar
    if err := ctrl.CreateUser(c); err != nil {
        t.Errorf("Handler returned error: %v", err)
    }
    
    // Verificar
    if rec.Code != http.StatusCreated {
        t.Errorf("Expected status 201, got %d", rec.Code)
    }
}
```

**Passo 2**: Usar banco de dados de teste
```go
func setupTestDB() *gorm.DB {
    db, err := gorm.Open(mysql.Open("test_dsn"), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    
    // Migrar schema
    db.AutoMigrate(&model.User{}, &model.Story{})
    
    return db
}

func teardownTestDB(db *gorm.DB) {
    // Limpar dados
    db.Exec("TRUNCATE users")
    db.Exec("TRUNCATE stories")
}
```

---

## Documentação

### ⚠️ 1. Swagger/OpenAPI (MELHORAR EXISTENTE)

**Status**: Existe `API_DOCUMENTATION.md`, mas não é interativo

**Como implementar Swagger**:

**Passo 1**: Instalar Swaggo
```bash
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/echo-swagger
```

**Passo 2**: Adicionar anotações nos controllers
```go
// controller/user.go

// CreateUser godoc
// @Summary      Criar novo usuário
// @Description  Registra um novo usuário no sistema
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      viewmodel.CreateUserRequest  true  "Dados do usuário"
// @Success      201   {object}  viewmodel.MessageResponse
// @Failure      400   {object}  errors.CustomError
// @Failure      409   {object}  errors.CustomError
// @Router       /user [post]
func (c *Controller) CreateUser(ctx echo.Context) error {
    // ... implementação
}
```

**Passo 3**: Gerar documentação
```bash
swag init -g main.go
```

**Passo 4**: Adicionar rota Swagger
```go
// server/router/router.go
import echoSwagger "github.com/swaggo/echo-swagger"

func NewRouter(e *echo.Echo, ctrl *controller.Controller) {
    // ... rotas existentes
    
    // Swagger
    e.GET("/swagger/*", echoSwagger.WrapHandler)
}
```

**Passo 5**: Acessar documentação
```
http://localhost:8080/swagger/index.html
```

---

### 🔴 2. Diagrama do Banco de Dados

**Status**: Não existe

**Como criar**:

**Opção 1**: Usar ferramenta online (dbdiagram.io, draw.io)

**Opção 2**: Gerar automaticamente com ferramenta
```bash
# Usar SchemaSpy ou MySQL Workbench
# Gerar diagrama ER visual
```

**Passo 3**: Adicionar ao repositório
```bash
# Salvar em docs/database_schema.png
# Referenciar no README.md
```

---

### ⚠️ 3. README Completo (MELHORAR)

**Status**: Pode existir mas provavelmente incompleto

**Seções necessárias**:

```markdown
# AMSVault

## 📖 Visão Geral
[Descrição do projeto]

## ✨ Funcionalidades
- Gerenciamento de usuários com autenticação JWT
- CRUD de histórias (animes, mangás, séries)
- Sistema de bookmarks para rastreamento de progresso
- Integração com MyAnimeList API

## 🛠️ Tecnologias
- Go 1.24+
- Echo Framework
- MySQL + MongoDB
- GORM
- JWT Authentication

## 📋 Pré-requisitos
- Go 1.24 ou superior
- MySQL 8.0+
- MongoDB 6.0+
- Docker (opcional)

## 🚀 Instalação

### Com Docker
```bash
cd build
docker-compose up -d
```

### Manual
```bash
# 1. Clonar repositório
git clone https://github.com/user/AMSVault

# 2. Instalar dependências
go mod download

# 3. Configurar .env
cp .env.example .env
# Editar .env com suas credenciais

# 4. Executar migrações
go run migrations/migrate.go

# 5. Iniciar aplicação
go run main.go
```

## ⚙️ Configuração
[Explicar variáveis de ambiente]

## 📚 Documentação da API
- Swagger: http://localhost:8080/swagger/
- Markdown: [API_DOCUMENTATION.md](API_DOCUMENTATION.md)

## 🧪 Testes
```bash
# Testes unitários
go test ./...

# Testes de integração
go test -tags=integration ./...

# Coverage
go test -cover ./...
```

## 📁 Estrutura do Projeto
[Explicar organização de pastas]

## 🤝 Contribuindo
[Guia de contribuição]

## 📄 Licença
[Licença do projeto]
```

---

## DevOps e Infraestrutura

### ⚠️ 1. CI/CD (IMPLEMENTAÇÃO PARCIAL)

**Status**: Pode existir docker-compose mas não CI/CD completo

**Como implementar**:

**Passo 1**: Criar GitHub Actions workflow
```yaml
# .github/workflows/ci.yml
name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      mysql:
        image: mysql:8.0
        env:
          MYSQL_ROOT_PASSWORD: root
          MYSQL_DATABASE: amsvault_test
        ports:
          - 3306:3306
        options: >-
          --health-cmd="mysqladmin ping"
          --health-interval=10s
          --health-timeout=5s
          --health-retries=3
      
      mongodb:
        image: mongo:6.0
        ports:
          - 27017:27017

    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.24'
    
    - name: Install dependencies
      run: go mod download
    
    - name: Run tests
      run: go test -v -cover ./...
      env:
        DB_HOST: localhost
        DB_PORT: 3306
        MONGO_URI: mongodb://localhost:27017
    
    - name: Run linting
      run: |
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
        golangci-lint run
    
    - name: Build
      run: go build -v ./...

  deploy:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
    - name: Deploy to production
      run: echo "Deploy steps here"
```

---

### ⚠️ 2. Docker (MELHORAR EXISTENTE)

**Status**: docker-compose.yaml existe, mas falta Dockerfile da aplicação

**Como implementar**:

**Passo 1**: Criar Dockerfile otimizado
```dockerfile
# Dockerfile
# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Instalar dependências de build
RUN apk add --no-cache git

# Copiar go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o amsvault .

# Stage 2: Runtime
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copiar binário do stage de build
COPY --from=builder /app/amsvault .
COPY --from=builder /app/.env .

# Expor porta
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Comando de execução
CMD ["./amsvault"]
```

**Passo 2**: Atualizar docker-compose.yaml
```yaml
# build/docker-compose.yaml
version: '3.8'

services:
  app:
    build:
      context: ..
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=mysql
      - MONGO_URI=mongodb://mongodb:27017
    depends_on:
      mysql:
        condition: service_healthy
      mongodb:
        condition: service_started
    networks:
      - amsvault-network

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: ${DB_PASSWORD}
      MYSQL_DATABASE: ${DB_NAME}
    volumes:
      - mysql-data:/var/lib/mysql
      - ./database/initial.sql:/docker-entrypoint-initdb.d/initial.sql
    ports:
      - "3306:3306"
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      interval: 10s
      timeout: 5s
      retries: 3
    networks:
      - amsvault-network

  mongodb:
    image: mongo:6.0
    volumes:
      - mongo-data:/data/db
    ports:
      - "27017:27017"
    networks:
      - amsvault-network

volumes:
  mysql-data:
  mongo-data:

networks:
  amsvault-network:
    driver: bridge
```

**Passo 3**: Criar .dockerignore
```
# .dockerignore
.git
.github
*.md
.env.example
build/
tmp/
*.log
```

---

### 🔴 3. Logs Estruturados

**Status**: Provavelmente usando fmt.Println ou log padrão

**Como implementar**:

**Passo 1**: Instalar biblioteca de logging
```bash
go get -u go.uber.org/zap
```

**Passo 2**: Criar pacote de logging
```go
// pkg/logger/logger.go
package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init(environment string) error {
    var config zap.Config
    
    if environment == "production" {
        config = zap.NewProductionConfig()
    } else {
        config = zap.NewDevelopmentConfig()
        config.EncodingConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    }
    
    var err error
    Log, err = config.Build()
    if err != nil {
        return err
    }
    
    return nil
}

func Info(msg string, fields ...zap.Field) {
    Log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
    Log.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
    Log.Warn(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
    Log.Debug(msg, fields...)
}
```

**Passo 3**: Usar no código
```go
// main.go
import "github.com.br/GregoryLacerda/AMSVault/pkg/logger"

func main() {
    if err := logger.Init(os.Getenv("ENVIRONMENT")); err != nil {
        panic(err)
    }
    defer logger.Log.Sync()
    
    logger.Info("Starting AMSVault",
        zap.String("version", "1.0.0"),
        zap.String("port", cfg.WebServerPort),
    )
    
    // ... resto do código
}

// Em controllers/services
logger.Error("Failed to create user",
    zap.Error(err),
    zap.String("email", email),
    zap.Int64("user_id", userID),
)
```

**Passo 4**: Middleware de logging
```go
// server/middleware/logger.go
func RequestLogger() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            start := time.Now()
            
            err := next(c)
            
            logger.Info("HTTP Request",
                zap.String("method", c.Request().Method),
                zap.String("path", c.Request().URL.Path),
                zap.Int("status", c.Response().Status),
                zap.Duration("latency", time.Since(start)),
                zap.String("ip", c.RealIP()),
            )
            
            return err
        }
    }
}
```

---

## Qualidade de Código

### ⚠️ 1. Linting e Formatação

**Status**: Provavelmente não configurado

**Como implementar**:

**Passo 1**: Instalar golangci-lint
```bash
# Linux/Mac
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Windows
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

**Passo 2**: Criar configuração
```yaml
# .golangci.yml
linters:
  enable:
    - gofmt
    - golint
    - govet
    - errcheck
    - staticcheck
    - unused
    - gosimple
    - structcheck
    - varcheck
    - ineffassign
    - deadcode
    - typecheck
    - gosec
    - gocyclo
    - dupl

linters-settings:
  gocyclo:
    min-complexity: 15
  golint:
    min-confidence: 0.8
  
issues:
  exclude-rules:
    - path: _test\.go
      linters:
        - errcheck
        - gosec

run:
  timeout: 5m
  skip-dirs:
    - vendor
    - tmp
```

**Passo 3**: Executar
```bash
# Formatar código
go fmt ./...

# Executar linting
golangci-lint run

# Corrigir automaticamente o que for possível
golangci-lint run --fix
```

**Passo 4**: Pre-commit hook
```bash
# .git/hooks/pre-commit
#!/bin/sh
echo "Running linter..."
golangci-lint run
if [ $? -ne 0 ]; then
    echo "Linting failed. Please fix errors before committing."
    exit 1
fi

echo "Running tests..."
go test ./...
if [ $? -ne 0 ]; then
    echo "Tests failed. Please fix before committing."
    exit 1
fi
```

---

## Checklist de Implementação

Use este checklist para acompanhar o progresso:

### Arquitetura (3/8)
- [x] Tratamento de erros centralizado
- [x] Interfaces para dependências
- [x] Separação Entidades/DTOs
- [ ] Autorização baseada em roles
- [ ] Remover panics
- [ ] Versionamento da API
- [ ] Graceful shutdown
- [x] Clean Architecture

### Banco de Dados (0/4)
- [ ] Sistema de migrações
- [ ] Índices de performance
- [ ] Transações
- [ ] Soft delete consistente

### Segurança (0/4)
- [ ] Rate limiting
- [ ] Validação robusta de entrada
- [ ] Headers de segurança
- [ ] CORS configurável

### Funcionalidades (0/3)
- [ ] Paginação
- [ ] Busca avançada
- [ ] Cache

### Testes (0/2)
- [ ] Testes unitários
- [ ] Testes de integração

### Documentação (1/3)
- [x] Documentação da API (Markdown)
- [ ] Swagger/OpenAPI
- [ ] Diagrama do BD

### DevOps (1/3)
- [x] Docker compose
- [ ] CI/CD
- [ ] Logs estruturados

### Qualidade (0/1)
- [ ] Linting configurado

---

## Priorização Sugerida

### 🔴 Alta Prioridade (Fazer Primeiro)
1. Remover panics e melhorar tratamento de erros
2. Implementar testes unitários (ao menos basics)
3. Adicionar validação robusta de entrada
4. Implementar rate limiting
5. Logs estruturados
6. Graceful shutdown

### 🟡 Média Prioridade
7. Sistema de migrações de BD
8. Paginação
9. Headers de segurança
10. Índices de performance
11. Versionamento da API
12. CI/CD básico

### 🟢 Baixa Prioridade (Pode Esperar)
13. Cache
14. Busca avançada
15. Swagger/OpenAPI
16. Soft delete consistente
17. Autorização baseada em roles

---

## Recursos Adicionais

- **Documentação Go**: https://go.dev/doc/
- **Echo Framework**: https://echo.labstack.com/
- **GORM**: https://gorm.io/docs/
- **Best Practices**: https://github.com/golang-standards/project-layout
- **Clean Architecture**: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

---

**Nota Final para IAs**: Ao implementar qualquer uma destas melhorias, sempre:
1. Consulte `agents.md` e `claude.md` para entender o contexto
2. Siga os padrões de código existentes
3. Adicione testes quando possível
4. Atualize a documentação
5. Faça commits atômicos e descritivos
