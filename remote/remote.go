package main

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-09-03 10:33
* @Description:
**/

var logger = zap.NewExample()

func init() {
	err := viper.AddRemoteProvider("consul", "localhost:8500", "config")
	if err != nil {
		logger.Error("read config",zap.Error(err))
		return
	}
	viper.SetConfigType("json") // Need to explicitly set this to json
	if err := viper.ReadRemoteConfig(); err != nil {
		logger.Error("read config",zap.Error(err))
		return
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