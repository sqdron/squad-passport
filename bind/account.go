package bind

import (
	"github.com/sqdron/squad"
)

type UserRequest struct {
	Email string
}

type IAccountBind interface {
	CreateAccount(email string) (int64, error)
	//	//CreateAccount(request *UserRequest) (*UserResponce, error)
	//	//GetAccount(request *UserRequest) (*UserResponce, error)
}
//
type AccountBind struct {
	squad.SquadBinder
}

func (f *AccountBind) CreateAccount(email string) (int64, error) {
	var res int64
	err := f.Invoke("create", &UserRequest{Email:email}, &res)
	return res, err
}





