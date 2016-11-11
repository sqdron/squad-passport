package services

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/sqdron/squad-passport/model"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type authService struct {
}

type IAuthService interface {
	//IService
	Token(*model.User) string
}

func (s *authService) Token(user *model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
		"User": user.Email,
	})

	tokenString, err := token.SignedString([]byte("mySigningKey"))
	fmt.Printf("%v %v", tokenString, err)
	return tokenString
}

func AuthService(connection string) *authService {
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()
	if (err != nil) {
		fmt.Errorf("DBError", err)
	}
	fmt.Println(db)
	db.AutoMigrate(&model.User{})
	return &authService{}
}