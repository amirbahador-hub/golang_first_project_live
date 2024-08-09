package migrations

import (
	"fmt"
	"digikala/logger"
	"digikala/shop"
	"digikala/auth"
	"digikala/initializers"
)

func Setup(){
	myslog := logger.GetLogger()
	config, err := initializers.LoadConfig(".")
	if err != nil {
		myslog.Error("🚀 Could not load environment variables", err)
	}
	initializers.ConnectDB(&config)
}

func RunMigrations() {
	Setup()
	initializers.DB.AutoMigrate(&shop.Product{}, &shop.Brand{}, &shop.Category{}, &auth.User{})
	fmt.Println("👍 Migration complete")
}