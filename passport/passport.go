package passport

import (
	"github.com/sqdron/squad-passport/services"
)

type passport struct {
	secret string
	auth services.IAuthService
}

type TokenRequest struct {
	Token string
}

func PassportController(auth services.IAuthService, secret string) *passport {
	return &passport{auth:auth, secret:secret}
}