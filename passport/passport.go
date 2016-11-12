package passport

import (
	"github.com/sqdron/squad-passport/services"
	"github.com/sqdron/squad-passport/bind"
)

type passport struct {
	secret  string
	auth    services.IAuthService
	account bind.IAccountBind
}

type TokenRequest struct {
	Token string
}

func PassportController(auth services.IAuthService, account bind.IAccountBind, secret string) *passport {
	return &passport{secret, auth, account}
}
