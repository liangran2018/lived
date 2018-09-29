package system

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/frame"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/materiel"

	"github.com/gin-gonic/gin"
)

func AddRoute(app *gin.Engine) {
	// 系统相关
	app.GET("", frame.MainWin)
	app.GET("/newgame", frame.NewGame)
	app.GET("/getbackup", frame.GetBackup)
	app.GET("/loadbackup", frame.ChooseBackup)
	app.GET("/save", frame.Save)

	app.GET("/herodetail", human.Show)
	app.GET("/ownthingshow", materiel.Show)
	app.POST("/ownthingadd", materiel.Add)
	app.POST("/ownthingplus", materiel.Plus)
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
