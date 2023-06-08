package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	PORT           int    `mapstructure:"PORT"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	// MongoURI       string `mapstructure:"MONGO_URI"`
	DBName string `mapstructure:"DB_NAME"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`

	//auth configs
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env: ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	if env.ServerAddress == "" {
		log.Fatal("SERVER_ADDRESS can't be empty in .env file ")
	} else {
		log.Println("server address is : ", env.ServerAddress)
	}

	if env.PORT <= 0 {
		log.Fatal("PORT can't be zero in .env file ")
	} else {
		log.Println("Application running on port  : ", env.PORT)
	}

	if env.PORT <= 0 {
		log.Fatal("CONTEXT_TIMEOUT should be a postive integer .env file ")
	} else {
		log.Println("Context timeout set to the value : ", env.ContextTimeout)
	}

	if env.DBName == `` {
		log.Fatal("DB_NAME cannot be an empty string .env file ")
	}

	return &env
}
