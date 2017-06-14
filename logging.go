package logging

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

type Position int

const (
	POSITION_SINGLE     Position = 1
	POSITION_IN_MANAGER Position = 2
)

func init() {
	log.SetFlags(log.LstdFlags)
}

type Logger interface {
	GetPosition() Position
	SetPosition(pos Position)
	Info(v ...interface{}) string
	Infof(format string, v ...interface{}) string
	Infoln(v ...interface{}) string
	Error(v ...interface{}) string
	Errorf(format string, v ...interface{}) string
	Errorln(v ...interface{}) string
	Panic(v ...interface{}) string
	Panicf(format string, v ...interface{}) string
	Panicln(v ...interface{}) string
	Warn(v ...interface{}) string
	Warnf(format string, v ...interface{}) string
	Warnln(v ...interface{}) string
	Fatal(v ...interface{}) string
	Fatalf(format string, v ...interface{}) string
	Fatalln(v ...interface{}) string
}

func getInvokerLocation(skipNum int) string {
	pc, file, line, ok := runtime.Caller(skipNum)
	if !ok {
		return ""
	}

	simpleFileName := ""
	if index := strings.LastIndex(file, "/"); index > 0 {
		simpleFileName = file[index+1:]
	}

	funcPath := ""
	funcPc := runtime.FuncForPC(pc)
	if funcPc != nil {
		funcPath = funcPc.Name()
	}

	return fmt.Sprintf("%s : (%s:%d)", funcPath, simpleFileName, line)
}

func generateLogContent(tag LoggerTag, pos Position, format string, v ...interface{}) string {
	skipNum := int(pos) + 2
	baseInfo := fmt.Sprintf("%s %s -", tag.TagPrefix(), getInvokerLocation(skipNum))

	var result string

	if len(format) > 0 {
		result = fmt.Sprintf(fmt.Sprint(baseInfo, format), v)
	} else {
		vLen := len(v)
		params := make([]interface{}, vLen+1)
		params[0] = baseInfo
		for i := 1; i <= vLen; i++ {
			params[i] = v[i-1]
		}
		result = fmt.Sprint(params)
	}

	return result
}

func NewSimpleLogger() Logger {
	logger := &ConsoleLogger{}
	logger.SetPosition(POSITION_SINGLE)
	return logger
}

func NewLogger(loggers []Logger) Logger {
	for _, logger := range loggers {
		logger.SetPosition(POSITION_IN_MANAGER)
	}

	return &LogManager{loggers: loggers}
}
