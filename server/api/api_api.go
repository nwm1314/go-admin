package api

import (
	"github.com/gin-gonic/gin"
	"go-admin/server/model/request"
	"go-admin/server/service"
	"net/http"
)

type ApiApi struct {
}

func (*ApiApi) HandleApiList(c *gin.Context) {
	var params request.PageParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data := service.GetApiListPage(params.Page, params.PageSize)
	c.JSON(http.StatusOK, data)
}
