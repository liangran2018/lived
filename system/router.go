package system

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/frame"

	"github.com/gin-gonic/gin"
)

func AddRoute(app *gin.Engine) {
	// 系统相关
	app.GET("/newgame", frame.NewGame)
	app.GET("/test", test)
	// 无路由或无方法
	app.NoRoute(error)
	app.NoMethod(error)
}

func test(c *gin.Context) {
	base.Output(c, 0, "test")
}

func error(c *gin.Context) {
	base.Output(c, 400, "access denied")
}
