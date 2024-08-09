package auth

import (
	"digikala/initializers"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createUser(userRequest UserRequest) (User, error){
	db := initializers.DB
	user := User{Name:userRequest.Name, Role: userRequest.Role, Password: userRequest.Password}
	user.hashPassword()
	result := db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func getUsers() []User{
	db := initializers.DB
	var users []User
	db.Find(&users)
	return users
}

func getUser(name string) (User, error){
	db := initializers.DB
	var user User
	 
	result := db.Where("Name = ?", name).First(&user)
	if result.Error != nil{
		return User{}, result.Error
	}
	return user, nil
}

func login(request LoginRequest) (LoginResponse, error){
	user, err := getUser(request.Name)
	if user.checkPassword(request.Password) == nil{
		claims := MyCustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			},
			Name : user.Name,
			Role:   user.Role,
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		access_toekn, _:= token.SignedString([]byte(initializers.SECRET_KEY))
		
		return LoginResponse{Token: access_toekn}, nil
	}
	return LoginResponse{}, err
}

func verify(request VerifyRequest) (VerifyResponse, error){
	token, err := jwt.ParseWithClaims(request.Token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(initializers.SECRET_KEY), nil
	})
	if _, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return VerifyResponse{IsValid: true}, err 
	} else {
		return VerifyResponse{IsValid: false}, err
	}
}