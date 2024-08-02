package apis

import (
	"github.com/labstack/echo/v4"
)


func GetRouter() *echo.Echo{
	e := echo.New()
	brandRouter := e.Group("/brands/")
	brandRouter.POST("", createBrand)
	brandRouter.GET("", listBrand)
	return e
	// brandRouter.GET(":id", getBrand)
	// brandRouter.PUT(":id", UpdateBrand)
	// brandRouter.DELETE(":id", DeleteBrand)

}