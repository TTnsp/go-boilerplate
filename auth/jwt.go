package auth

import (
	"log"
	"time"

	"github.com/ttnsp/go-boilerplate/models"

	"github.com/ttnsp/go-boilerplate/configuration"

	"github.com/golang-jwt/jwt"
)

type TokenPayload struct {
	Name string
}

func GenerateJWT(user models.Users) (string, error) {
	var mySigningKey = []byte(configuration.App.JWT.SecretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Duration(configuration.App.JWT.ExpireIn * 1000000000)).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Print(err)
		return "", err
	}
	return tokenString, nil
}

func IsAuthorized(token string) (bool, *TokenPayload) {
	var mySigningKey = []byte(configuration.App.JWT.SecretKey)

	parsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	// err means token invalid. No err means token valid !
	if err != nil {
		return false, nil
	}

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok && parsed.Valid {
		var payload TokenPayload
		if str, ok := claims["name"].(string); ok {
			payload.Name = str
		} else {
			return false, nil
		}
		return true, &payload
	}

	return false, nil
}
