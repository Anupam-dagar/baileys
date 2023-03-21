package configuration

import "github.com/spf13/viper"

var configuration *viper.Viper

func Init(config interface{}) {
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

	err = viper.Unmarshal(&config)
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
