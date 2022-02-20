package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type conf struct {
	Url              string `yaml:"url"`
	UserName         string `yaml:"userName"`
	Password         string `yaml:"password"`
	DbName           string `yaml:"dbname"`
	Port             string `yaml:"post"`
	MaxIdleConnCount int    `yaml:"max_idle_conn_count"`
	MaxOpenConnCount int    `yaml:"max_open_conn_count"`
}

var Conf conf

func init() {
	yamlFile, err := ioutil.ReadFile("../resources/application.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetDBConnectUrl() string {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Conf.UserName,
		Conf.Password,
		Conf.Url,
		Conf.Port,
		Conf.DbName,
	)
	return dbUrl
}

func GetMaxIdleConnCount() int {
	return Conf.MaxIdleConnCount
}

func GetMaxOpenConnCount() int {
	return Conf.MaxOpenConnCount
}

func GetConnMaxLifetime() time.Duration {
	return time.Minute
}
