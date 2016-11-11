package main

import (
	"github.com/sqdron/squad"
	"fmt"
	"github.com/sqdron/squad-passport/passport"
	"github.com/sqdron/squad-passport/services"
	"github.com/sqdron/squad-passport/bind"
)

type Options struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	Secret     string
}

func main() {

	o := &Options{}
	client := squad.Client("squad.passport", o)

	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		o.DbHost, o.DbPort, o.DbUser, o.DbName, o.DbPassword)

	accoundService := client.Bind.ToController("account.provider", &bind.AccountBind{}).(bind.IAccountBind)
	fmt.Println("accoundService", accoundService)
	controller := passport.PassportController(services.AuthService(connection), o.Secret)
	accoundService.CreateAccount("dddd@ddd.dd")
	client.Api.Controller(controller)
	<-client.Run()
}
