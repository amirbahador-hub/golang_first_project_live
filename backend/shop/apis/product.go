package apis

import (
	"digikala/shop"
	"fmt"
	"net/http"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)


func ListProduct(c echo.Context) error {
	// Get team and member from the query string
	products := shop.ListProductService()
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	// db := initializers.DB
	var product shop.ProductRequest
	err := c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	validation_err := validation.ValidateStruct(&product,
		validation.Field(&product.Title, validation.Required, validation.Length(2, 50)),
		validation.Field(&product.Price, validation.Required),
	)

	if validation_err != nil {
		return c.JSON(http.StatusBadRequest, validation_err)
	}

	db_err := shop.CreateProductService(product)
	if db_err != nil {
		fmt.Println(db_err)
		return c.JSON(http.StatusBadRequest, db_err)
	}
	// db.Create(&user)
	return c.JSON(http.StatusOK, true)
}


func GetProduct(c echo.Context) error {
	// Get team and member from the query string
	title := c.QueryParam("title")
	product := shop.GetProductService(title)
	return c.JSON(http.StatusOK, product)
}