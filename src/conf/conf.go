package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Url              string `mapstructure:"url"`
	UserName         string `mapstructure:"userName"`
	Password         string `mapstructure:"password"`
	DbName           string `mapstructure:"dbname"`
	Port             string `mapstructure:"post"`
	MaxIdleConnCount int    `mapstructure:"max_idle_conn_count"`
	MaxOpenConnCount int    `mapstructure:"max_open_conn_count"`
}

const appConfigFilePath = "./src/conf/application.yaml"

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(appConfigFilePath)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = v.Unmarshal(&config)
	return config, nil
}

func (cfg *Config) GetDBConnectUrl() string {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.UserName,
		cfg.Password,
		cfg.Url,
		cfg.Port,
		cfg.DbName,
	)
	return dbUrl
}

func (cfg *Config) GetMaxIdleConnCount() int {
	return cfg.MaxIdleConnCount
}

func (cfg *Config) GetMaxOpenConnCount() int {
	return cfg.MaxOpenConnCount
}

func (cfg *Config) GetConnMaxLifetime() time.Duration {
	return time.Minute
}
