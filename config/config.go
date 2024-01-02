package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const NAMESPACE = "MOVIERST"

type Configuration struct {
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort int    `mapstructure:"SERVER_PORT"`

	MySQLHost     string `mapstructure:"MYSQL_HOST"`
	MySQLPort     int    `mapstructure:"MYSQL_PORT"`
	MySQLUsername string `mapstructure:"MYSQL_USERNAME"`
	MySQLPassword string `mapstructure:"MYSQL_PASSWORD"`
	MySQLDbName   string `mapstructure:"MYSQL_DB_NAME"`
}

func (cf Configuration) ServerAddress() string {
	return fmt.Sprintf("%s:%d", cf.ServerHost, cf.ServerPort)
}

func (cf Configuration) MySqlDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cf.MySQLUsername,
		cf.MySQLPassword,
		cf.MySQLHost,
		cf.MySQLPort,
		cf.MySQLDbName,
	)
}

func Get(path string) (Configuration, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.SetEnvPrefix(NAMESPACE)
	viper.AutomaticEnv()

	var cfg Configuration
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("err loading .env err: %v", err)
		return Configuration{}, err
	}

	if err := viper.UnmarshalExact(&cfg); err != nil {
		log.Printf("err loading .env err: %v\n", cfg)

		return Configuration{}, err
	}

	return cfg, nil
}
