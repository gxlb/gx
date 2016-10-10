package log

type LogType uint

const (
	LogTrace = iota
	LogConsole
	LogDebug
	LogLog
	LogWarn
	LogErr
	LogPanic
	LogUserBegin
)

type LogOption uint

const (
	LogORemote LogOption = 1 << iota
	LogODb
	LogOCompress
)

type Log struct {
}
