package middleware

import (
	"net/http"
	"strings"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/pkg/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func JWTMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing or malformed JWT"})
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid authorization header format",
				})
			}

			token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.NewUnauthorizedError("invalid token signing method")
				}
				return []byte(cfg.JWTSecret), nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid or expired JWT",
				})
			}

			if !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid or expired JWT",
				})
			}

			claims, ok := token.Claims.(*JWTClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid token claims",
				})
			}

			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)

			return next(c)
		}
	}
}

func GetUserIDFromContext(c echo.Context) (int64, error) {
	userID := c.Get("user_id")
	if userID == nil {
		return 0, errors.NewUnauthorizedError("user_id not found in context")
	}

	id, ok := userID.(int64)
	if !ok {
		return 0, errors.NewInternalError("user_id has invalid type", nil)
	}

	return id, nil
}

func GetEmailFromContext(c echo.Context) (string, error) {
	userEmail := c.Get("email")
	if userEmail == nil {
		return "", errors.NewUnauthorizedError("email not found in context")
	}

	email, ok := userEmail.(string)
	if !ok {
		return "", errors.NewInternalError("email has invalid type", nil)
	}

	return email, nil
}
