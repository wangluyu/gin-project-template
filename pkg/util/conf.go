package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configFile []byte
var appConfigFile []byte

type AppConf struct {
	Evn       string `yaml:"evn"`
	JwtExpire int    `yaml:"jwtExpire"`
}

type Conf struct {
	Mysql MysqlConf `yaml:"mysql"`
	Redis RedisConf `yaml:"redis"`
}

type MysqlConf struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Timeout  int    `yaml:"timeout"`
}

type RedisConf struct {
	Host string `yaml:"host"`
}

func FetchAppConf() (c *AppConf, err error) {
	err = yaml.Unmarshal(appConfigFile, &c)
	return c, err
}

func FetchConf() (c *Conf, err error) {
	err = yaml.Unmarshal(configFile, &c)
	return c, err
}

func FetchMysqlConf() (c MysqlConf, err error) {
	conf, err := FetchConf()
	if err == nil {
		c = conf.Mysql
	}
	return c, err
}

func FetchRedisConf() (c RedisConf, err error) {
	conf, err := FetchConf()
	if err == nil {
		c = conf.Redis
	}
	return c, err
}

func init() {
	var err error
	appConfigFile, err = ioutil.ReadFile("./config/app.yaml")
	if err == nil {
		appConf, err := FetchAppConf()
		if err == nil {
			fileName := "./config/" + appConf.Evn + ".yaml"
			configFile, err = ioutil.ReadFile(fileName)
		}
	}
}
