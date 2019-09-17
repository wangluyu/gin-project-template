package util

import (
	"gin-project-template/pkg/err"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func (g *Gin) Success(errCode int, data interface{}) {
	g.Response(http.StatusOK, errCode, data)
	return
}

func (g *Gin) Error(errCode int, data interface{}) {
	g.Response(http.StatusBadRequest, errCode, data)
	return
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, response{
		Code: errCode,
		Msg:  err.GetErrMsg(errCode),
		Data: data,
	})
	return
}
