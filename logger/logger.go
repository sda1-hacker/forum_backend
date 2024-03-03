package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"
)

var logDir = "./runtime/log"

func init() {
	// 设置日志的格式为json
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

func Write(msg string, filename string) { // 用户自定义日志
	setOutPutFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args)
}

func Info(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Info(args)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Warn(args)
}

func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Fatal(false)
}

func Error(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Error(args)
}

func Panic(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.PanicLevel, "panic")
	logrus.WithFields(fields).Panic(args)
}

func Trace(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.TraceLevel, "trace")
	logrus.WithFields(fields).Trace(args)
}

// 自定义日志
func setOutPutFile(level logrus.Level, logName string) {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error '%s'", logDir, err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	logFileName := path.Join(logDir, logName+"_"+timeStr+".log")

	var err error
	os.Stderr, err = os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file err", err)
	}
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	return
}

// 访问日志
func LoggerToFile() gin.LoggerConfig {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error '%s'", logDir, err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	logFileName := path.Join(logDir, "success_"+timeStr+".log")

	var err error
	os.Stderr, err = os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file err", err)
	}

	config := gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s\" %s\"\n",
				params.TimeStamp.Format("2006-01-02 13:04:05"), // 访问时间
				params.ClientIP, // 访问的ip
				params.Method,   // 请求方式
				params.Path,     // 请求路径
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, os.Stderr),
	}

	return config
}

// panic日志
func Recover(context *gin.Context) {
	// recover 只会在defer中执行
	// 如果出现了panic，通过defer中的recover就可以捕获这个panic
	// 通过 context.Abort()终止后续的handler继续执行
	defer func() {
		if err := recover(); err != nil {
			if _, errDir := os.Stat(logDir); os.IsNotExist(errDir) {
				errDir = os.MkdirAll(logDir, 0777)
				if errDir != nil {
					panic(fmt.Errorf("create log dir '%s' error '%s'", logDir, err))
				}
			}

			timeStr := time.Now().Format("2006-01-02")
			logFileName := path.Join(logDir, "error_"+timeStr+".log")

			f, errFile := os.OpenFile(logFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if errFile != nil {
				fmt.Println("open log file err", err)
			}
			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			f.WriteString("panic error time:" + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v", err) + "\n")
			f.WriteString("stacktrace from panic:" + string(debug.Stack()) + "\n")
			f.Close()
			context.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("%v", err),
			})
			// 不再调用后续的handler
			context.Abort()
		}
	}()

	// 如果没有出现panic，则不会调用defer中的recover，需要继续调用后续的handler
	context.Next()
}
