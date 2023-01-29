package logger

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

var (
	//F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "NOTICE", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	NOTICE
	WARNING
	ERROR
	FATAL
)

func getNowTime() string {
	var nowTime = time.Now()
	return nowTime.Format("2006年01月02日 15:04:05")
}

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

func GetLogPath() string {
	var path string
	if IsWindows() {
		path = "logs"
	} else {
		path = filepath.Join(GetAppPath(), "logs")
	}
	return path
}

// Get the absolute path to the running directory
func GetAppPath() string {
	if path, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		return path
	}
	return os.Args[0]
}

var isDebug = false

func ResetLog(debug bool) {
	isDebug = debug
	timeWriter := &TimeWriter{
		Dir:        GetLogPath(),
		Compress:   true,
		ReserveDay: 30,
	}
	logger = log.New(timeWriter, DefaultPrefix, log.LstdFlags)

	logger.Printf("***************%s**********************", getNowTime())
	logger.Printf("***************ipidea代理启动1.2**********************")
}

func GetLog() *log.Logger {
	return logger
}

func Info(format string, v ...interface{}) {
	setPrefix(INFO)
	if isDebug {
		logs.Notice(format, v...)
	}
	logger.Printf(format, v...)
}

func Notice(format string, v ...interface{}) {
	setPrefix(NOTICE)
	if isDebug {
		logs.Notice(format, v...)
	}
	logger.Printf(format, v...)
}

func Warn(format string, v ...interface{}) {
	setPrefix(WARNING)
	logger.Printf(format, v...)
}

func Error(format string, v ...interface{}) {
	setPrefix(ERROR)
	if isDebug {
		logs.Error(format, v...)
	}
	logger.Printf(format, v...)
}

func Debug(format string, v ...interface{}) {
	//setPrefix(DEBUG)
	if isDebug {
		logs.Debug(format, v...)
		logger.Printf(format, v...)
	}
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
