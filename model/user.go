package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
}

type Token struct {
	ID         uint   `json:"id"`
	IdCustomer string `json:"id_customer"`
	Phone      string `json:"phone"`
	*jwt.StandardClaims
}
type Empty struct {
}

type UserView struct {
	Username  string `json:"username"`
	AuthToken string `json:"auth_token"`
}
type UserRegisView struct {
	Username  string `json:"username"`
}
