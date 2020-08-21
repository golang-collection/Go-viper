package main

import (
	"fmt"
	"log"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-21 10:00
* @Description:
**/

func TestGetMysqlUrl(t *testing.T) {
	mysqlUrl, err := GetMysqlUrl()
	if err != nil{
		log.Println(err)
	}
	fmt.Println(mysqlUrl)
}

func TestGetRedisUrl(t *testing.T) {
	mysqlUrl, err := GetRedisUrl()
	if err != nil{
		log.Println(err)
	}
	fmt.Println(mysqlUrl)
}

func TestGetRabbitMQUrl(t *testing.T) {
	mysqlUrl, err := GetRabbitMQUrl()
	if err != nil{
		log.Println(err)
	}
	fmt.Println(mysqlUrl)
}