package base

// 10000~20000系统错误
// 20000~30000操作错误
const (
	OpenFileErr    = 10000
	ReadFileErr    = 10001
	PrintHtmlErr   = 10002
	TypeConvertErr = 10003
	DelFileErr     = 10004
	SaveFileErr    = 10005

	NoBackupFile = 20000
	NoOwnThing   = 20001

	ParaInvalid = 40000
)
