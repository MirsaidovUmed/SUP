package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(secretKey string, id int) (accessToken string, err error) {
	fmt.Println(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
			"user_id": id,
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
