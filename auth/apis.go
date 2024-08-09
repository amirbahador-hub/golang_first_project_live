package auth

import (
	"fmt"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)



func CreateUser(c echo.Context) error {
	// db := initializers.DB
	var user UserRequest
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	validation_err := validation.ValidateStruct(&user,
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Field(&user.Name, validation.Required, validation.Length(5, 20)),
		validation.Field(&user.Password, validation.Required, validation.Length(5, 20)),
		// validation.Field(&c.Gender, validation.In("Female", "Male")),
		// validation.Field(&c.Email, validation.Required, is.Email),
		// Validate Address using its own validation rules
		validation.Field(&user.Role),
	)

	if validation_err != nil {
		return c.JSON(http.StatusBadRequest, validation_err)
	}

	user_obj, db_err := createUser(user)
	if db_err != nil {

		return c.JSON(http.StatusBadRequest, db_err)
	}
	// db.Create(&user)
	return c.JSON(http.StatusOK, user_obj)
}

func ListUser(c echo.Context) error {
	users := getUsers()
	return c.JSON(http.StatusOK, users)
}

func Login(c echo.Context) error {
	var request LoginRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	validation_err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required, validation.Length(5, 20)),
		validation.Field(&request.Password, validation.Required, validation.Length(5, 20)),
	)

	if validation_err != nil {
		return c.JSON(http.StatusBadRequest, validation_err)
	}

	response, db_err := login(request)
	if db_err != nil {
		return c.JSON(http.StatusBadRequest, db_err)
	}
	return c.JSON(http.StatusOK, response)
}

func Verify(c echo.Context) error {
	var request VerifyRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	validation_err := validation.ValidateStruct(&request,
		validation.Field(&request.Token, validation.Required),
	)

	if validation_err != nil {
		return c.JSON(http.StatusBadRequest, validation_err)
	}

	response, db_err := verify(request)
	if db_err != nil {
		return c.JSON(http.StatusBadRequest, db_err)
	}
	return c.JSON(http.StatusOK, response)
}