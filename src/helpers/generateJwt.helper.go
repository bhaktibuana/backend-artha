package helpers

import (
	"api-artha/src/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(payload jwt.Claims, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	if expiresIn > 0 {
		token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(expiresIn).Unix()
	}
	signedToken, err := token.SignedString([]byte(configs.AppConfig().JwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
