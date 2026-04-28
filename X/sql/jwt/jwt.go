package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256,
	jwt.MapClaims{
		"id": id,
		"username": username,
		"exp": time.Now().Add(60*time.Minute).Unix(),
	})

	key:=[]byte(secretKey)
	token_str,err:=token.SignedString(key)
	if err!=nil{
		return "", err
	}
	return token_str, nil
}

func ValidateToken(tokenStr,secretKey string, withClaimValidation bool) (*jwt.Token, error) {
	var (
		key = []byte(secretKey)
		claims= jwt.MapClaims{}
		token *jwt.Token
		err error
	)

	if withClaimValidation {
		token,err= jwt.ParseWithClaims(tokenStr,claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
	} else {
		token,err= jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return key,nil
		}, jwt.WithoutClaimsValidation())
	}

	if err!=nil{
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	userID:=claims["id"].(float64)
	username:=claims["username"].(string)
	claims["id"]=int64(userID)
	claims["username"]=username
	return token, nil
}