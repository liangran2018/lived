package frame

import (
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

func MainWin(c *gin.Context) {
	/*
	f, err := os.Open("./pages/index.html")
	if err != nil {
		base.Output(c, base.OpenFileErr, err.Error())
	}

	b := make([]byte, 1024)
	n, err := f.Read(b)
	if err != nil {
		base.Output(c, base.ReadFileErr, err.Error())
	}

	_, err = fmt.Fprint(c.Writer, string(b[:n]))
	if err != nil {
		base.Output(c, base.PrintHtmlErr, err.Error())
	}
*/
	//新建日志文件
	log.NewLogFile()
	//记录
	log.GetLogger().Log(log.Info, "open program")
}