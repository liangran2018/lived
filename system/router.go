package system

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/frame"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/explore"
	"github.com/liangran2018/lived/plat/home"

	"github.com/gin-gonic/gin"
)

func AddRoute(app *gin.Engine) {
	// 系统相关
	app.GET("", frame.MainWin)
	app.GET("/api/newgame", frame.NewGame)
	app.GET("/api/getbackup", frame.GetBackup)
	app.GET("/api/loadbackup", frame.ChooseBackup)
	app.GET("/api/save", frame.Save)

	app.GET("/api/herodetail", human.Show)
	app.GET("/api/ownthingshow", materiel.Show)
	app.POST("/api/ownthingadd", materiel.Add)
	app.POST("/api/ownthingplus", materiel.Plus)

	app.GET("/api/equipchoose", explore.EquipChoose)
	app.POST("/api/equip", explore.Equip)
	app.GET("/api/equipshow", explore.Show)

	app.POST("/bag", explore.Bag) //bug


	app.GET("/api/buildshow", home.Show)
	app.GET("/api/buildnotice", home.Notice)
	app.GET("/api/buildupdate", home.Update)
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
