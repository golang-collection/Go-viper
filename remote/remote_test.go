package main

import "testing"

/**
* @Author: super
* @Date: 2020-09-03 10:37
* @Description:
**/

func TestGetMysqlUrl(t *testing.T) {
	s, _ := GetMysqlUrl()
	t.Log(s)
}