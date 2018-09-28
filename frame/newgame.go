package frame

import (
	"github.com/liangran2018/lived/base"

	"github.com/gin-gonic/gin"
)

const defaultName = "一一"

func NewGame(c *gin.Context) {
	name := c.Query("name")
	if base.Empty(name) {
		name = defaultName
	}

	base.Output(c, 0, nil)
}
