package memo28Gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type CatchGlobalErrorOptions struct {
	Callback func(errMsg string, c *gin.Context)
}

func CatchGlobalError(catchGlobalErrorOptions CatchGlobalErrorOptions) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				// 发生全局错误时的处理逻辑
				catchGlobalErrorOptions.Callback(fmt.Sprint(err), c)
			}
		}()

		c.Next()
	}
}
