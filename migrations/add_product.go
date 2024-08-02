package migrations

import (
	"fmt"
	"digikala/logger"
	"digikala/models"
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
	initializers.DB.AutoMigrate(&models.Product{}, &models.Brand{}, &models.Category{})
	fmt.Println("👍 Migration complete")
}