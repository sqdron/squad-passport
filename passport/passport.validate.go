package passport

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func (c *passport) Validate(r *TokenRequest) error {
	fmt.Println("Passport validations", r)

	token, err := jwt.Parse(r.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.secret), nil
	})
	if token.Valid {
		fmt.Println("Token is valid")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors & jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors & (jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	return err
}

