package swagger

import (
	"gin-project-template/app/api/v1"
)

type Success struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"success"`
	Data interface{} `json:"data" example:"[]"`
}

type Failure struct {
	Code int         `json:"code" example:"500"`
	Msg  string      `json:"msg" example:"error"`
	Data interface{} `json:"data" example:"[]"`
}

type GetExample struct {
	Success
	Data v1.User `json:"data"`
}
