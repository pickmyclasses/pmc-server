package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	Secret              = []byte("pmc-jwt-secrete")
	TokenExpireDuration = time.Hour * 24
)

type UniqueClaims struct {
	UserID    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.StandardClaims
}

func GenToken(userID int64, firstName string, lastName string) (string, error) {
	c := UniqueClaims{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "pmc",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

func ParseToken(tokenStr string) (*UniqueClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UniqueClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UniqueClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
