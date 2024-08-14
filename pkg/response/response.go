package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code int   `json:"code"`
	Msg  error `json:"msg"`
	Data any   `json:"data"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, response{
		Code: 0,
		Msg:  nil,
		Data: data,
	})
	c.Abort()
}

func Fail(c *gin.Context, msg error) {
	c.JSON(http.StatusOK, response{
		Code: 1,
		Msg:  msg,
		Data: nil,
	})
	c.Abort()
}
