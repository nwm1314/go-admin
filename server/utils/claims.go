package utils

import (
	"github.com/gin-gonic/gin"
	"go-admin/server/model/request"
	"log"
)

func GetUserId(c *gin.Context) uint {
	claims, _ := c.Get("claims")
	if claims != nil {
		if baseClaims, ok := claims.(request.BaseClaims); ok {
			return baseClaims.ID
		}
	}
	return 0
}

func GetAuthorityId(c *gin.Context) string {
	claims, _ := c.Get("claims")
	if claims != nil {
		if baseClaims, ok := claims.(request.BaseClaims); ok {
			log.Println(baseClaims.AuthorityId)
			return baseClaims.AuthorityId
		}
	}
	return ""
}

func GetClaims(c *gin.Context) *request.BaseClaims {
	claims, _ := c.Get("claims")
	if claims != nil {
		if baseClaims, ok := claims.(request.BaseClaims); ok {
			return &baseClaims
		}
	}
	return nil
}
