// Package utils
// @Project:      nft-studio-backend
// @File:          signature.go
// @Author:        eagle
// @Create:        2021/08/11 16:19:36
// @Description:
package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

func Sign(m map[string]interface{}, secret string, signType string) string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var pairs []string
	var stringA string
	for _, k := range keys {
		pairs = append(pairs, fmt.Sprintf("%v=%v", k, m[k]))
	}
	pairs = append(pairs, "key="+secret)
	stringA = strings.Join(pairs, "&")
	logrus.Debugf("stringA: %v\n", stringA)
	switch signType {
	case "HMAC-SHA256":
		return strings.ToUpper(HmacSha256(stringA, secret))
	default:
		return ""
	}
}

func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
