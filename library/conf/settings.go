package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	settings *viper.Viper
)

type GinConfig struct {
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
}

type PostgreSQL struct {
	User        string
	Password    string
	Host        string
	Port        int
	Name        string
	TablePrefix string
}

type RedisConfig struct {
	Host string
	Port int
}

type ElasticSearchConfig struct {
	Host string
	Port int
}

type SentryConfig struct {
	Url string
}

type AppConfig struct {
	RunMode   string
	SecretKey string

	Gin           GinConfig
	PostgreSQL    PostgreSQL
	Redis         RedisConfig
	ElasticSearch ElasticSearchConfig
	Sentry        SentryConfig
}

var C AppConfig

func Init(env string) {
	var err error
	settings = viper.New()
	settings.SetConfigType("yaml")
	settings.SetConfigName(env)
	settings.AddConfigPath("./library/conf/")
	err = settings.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = settings.Unmarshal(&C)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
}

func GetConfig() AppConfig {
	return C
}
