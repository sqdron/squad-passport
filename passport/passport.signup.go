package passport

import "fmt"

type SignupRequest struct {
	Email string
}

func (c *passport) Signup(r SignupRequest) error {
	fmt.Println("Signup", r)
	return nil
}

