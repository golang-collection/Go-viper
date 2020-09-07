package main

import (
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-09-03 10:37
* @Description:
**/

func TestGetMysqlUrl(t *testing.T) {
	if err := Init("remote"); err != nil {
		t.Error(err)
		t.FailNow()
	}
	s, _ := GetMysqlUrl()
	t.Log(s)
}

func TestConfigFileChanged(t *testing.T) {
	ch := make(chan string)

	if err := Init("remote"); err != nil {
		t.Error(err)
		t.FailNow()
	}
	s, _ := GetMysqlUrl()
	fmt.Println(s)

	<-ch
}
