package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func generateJWT() (string, error) {
	var sampleSecretKey = []byte("SecretYouShouldHide")
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
