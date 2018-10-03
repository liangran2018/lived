package system

import (
	"fmt"
	"bytes"
	"strings"
	"runtime"

	"github.com/liangran2018/lived/base"

	"github.com/gin-gonic/gin"
)

// 系统出错后的恢复
func AddRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case base.DEAD:
					base.Output(c, base.HeroDead, e.Reason)
				default:
					base.Output(c, base.ProgramPanic, "系统错误:"+ base.StrVal(err))
					c.Writer.Write(panicFileAndLine(3))
					c.AbortWithStatus(200)
				}
			}
		}()
		c.Next()
	}
}

func panicFileAndLine(skip int) []byte {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "\n")
	//var str string
	for i := skip; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		// vendor：去掉gin框架的文件路径
		if strings.Contains(file, "vendor") || strings.Contains(file, "system") ||
			strings.Contains(file, "panic") {
		//	continue
		}

		fmt.Fprintf(buf, "%s:%d \n", file, line)
	}
	return buf.Bytes()
}

