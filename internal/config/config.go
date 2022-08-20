package config

import "github.com/spf13/viper"

type Config struct {
	Port             string `mapstructure:"PORT"`
	AuthSvcUrl       string `mapstructure:"AUTH_SVC_URL"`
	PassengerSvcUrl  string `mapstructure:"PASSENGER_SVC_URL"`
	DriverSvcUrl     string `mapstructure:"DRIVER_SVC_URL"`
	CallCenterSvcUrl string `mapstructure:"CALLCENTER_SVC_URL"`
	BookingSvcUrl    string `mapstructure:"BOOKING_SVC_URL"`
	LocationSvcUrl   string `mapstructure:"LOCATION_SVC_URL"`

	SecretKey string `mapstructure:"SECRET_KEY"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
