package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/server/utils"
	"log"
	"net/http"
	"strings"
)

func TokenMiddle() gin.HandlerFunc {
	return func(con *gin.Context) {
		// 根据headers 中的 Authorization，判断用户是否登录
		path := con.Request.URL.Path
		if strings.Index(path, "login") < 0 && !strings.Contains(path, "captcha") {
			// 进行token验证
			token := con.Request.Header.Get("x-token")
			if token == "" {
				con.JSON(
					http.StatusOK,
					gin.H{
						"code": 405,
						"msg":  "无权访问",
					},
				)
				con.Abort()
				return
			}
			claims, err := utils.ParseToken(token)
			if err != nil {
				log.Println(err.Error())
				con.JSON(
					http.StatusOK,
					gin.H{
						"code": 405,
						"msg":  "无权访问",
					},
				)
				con.Abort()
				return
			}
			con.Set("claims", *claims)
			con.Next()
		} else {
			con.Next()
		}
	}
}
