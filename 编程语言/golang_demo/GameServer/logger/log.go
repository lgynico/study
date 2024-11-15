package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	write       io.Writer
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func Config(filename string) {
	path := "logs/"
	if ex, err := os.Executable(); err == nil {
		path = filepath.Dir(ex) + "/logs/"
	}

	write = &timeRollingFileWriter{
		filename: path + filename,
	}
	log.Printf("Init logger: %s", path+filename)
	infoLogger = log.New(write, "[INFO]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	errorLogger = log.New(write, "[ERROR]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

func Info(format string, args ...any) {
	_ = infoLogger.Output(2, fmt.Sprintf(format, args...))
}

func Error(format string, args ...any) {
	_ = errorLogger.Output(2, fmt.Sprintf(format, args...))
}

func Debug(format string, args ...any) {
}

func Warn(format string, args ...any) {
}
