package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"staging/pkg/response"
)

func test(c *gin.Context) {
	// todo 接收参数
	var input struct {
		Test string `form:"test"`
	}
	if err := c.Bind(&input); err != nil {
		response.Fail(c, err)
	}
	// todo 参数校验

	data, err := svc.Test()
	if err != nil {
		zap.L().Error("test failed,err: %v", zap.Error(err))
		response.Fail(c, err)
	}
	response.Success(c, data)
}
