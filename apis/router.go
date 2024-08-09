package apis

import (
	"github.com/labstack/echo/v4"
	shopApi "digikala/shop/apis"
	authApi "digikala/auth"
)


func GetRouter() *echo.Echo{
	e := echo.New()
	brandRouter := e.Group("/brands/")
	brandRouter.POST("", shopApi.CreateBrand)
	brandRouter.GET("", shopApi.ListBrand)

	userRouter := e.Group("/users/")
	userRouter.POST("", authApi.CreateUser)
	userRouter.GET("", authApi.ListUser)
	userRouter.POST("login/", authApi.Login)
	userRouter.POST("verify/", authApi.Verify)
	return e

}