package apis

import (
	"digikala/initializers"
	"digikala/models"
	"net/http"

	"github.com/labstack/echo/v4"
)


func listBrand(c echo.Context) error {
	// Get team and member from the query string
	db := initializers.DB
	var brands []models.Brand
	db.Find(&brands)
	return c.JSON(http.StatusOK, brands)
}

func createBrand(c echo.Context) error {
	// Get team and member from the query string
	db := initializers.DB
	var brand models.Brand
	err := c.Bind(&brand)
	if err != nil{
		panic(err)
	}
	db.Create(&brand)
	return c.JSON(http.StatusOK, brand)
}