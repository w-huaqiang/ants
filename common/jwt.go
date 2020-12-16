package common

import (
	"time"

	"bjzdgt.com/ants/model"
	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte("bca_bjzdgt_ant")

// Claims is struct for jwt
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// ReleaseToken for release token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "bjzdgt.com",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtkey)

	if err != nil {
		return "", err
	}

	return tokenString, err
}

//ParseToken for middleware
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (i interface{}, err error) {
			return jwtkey, nil
		})

	return token, claims, err
}
