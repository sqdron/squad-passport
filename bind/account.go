package bind

import (
	"github.com/sqdron/squad"
	"fmt"
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

//func AccountBind(c connect.ITransport) *accountBind {
//	return &accountBind{connect: c}
//}
//
func (f *AccountBind) CreateAccount(email string) (int64, error) {
	res, err := f.Invoke(&UserRequest{Email:email})
	fmt.Println(res)
	return 0, err
	//return f.connect.RequestSync("account.create", &UserRequest{Email:email}, 3 * time.Second)
}





