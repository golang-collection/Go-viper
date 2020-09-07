package normal

import (
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-09-07 08:03
* @Description:
**/

func TestGetMysqlUrl(t *testing.T) {
	if err := Init("local"); err != nil {
		t.Error(err)
		t.FailNow()
	}
	s, _ := GetMysqlUrl()
	t.Log(s)
}

func TestGetRedisUrl(t *testing.T) {
	if err := Init("local"); err != nil {
		t.Error(err)
		t.FailNow()
	}
	s, _ := GetRedisUrl()
	t.Log(s)
}

func TestGetRabbitMQUrl(t *testing.T) {
	if err := Init("local"); err != nil {
		t.Error(err)
		t.FailNow()
	}
	s, _ := GetRabbitMQUrl()
	t.Log(s)
}

func TestConfigFileChanged(t *testing.T){
	ch := make(chan string)

	if err := Init("local"); err != nil {
		t.Error(err)
		t.FailNow()
	}
	s, _ := GetMysqlUrl()
	fmt.Println(s)

	<-ch
}