package initializers

import (
	"github.com/spf13/viper"
)

const SECRET_KEY="nhdx(=+kek#x+mzyd@qnk!ghu5vc+r($w4*zl^ry6sw1v#2yat"
type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`
	MongoDbUri 	   string `mapstructure:"MONGODB"`

}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	//viper.SetConfigType("env")
	//viper.SetConfigName("")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}