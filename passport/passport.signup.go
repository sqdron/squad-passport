package passport

import "fmt"

type SignupRequest struct {
	Email string
}

func (c *passport) Signup(r SignupRequest) error {
	fmt.Println("c.account.CreateAccount(r.Email)", r.Email)
	res, err := c.account.CreateAccount(r.Email)
	fmt.Println(res, err)
	return err
}

