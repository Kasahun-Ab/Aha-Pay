package utils

import (
	"go_ecommerce/internal/models"
	"time"
	"github.com/dgrijalva/jwt-go/v4"

)

func GenerateJWT(user models.User, secretKey string) (string, error) {

	claims := jwt.MapClaims{
		"id":        user.ID,
		"username":  user.Username,
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"exp":       time.Now().Add(24 * time.Hour).Unix(), 
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey)) 
}


func ParseToken(tokenStr, secretKey string) (jwt.MapClaims, error) {
	
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		
		return []byte(secretKey), nil

	})

	if err != nil {
		return nil, err
	}


	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid 
	}
	return claims, nil
}

