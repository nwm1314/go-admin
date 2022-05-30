package api

import (
	"github.com/gin-gonic/gin"
	"go-admin/server/model/request"
	"go-admin/server/service"
	"net/http"
)

type AuthorityApi struct {
}

func (*AuthorityApi) GetAuthorityList(c *gin.Context) {
	var params request.PageParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data := service.GetAuthorityList(params.Page, params.PageSize)
	c.JSON(http.StatusOK, data)
}
