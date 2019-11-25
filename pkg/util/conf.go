package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configFile []byte
var appConfigFile []byte

type AppConf struct {
	Evn         string          `yaml:"evn"`
	JwtExpire   int             `yaml:"jwtExpire"`
	MailProduct MailProductConf `yaml:"mailProduct"`
	Mail        MailConf        `yaml:"mail"`
	Tencent     TencentConf     `yaml:"tencent"`
	Sms         SmsConf         `yaml:"sms"`
}

type SmsConf struct {
	Sign     string `yaml:"sign"`
	SdkAppid string `yaml:"sdkAppid"`
}

type TencentConf struct {
	SecretId  string `yaml:"secretId"`
	SecretKey string `yaml:"secretKey"`
}
type MailConf struct {
	From     string `yaml:"from"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type MailProductConf struct {
	Name      string `yaml:"name"`
	Link      string `yaml:"link"`
	Logo      string `yaml:"logo"`
	Copyright string `yaml:"copyright"`
}

type Conf struct {
	LogPath string    `yaml:"logPath"`
	Mysql   MysqlConf `yaml:"mysql"`
	Redis   RedisConf `yaml:"redis"`
}

type MysqlConf struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
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
