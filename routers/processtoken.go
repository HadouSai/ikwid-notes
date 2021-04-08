package routers

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/HadouSai/ikwid-notes/db"
	"github.com/HadouSai/ikwid-notes/models"
)

var Email string
var IdUser string

// ProcessToken:
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myUltraSuperSecretKey := []byte("IlOvEuReAcT<3")
	claims := &models.Claim{}

	tok, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) { return myUltraSuperSecretKey, nil })

	if err == nil {
		_, exists, _ := db.CheckUserExist(claims.Email)
		if exists {
			Email = claims.Email
			IdUser = claims.ID.Hex()
		}
		return claims, exists, IdUser, nil
	}

	if !tok.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err

}
