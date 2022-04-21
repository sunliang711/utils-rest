// Package utils
// @Project:      nft-studio-backend
// @File:          id_test.go
// @Author:        eagle
// @Create:        2021/08/11 13:22:31
// @Description:
package utils

import (
	"testing"
)

func TestOrderId(t *testing.T) {
	Init(100)
	for i := 0; i < 10; i++ {
		oId := OrderId()
		t.Logf("order id: %v", oId)
	}
}
