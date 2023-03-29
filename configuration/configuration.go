package configuration

import "github.com/spf13/viper"

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

func GetStringConfig(key string) string {
	return configuration.GetString(key)
}

func GetIntConfig(key string) int {
	return configuration.GetInt(key)
}

func GetBoolConfig(key string) bool {
	return configuration.GetBool(key)
}

func GetConfig(key string) any {
	return configuration.Get(key)
}
