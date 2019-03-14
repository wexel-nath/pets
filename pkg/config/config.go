package config

import (
	"github.com/spf13/viper"
)

func Configure() {
	viper.SetEnvPrefix("PETS")

	// Postgres
	viper.SetDefault("DB_HOST", "")
	viper.BindEnv("DB_HOST")
	viper.SetDefault("DB_NAME", "pets")
	viper.BindEnv("DB_NAME")
	viper.SetDefault("DB_USER", "postgres")
	viper.BindEnv("DB_USER")
	viper.SetDefault("DB_PASS", "")
	viper.BindEnv("DB_PASS")
}

func GetDBHost() string {
	return viper.GetString("DB_HOST")
}

func GetDBName() string {
	return viper.GetString("DB_NAME")
}

func GetDBUser() string {
	return viper.GetString("DB_USER")
}

func GetDBPass() string {
	return viper.GetString("DB_PASS")
}
