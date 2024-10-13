package usecase

import (
	"device_management/core/configs"
	"device_management/core/entities"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UseCaseJwt struct{}

func NewUseCaseJwt() *UseCaseJwt {
	return &UseCaseJwt{}
}

func (u *UseCaseJwt) GenToken(userId int64, role int, createdAt int64, userName string) (*entities.UseCaseJwt, error) {
	expirationDuration, err := time.ParseDuration(configs.Get().Expiration)
	if err != nil {
		return nil, fmt.Errorf("invalid token expiration duration: %v", err)
	}

	claims := entities.UserClaims{
		UserName:  userName,
		UserId:    userId,
		Role:      role,
		UpdatedAt: createdAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(configs.Get().Secret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign token: %v", err)
	}

	return &entities.UseCaseJwt{
		Token:     tokenString,
		UserId:    userId,
		Role:      role,
		UpdatedAt: createdAt,
		UserName:  userName,
	}, nil
}

func (u *UseCaseJwt) Verify(tokenString string) (*entities.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &entities.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(configs.Get().Secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(*entities.UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
