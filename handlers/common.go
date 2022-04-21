package handlers

import (
	"errors"
	"fmt"
	"math/big"
	"nft-studio-backend/types"
	"nft-studio-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	utils.Init(100)
}

func getBaseURI(baseUriPrefix string, contractId uint) string {
	return fmt.Sprintf("%v%v/", baseUriPrefix, contractId)
}

func getNftURI(baseURI string, nftId uint) string {
	bi := big.NewInt(0)
	bi.SetUint64(uint64(nftId))
	id := bi.Text(16)
	return fmt.Sprintf("%v%064v.json", baseURI, id)
}

func GetAddress(ctx *gin.Context) (string, error) {
	address, err := utils.GetValueFromJwtToken(ctx, types.JwtKeyAddress)
	if err != nil {
		return "", err
	}

	addressStr, ok := address.(string)
	if !ok {
		return "", errors.New("address is not string")
	}
	return addressStr, nil

}

func PagingResponse(ctx *gin.Context, pResult types.PagingResult, items interface{}) {
	utils.Response(ctx, types.Ok, "", struct {
		PageResult types.PagingResult `json:"page_result"`
		Items      interface{}        `json:"items"`
	}{
		PageResult: pResult,
		Items:      items,
	})
}

func UploadFile(ctx *gin.Context) {
	path := ctx.Request.FormValue("path")
	if path == "" {
		path = viper.GetString("upload.default_path")
	}
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		utils.ResponseErrorWithLog(ctx, types.ErrInvalidForm, err.Error())
		return
	}

	size := fileHeader.Size
	if size > viper.GetInt64("upload.max_size") {
		utils.ResponseErrorWithLog(ctx, types.ErrExceedMaxUpload, "")
		return
	}

	fileName := fileHeader.Filename
	f, err := fileHeader.Open()
	if err != nil {
		utils.ResponseErrorWithLog(ctx, types.ErrOpenUploadFile, err.Error())
		return
	}

	result, err := utils.Uploader.Upload(f, viper.GetString("upload.s3_bucket"), fmt.Sprintf("%v/%v", path, fileName), types.DefaultCacheControl)
	if err != nil {
		utils.ResponseErrorWithLog(ctx, types.ErrUploadFile, err.Error())
		return
	}

	logrus.WithField("path", result.Location).Info("file uploaded")
	utils.Response(ctx, types.Ok, "", result.Location)

}
