package auth 

import (
	"errors"
	"gorm.io/gorm"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)


type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Name string 
	Role string 
	Password string `gorm:"type:varchar(500)"`
}

func (user *User) hashPassword() error {
	if len(user.Password) < 5 {
		return errors.New("Passowrd must be more than 5 letters")
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(passwordHash)
	return nil
}

func (u *User) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

type MyCustomClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}