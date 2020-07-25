package macaw

import (
	"fmt"
	"strings"
	"time"
)

const (
	resetColor   = "\033[0m"
	infoColor    = "\033[0;34m"
	noticeColor  = "\033[0;36m"
	warningColor = "\033[0;33m"
	errorColor   = "\033[0;31m"
	debugColor   = "\033[0;36m"
)

// Logger Macaw instance
type Logger struct {
	Name string
}

// CreateSubLogger Create new sub logger
func (m *Logger) CreateSubLogger(name ...string) Logger {
	return Logger{Name: m.Name + " | " + strings.Join(name[:], " | ")}
}

func (m *Logger) print(logLevel string, color string, params ...interface{}) {
	paramStr := ""
	output := []interface{}{
		color, time.Now().Format("2006-01-02T15:04:05-0700"), logLevel, m.Name,
	}

	for i := 0; i < len(params); i++ {
		paramStr += "%v"
		output = append(output, params[i])
	}

	fmt.Printf("%s%s (%s) [%v] "+paramStr+"\033[0m\n", output...)
}

// Debug Print message with severity level debug
func (m *Logger) Debug(Params ...interface{}) {
	m.print("Debug", debugColor, Params...)
}

// Info Print message with severity level debug
func (m *Logger) Info(Params ...interface{}) {
	m.print("Info", infoColor, Params...)
}

// Warning Print message with severity level debug
func (m *Logger) Warning(Params ...interface{}) {
	m.print("Warning", warningColor, Params...)
}

// Error Print message with severity level error
func (m *Logger) Error(Params ...interface{}) {
	m.print("Error", errorColor, Params...)
}

// CreateLogger Create a new logger instance
func CreateLogger(name string) Logger {
	logger := Logger{Name: name}

	return logger
}
