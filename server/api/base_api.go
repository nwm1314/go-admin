package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go-admin/server/common"
	"go-admin/server/model/request"
	"go-admin/server/model/response"
	"go-admin/server/service"
	"net/http"
)

type BaseApi struct {
}

var captcha *base64Captcha.Captcha

func (*BaseApi) HandleCaptcha(c *gin.Context) {
	var driverDigit = base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	var store = base64Captcha.DefaultMemStore
	captcha = base64Captcha.NewCaptcha(driverDigit, store)
	id, b64s, err := captcha.Generate()
	var captchaResp = response.CaptchaResp{
		id,
		6,
		b64s,
	}
	var data common.Data
	if err != nil {
		data = common.NewData(-999, nil, err.Error())
	} else {
		data = common.NewData(0, captchaResp, "验证码获取成功")
	}
	c.JSON(http.StatusOK, data)
}

func (*BaseApi) HandleLogin(c *gin.Context) {
	var loginData request.Login
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	username := loginData.Username
	password := loginData.Password
	captchaId := loginData.CaptchaId
	captchaContent := loginData.Captcha
	data := service.VerifyUser(username, password, captchaContent, captchaId, captcha)
	c.JSON(http.StatusOK, data)
}
