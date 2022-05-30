package api

import (
	"github.com/gin-gonic/gin"
	"go-admin/server/model/request"
	"go-admin/server/service"
	"go-admin/server/utils"
	"net/http"
)

type UserApi struct {
}

func (*UserApi) GetUserInfo(c *gin.Context) {
	id := utils.GetUserId(c)
	data := service.GetUserInfo(id)
	c.JSON(http.StatusOK, data)
}

func (*UserApi) GetUserList(c *gin.Context) {
	var params request.PageParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data := service.GetUserList(params.Page, params.PageSize)
	c.JSON(http.StatusOK, data)
}
