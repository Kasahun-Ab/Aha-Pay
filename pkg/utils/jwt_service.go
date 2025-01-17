package utils

import (
	"go_ecommerce/internal/models"
	"io/ioutil"
	"time"
	"github.com/dgrijalva/jwt-go/v4"
	"gopkg.in/yaml.v3"
)

type Config struct {
	API struct {
		SecretKey string `yaml:"secret_key"` 
	} `yaml:"api"`
}

func LoadKey(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}


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

