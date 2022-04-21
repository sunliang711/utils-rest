package utils

import (
	"io"
	"net/http"

	"nft-studio-backend/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Response(ctx *gin.Context, code int, additionMsg string, data interface{}) {
	// get message from error table
	message := types.ErrorMsg(code, additionMsg)
	ctx.JSON(
		http.StatusOK,
		types.Resp{
			Code: code,
			Msg:  message,
			Data: data,
		})
}

func ResponseErrorWithLog(ctx *gin.Context, code int, additionMsg string) {
	message := types.ErrorMsg(code, additionMsg)
	logrus.Error(message)
	ctx.JSON(
		http.StatusOK, types.Resp{
			Code: code,
			Msg:  message,
		})
}

// HandleError deal with error in handlers
func HandleError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	logrus.Error(err.Error())
	if err == io.EOF {
		Response(ctx, types.ErrEOF, "", nil)
		return true
	}

	Response(ctx, types.ErrGeneral, err.Error(), nil)
	return true
}
