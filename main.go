package main


import(
	"github.com/spf13/viper"
	"go.uber.org/zap"
)
/**
* @Author: super
* @Date: 2020-08-21 09:57
* @Description:
**/

var logger = zap.NewExample()

func init() {
	viper.SetConfigFile("config/config.json") //文件名
	err := viper.ReadInConfig() // 会查找和读取配置文件
	if err != nil {             // Handle errors reading the config file
		logger.Error("viper read config error", zap.Error(err))
	}
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