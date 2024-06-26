package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type JwtClaims struct {
	UserUuid string `json:"user_uuid,omitempty"`
	jwt.StandardClaims
}

func ResponseRenderer(message string, payload ...gin.H) gin.H {
	response := gin.H{
		"meta": gin.H{
			"message": message,
		},
	}
	if len(payload) > 0 {
		response["payload"] = payload[0]
	}
	return response
}

func Hash(payload string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload), 10)
	return string(bytes), err
}

func ValidatePassword(incomingPlainTextPassword string, actualHashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(actualHashedPassword), []byte(incomingPlainTextPassword))
	return err == nil
}

func GenerateJWT(uuid, secret string) (string, error) {
	expiresAt := time.Now().Add(7 * 24 * time.Hour).Unix()

	claims := &JwtClaims{
		UserUuid: uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func FromPtr[T any](obj *T) T {
	if obj == nil {
		var zero T
		return zero
	}
	return *obj
}

func ToPtr[T any](obj T) *T {
	return &obj
}

func PtrCopy[T any](obj *T) *T {
	if obj == nil {
		return nil
	}
	obj2 := *obj
	return &obj2
}
