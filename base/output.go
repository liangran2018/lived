package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type outputStr struct {
	Result int			`json:"result"`
	Data interface{}	`json:"data"`
}

func Output(c *gin.Context, result int, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	o := outputStr{Result:result, Data:data}
	c.JSON(http.StatusOK, o)
	return
}
