package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUint gets 'key' value from Query
func GetUint(ctx *gin.Context, key string) (uint, error) {
	val := ctx.Query(key)
	if val == "" {
		return 0, fmt.Errorf("%s is empty", key)
	}
	value, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}
