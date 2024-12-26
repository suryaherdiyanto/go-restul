package token

import (
	"time"

	"github.com/go-restful/app/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	*jwt.RegisteredClaims
}

func GenerateToken(user *model.User, secret string, duration time.Duration) (string, error) {
	ID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.FirstName,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.Email,
			ID:        ID.String(),
		},
	})

	return jwtToken.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secret string) (*UserClaims, error) {
	claims := &UserClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		expirationTime, err := token.Claims.GetExpirationTime()

		if err != nil {
			return nil, err
		}

		if time.Now().After(expirationTime.UTC()) {
			return nil, jwt.ErrTokenExpired
		}

		return []byte(secret), nil
	})

	if err != nil {
		return claims, err
	}

	return claims, nil
}
