/*
 * @Author: @memo28.repo
 * @Date: 2024-07-17 15:42:03
 * @LastEditTime: 2024-07-17 15:54:01
 * @Description:
 * @FilePath: /memo28.gin/response.go
 */
package memo_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	memoMap "github.com/memo28-space-golang/memo28.map"
)

type Response struct {
}

// Suc 处理成功返回值
func (receiver Response) Suc(context *gin.Context, data interface{}) {
	context.JSON(http.StatusOK, receiver.EncapsulatesResponse(http.StatusOK, data, "ok"))
	context.Abort()
}

// Error 处理错误返回
func (receiver Response) Error(context *gin.Context, code int, data interface{}, msg string) {
	context.JSON(http.StatusOK, receiver.EncapsulatesResponse(code, data, msg))
	context.Abort()
}

// EncapsulatesResponse 组装响应数据
func (receiver Response) EncapsulatesResponse(code int, data interface{}, msg string) memoMap.Map[string, any] {
	u := memoMap.Map[string, any]{}
	u.Add("code", code).Add("data", data).Add("msg", msg)
	return u
}
