package passport

import (
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type LoginRequest struct {
	User     string
	Password string
}

func (c *passport) Login(r *LoginRequest) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
		"User": r.User,
	})

	tokenString, err := token.SignedString([]byte(c.secret))
	fmt.Printf("%v %v", tokenString, err)
	return "", nil
}
