package normal

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

/**
* @Author: super
* @Date: 2020-09-07 07:59
* @Description: 读取本地配置文件
**/

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	viper.SetConfigFile("/Users/super/develop/Go-viper/config/config.json") //文件名
	err := viper.ReadInConfig() // 会查找和读取配置文件
	if err != nil {             // Handle errors reading the config file
		return err
	}
	return nil
}

func GetMysqlUrl() (string, error) {
	mysqlHost := viper.GetString("mysql.host")
	mysqlUser := viper.GetString("mysql.user")
	mysqlPassword := viper.GetString("mysql.password")
	mysqlDBName := viper.GetString("mysql.db_name")
	mysqlURL := mysqlUser + ":" + mysqlPassword + "@(" + mysqlHost + ")/" + mysqlDBName
	return mysqlURL, nil
}

func GetRedisUrl() (string, error) {
	redisURL := viper.GetString("redis.host")
	return redisURL, nil
}

func GetRabbitMQUrl() (string, error) {
	mqHost := viper.GetString("rabbitmq.host")
	mqUser := viper.GetString("rabbitmq.user")
	mqPassword := viper.GetString("rabbitmq.password")
	mqURL := "amqp://" + mqUser + ":" + mqPassword + "@"+mqHost+"/"
	return mqURL, nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s\n", e.Name)
	})
}