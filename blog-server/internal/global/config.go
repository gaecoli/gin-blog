package global

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Server struct {
		Mode          string
		Port          string
		DbAutoMigrate bool
		DbLogMode     string
	}

	Log struct {
		Level  string
		Prefix string
		Format string
		Path   string
	}

	MySql struct {
		Host     string
		Port     string
		Config   string
		Dbname   string
		Username string
		Password string
	}
}

var Conf *Config

// GetConfig 获取配置文件
func GetConfig() *Config {
	if Conf == nil {
		log.Panic("config file not init...")
		return nil
	}
	return Conf
}

func ReadConfig(path string) *Config {
	v := viper.New()
	v.SetConfigFile(path)
	v.AutomaticEnv()                                   // 可以使用环境变量
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // xxx_xx => xxx.xx

	if err := v.ReadInConfig(); err != nil {
		panic("Config file read failed: " + err.Error())
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic("配置文件反序列化失败: " + err.Error())
	}

	log.Println("配置文件内容加载成功：", path)
	return Conf
}

// DbDSN 返回数据库 dsn
func (*Config) DbDSN() string {
	conf := Conf.MySql
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Dbname, conf.Config)
}
