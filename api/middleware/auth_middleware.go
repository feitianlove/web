package middleware

import (
	"fmt"
	"github.com/feitianlove/golib/common/ecode"
	"github.com/feitianlove/web/auth"
	"github.com/gin-gonic/gin"
)

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		url := c.Request.URL.String()
		var role string
		if method == "GET" {
			role = c.Query("name")
		} else {
			role = c.PostForm("name")
		}
		if len(role) == 0 {
			ecode.RespErrCode(c, ecode.AuthEmpty, "don't get username")
			c.Abort()
		}
		ok, err := auth.CheckPolicy(role, url, method)
		if err != nil {
			ecode.RespErrCode(c, ecode.AuthDenied, err.Error())
		}
		if ok {
			c.Next()
		} else {
			ecode.RespErrCode(c, ecode.AuthDenied,
				fmt.Sprintf("%s don't have permission access  [%s] [%s] ", role, url, method))
			c.Abort()
		}
	}
}
