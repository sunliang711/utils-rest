package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"nft-studio-backend/service/user"
	"nft-studio-backend/types"
	"nft-studio-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get(viper.GetString("jwt.header_name"))
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "need jwt token",
		})
		c.Abort()
		return
	}
	_, err := utils.ParseJwtToken(token, viper.GetString("jwt.key"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  err.Error(),
		})
		logrus.Errorf("Parse jwt token error: %v", err.Error())
		c.Abort()
		return
	}
	c.Next()
}

func User(ctx *gin.Context) {
	if err := auth(ctx, types.RoleUser); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": types.ErrUserRole,
			"msg":  types.ErrorMsg(types.ErrUserRole, err.Error()),
		})
		ctx.Abort()
	} else {
		ctx.Next()
	}
}

func Admin(ctx *gin.Context) {
	if err := auth(ctx, types.RoleUser); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": types.ErrAdminRole,
			"msg":  types.ErrorMsg(types.ErrAdminRole, err.Error()),
		})
		ctx.Abort()
	} else {
		ctx.Next()
	}

}

func auth(ctx *gin.Context, expectedRole int) error {
	token := ctx.Request.Header.Get(viper.GetString("jwt.header_name"))
	if token == "" {
		return errors.New("need jwt token")
	}

	role, err := utils.GetValueFromJwtToken(ctx, types.JwtKeyRole)
	if err != nil {
		return err
	}

	// logrus.Debugf("role id: %v", role)
	if fmt.Sprintf("%v", role) != fmt.Sprintf("%v", expectedRole) {
		switch expectedRole {
		case types.RoleAdmin:
			return errors.New("not admin")
		case types.RoleUser:
			return errors.New("not user")
		default:
			return errors.New("expected role invalid")
		}
	}

	address, err := utils.GetValueFromJwtToken(ctx, types.JwtKeyAddress)
	if err != nil {
		return err
	}

	_, err = user.NewService().QueryUser(address.(string))
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return errors.New("no such user")
		} else {
			return err
		}
	}
	return nil
}
