package log

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"io/ioutil"

	"github.com/liangran2018/lived/base"
)

type logging struct {
	l *os.File
}

type LogLvl int

var logFile *logging

const (
	Info LogLvl = iota
	Warning
	Wrong
)

func init() {
	logFile = &logging{}
}

func NewLogFile() {
	if ok := base.Exists("./operaLog"); !ok {
		if err := os.Mkdir("./operaLog", os.ModePerm); err != nil {
			panic(err)
		}
	}
	files, err := ioutil.ReadDir("./operaLog")
	if err != nil {
		panic(err)
	}

	i := 1
	for _, file := range files {
		if file.IsDir() || base.Exists(fmt.Sprintf("./operaLog/log%d.txt", i)) {
			i++
			continue
		}

		break
	}

	file, err := os.Create(fmt.Sprintf("./operaLog/log%d.txt", i))
	if err != nil {
		panic(err)
	}

	logFile.l = file
}

func OpenLogFile(i int) {
	file, err := os.Open(fmt.Sprintf("./operaLog/log%s.txt", strconv.Itoa(i)))
	if err != nil {
		panic(err)
	}

	logFile.l = file
}

func GetLogger() *logging {
	return logFile
}

func (this LogLvl) notice() string {
	switch this {
	case Info:
		return "[info] "
	case Warning:
		return "[Warn] "
	case Wrong:
		return "[ERR] "
	default:
	}

	return ""
}

func (this *logging) Log(lvl LogLvl, opera ...interface{}) {
	this.l.WriteString(lvl.notice())
	this.l.WriteString(time.Now().Format("2006-01-02 15:04:05") + " ")
	if len(opera) == 0 {
		return
	}

	for _, o := range opera {
		this.l.WriteString(base.StrVal(o) + " ")
	}
	this.l.WriteString("\n")
}

func (this *logging) Close() {
	this.l.Close()
}
