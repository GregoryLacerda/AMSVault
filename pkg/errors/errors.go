package errors

import (
	"fmt"
	"net/http"
)

// ErrorType representa o tipo de erro
type ErrorType string

const (
	// Tipos de erro de domínio
	ErrorTypeValidation    ErrorType = "VALIDATION_ERROR"
	ErrorTypeNotFound      ErrorType = "NOT_FOUND"
	ErrorTypeAlreadyExists ErrorType = "ALREADY_EXISTS"
	ErrorTypeUnauthorized  ErrorType = "UNAUTHORIZED"
	ErrorTypeForbidden     ErrorType = "FORBIDDEN"
	ErrorTypeInternal      ErrorType = "INTERNAL_ERROR"
	ErrorTypeDatabase      ErrorType = "DATABASE_ERROR"
	ErrorTypeExternal      ErrorType = "EXTERNAL_SERVICE_ERROR"
)

// AppError representa um erro da aplicação
type AppError struct {
	Type       ErrorType `json:"type"`
	Message    string    `json:"message"`
	StatusCode int       `json:"-"`
	Cause      error     `json:"-"`
}

// Error implementa a interface error
func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (caused by: %v)", e.Type, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// GetStatusCode retorna o código HTTP apropriado
func (e *AppError) GetStatusCode() int {
	return e.StatusCode
}

// Unwrap implementa a interface para unwrapping de erros
func (e *AppError) Unwrap() error {
	return e.Cause
}

// Construtores para tipos específicos de erro

// NewValidationError cria um erro de validação
func NewValidationError(message string) *AppError {
	return &AppError{
		Type:       ErrorTypeValidation,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func New(message string) *AppError {
	return &AppError{
		Message: message,
	}
}

// NewNotFoundError cria um erro de recurso não encontrado
func NewNotFoundError(resource string) *AppError {
	return &AppError{
		Type:       ErrorTypeNotFound,
		Message:    fmt.Sprintf("%s not found", resource),
		StatusCode: http.StatusNotFound,
	}
}

// NewAlreadyExistsError cria um erro de recurso já existente
func NewAlreadyExistsError(resource string) *AppError {
	return &AppError{
		Type:       ErrorTypeAlreadyExists,
		Message:    fmt.Sprintf("%s already exists", resource),
		StatusCode: http.StatusConflict,
	}
}

// NewUnauthorizedError cria um erro de não autorizado
func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Type:       ErrorTypeUnauthorized,
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

// NewForbiddenError cria um erro de acesso proibido
func NewForbiddenError(message string) *AppError {
	return &AppError{
		Type:       ErrorTypeForbidden,
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}

// NewInternalError cria um erro interno
func NewInternalError(message string, cause error) *AppError {
	return &AppError{
		Type:       ErrorTypeInternal,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Cause:      cause,
	}
}

// NewDatabaseError cria um erro de banco de dados
func NewDatabaseError(operation string, cause error) *AppError {
	return &AppError{
		Type:       ErrorTypeDatabase,
		Message:    fmt.Sprintf("database error: %s", operation),
		StatusCode: http.StatusInternalServerError,
		Cause:      cause,
	}
}

// NewExternalServiceError cria um erro de serviço externo
func NewExternalServiceError(service string, cause error) *AppError {
	return &AppError{
		Type:       ErrorTypeExternal,
		Message:    fmt.Sprintf("external service error: %s", service),
		StatusCode: http.StatusServiceUnavailable,
		Cause:      cause,
	}
}

// Helpers para verificação de tipos de erro

// IsValidationError verifica se é um erro de validação
func IsValidationError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Type == ErrorTypeValidation
	}
	return false
}

// IsNotFoundError verifica se é um erro de não encontrado
func IsNotFoundError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Type == ErrorTypeNotFound
	}
	return false
}

// IsDatabaseError verifica se é um erro de banco de dados
func IsDatabaseError(err error) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Type == ErrorTypeDatabase
	}
	return false
}
