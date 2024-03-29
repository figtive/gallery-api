package configs

import "github.com/spf13/viper"

var AppConfig Config

type Config struct {
	Port   int
	Debug  bool
	Secret string

	StaticBaseDir string

	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string

	GoogleClientID string
}

func InitializeConfig() {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	AppConfig.Port = viper.GetInt("PORT")
	AppConfig.Debug = viper.GetBool("DEBUG")
	AppConfig.Secret = viper.GetString("SECRET")

	AppConfig.StaticBaseDir = viper.GetString("STATIC_BASE_DIR")

	AppConfig.DbHost = viper.GetString("DB_HOST")
	AppConfig.DbPort = viper.GetInt("DB_PORT")
	AppConfig.DbUser = viper.GetString("DB_USER")
	AppConfig.DbPassword = viper.GetString("DB_PASSWORD")
	AppConfig.DbName = viper.GetString("DB_NAME")
	AppConfig.GoogleClientID = viper.GetString("GOOGLE_CLIENT_ID")
}
