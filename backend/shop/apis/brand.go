package apis

import (
	"digikala/initializers"
	"net/http"
	"digikala/shop"
	"github.com/labstack/echo/v4"
)


func ListBrand(c echo.Context) error {
	// Get team and member from the query string
	db := initializers.DB
	var brands []shop.Brand
	db.Find(&brands)
	return c.JSON(http.StatusOK, brands)
}

func CreateBrand(c echo.Context) error {
	// Get team and member from the query string
	db := initializers.DB
	var brand shop.Brand
	err := c.Bind(&brand)
	if err != nil{
		panic(err)
	}
	db.Create(&brand)
	return c.JSON(http.StatusOK, brand)
}