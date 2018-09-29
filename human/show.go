package human

import (
	"github.com/liangran2018/lived/base"

	"github.com/gin-gonic/gin"

)

func Show(c *gin.Context) {
	base.Output(c, 0, GetHuman())
	return
}
