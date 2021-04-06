package jwt

import (
	"time"

	"github.com/HadouSai/ikwid-notes/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* GenerateJWT generate a jwt  */
func GenerateJWT(user *models.User) (string, error) {
	myUltraSuperSecretKey := []byte("IlOvEuReAcT<3")

	payload := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"_id":   user.ID.Hex(),
		"exp":   time.Now().Add(time.Hour * 16).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(myUltraSuperSecretKey)

	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}
