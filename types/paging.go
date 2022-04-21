package types

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PagingParam struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type PagingResult struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

func ExtractPagingQuery(ctx *gin.Context) PagingParam {
	return PagingParam{
		Offset: tryParseIntDefault(ctx.DefaultQuery("offset", "0"), 0),
		Limit:  tryParseIntDefault(ctx.DefaultQuery("limit", "10"), 10),
	}
}

func tryParseIntDefault(v string, d int) int {
	c, err := strconv.Atoi(v)
	if err != nil {
		return d
	}
	return c
}

func GetMaxPagingParam() PagingParam {
	return PagingParam{
		Offset: 0,
		Limit:  math.MaxInt64,
	}
}
