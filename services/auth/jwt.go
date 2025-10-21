package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT (secret []byte, userID int)(string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":strconv.Itoa(userID),
		"expiredAt" : jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	ss, err:= token.SignedString(secret)
	return ss, err
	
}