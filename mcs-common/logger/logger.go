package logger

import (
	"fmt"
	"mcs-nghiadeptrai/mcs-common/config"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logger.Entry) ([]byte, error) {

	clientIP := entry.Data["clientIP"]
	method := entry.Data["method"]
	path := entry.Data["path"]
	proto := entry.Data["proto"]
	statusCode := fmt.Sprintf("%v", entry.Data["statusCode"])
	userAgent := entry.Data["userAgent"]
	traceid := entry.Data["traceId"]
	level := strings.ToUpper(fmt.Sprintf("%v", entry.Level))

	switch level {
	case "INFO":
		level = green(level)
	case "ERROR":
		level = red(level)
	case "WARN":
		level = yellow(level)
	case "DEBUG":
		level = blue(level)
	}

	log := fmt.Sprintf("%s %s [%s] [%s] [%s] [%s] [%s] [%s] [%s]  - %s\n",
		entry.Time.Format("2006-01-02 15:04:05"), level, clientIP, method, path, proto, traceid, userAgent, statusCode, entry.Message)

	return []byte(log), nil
}

func Init() {
	customFormatter := new(CustomFormatter)

	logger.SetFormatter(customFormatter)
	logLevel := config.Appconfig.GetString("Logging.level")
	setLogLevel(logLevel)
	if config.Appconfig.GetBool("Logging.stdout") {
		logger.New().Out = os.Stdout
	} else {
		file, err := os.OpenFile(config.Appconfig.GetString("Logging.path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.SetOutput(file)
		} else {
			fmt.Println("Failed to log to file, using default stderr", err.Error())
		}
	}

}

func setLogLevel(LogLevel string) {

	switch strings.ToLower(LogLevel) {
	case "debug":
		logger.SetLevel(logger.DebugLevel)
	case "info":
		logger.SetLevel(logger.InfoLevel)
	case "warn":
		logger.SetLevel(logger.WarnLevel)
	case "error":
		logger.SetLevel(logger.ErrorLevel)
	default:
		logger.SetLevel(logger.DebugLevel)

	}
}

// LogInfo ...
func LogInfo(msg string, c *gin.Context) {

	logger.WithFields(logger.Fields{
		"clientIP":   c.ClientIP(),
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"proto":      c.Request.Proto,
		"statusCode": c.Writer.Status(),
		"userAgent":  c.Request.UserAgent(),
		"traceId":    c.Request.Header.Get("x-request-id"),
	}).Info(msg)
}

// LogError ...
func LogError(msg string, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"clientIP":   c.ClientIP(),
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"proto":      c.Request.Proto,
		"statusCode": c.Writer.Status(),
		"userAgent":  c.Request.UserAgent(),
		"traceId":    c.Request.Header.Get("x-request-id"),
	}).Error(msg)
}

// LogWarn ...
func LogWarn(msg string, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"clientIP":   c.ClientIP(),
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"proto":      c.Request.Proto,
		"statusCode": c.Writer.Status(),
		"userAgent":  c.Request.UserAgent(),
		"traceId":    c.Request.Header.Get("x-request-id"),
	}).Warn(msg)
}

// LogDebug ...
func LogDebug(msg string, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"clientIP":   c.ClientIP(),
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"proto":      c.Request.Proto,
		"statusCode": c.Writer.Status(),
		"userAgent":  c.Request.UserAgent(),
		"traceId":    c.Request.Header.Get("x-request-id"),
	}).Debug(msg)
}

func LogInfoNoContext(msg string) {
	logger.WithFields(logger.Fields{
		"Time":       "",
		"clientIP":   "",
		"method":     "",
		"path":       "",
		"proto":      "",
		"statusCode": "",
		"userAgent":  "",
		"traceId":    "",
	}).Infoln(msg)
}

func LogErrorNoContext(msg string) {
	logger.WithFields(logger.Fields{
		"Time":       "",
		"clientIP":   "",
		"method":     "",
		"path":       "",
		"proto":      "",
		"statusCode": "",
		"userAgent":  "",
		"traceId":    "",
	}).Errorln(msg)
}

func LogWarnNoContext(msg string) {
	logger.WithFields(logger.Fields{
		"Time":       "",
		"clientIP":   "",
		"method":     "",
		"path":       "",
		"proto":      "",
		"statusCode": "",
		"userAgent":  "",
		"traceId":    "",
	}).Warnln(msg)
}

func LogDebugNoContext(msg string) {
	logger.WithFields(logger.Fields{
		"Time":       "",
		"clientIP":   "",
		"method":     "",
		"path":       "",
		"proto":      "",
		"statusCode": "",
		"userAgent":  "",
		"traceId":    "",
	}).Debugln(msg)
}

func LogFatalNoContext(msg string) {
	logger.WithFields(logger.Fields{
		"Time":       "",
		"clientIP":   "",
		"method":     "",
		"path":       "",
		"proto":      "",
		"statusCode": "",
		"userAgent":  "",
		"traceId":    "",
	}).Fatalln(msg)
}
