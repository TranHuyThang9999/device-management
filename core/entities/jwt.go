package entities

import "github.com/golang-jwt/jwt/v5"

// UseCaseJwt holds the token information and user details
type UseCaseJwt struct {
	Token     string `json:"token"`
	UserName  string `json:"user_name"`
	UserId    int64  `json:"user_id"`
	Role      int    `json:"role,omitempty"`
	UpdatedAt int64  `json:"updated_at"`
}

// UserClaims holds the claims for the JWT token
type UserClaims struct {
	jwt.RegisteredClaims
	UserName  string `json:"user_name"`
	UserId    int64  `json:"user_id"`
	Role      int    `json:"role"`
	UpdatedAt int64  `json:"updated_at"`
}
