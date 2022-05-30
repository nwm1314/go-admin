package api

import (
	"github.com/gin-gonic/gin"
	"go-admin/server/model/request"
	"go-admin/server/service"
	"go-admin/server/utils"
	"net/http"
)

type MenuApi struct {
}

func (*MenuApi) HandleGetMenu(c *gin.Context) {
	id := utils.GetAuthorityId(c)
	if id != "" {
		data := service.GetMenu(id)
		c.JSON(http.StatusOK, data)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "解析token参数出错"})
}

func (*MenuApi) HandleGetMenuList(c *gin.Context) {
	var params request.PageParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data := service.GetMenuList(params.Page, params.PageSize)
	c.JSON(http.StatusOK, data)
}
