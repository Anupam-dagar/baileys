package configuration

import (
	"github.com/Anupam-dagar/baileys/constant/types"
	"github.com/spf13/viper"
)

var configuration *viper.Viper

func Init() {
	viper.AutomaticEnv()

	if viper.Get("ENV") == nil {
		viper.Set("ENV", "dev")
	}

	viper.SetConfigName(viper.GetString("ENV"))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	configuration = viper.GetViper()
}

func GetConfiguration() *viper.Viper {
	return configuration
}

func GetStringConfig(key types.ConfigurationKey) string {
	return configuration.GetString(string(key))
}

func GetIntConfig(key types.ConfigurationKey) int {
	return configuration.GetInt(string(key))
}

func GetBoolConfig(key types.ConfigurationKey) bool {
	return configuration.GetBool(string(key))
}

func GetConfig(key types.ConfigurationKey) any {
	return configuration.Get(string(key))
}
