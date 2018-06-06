package vslogger

import (
        "log"
        "fmt"
        "io"
)

const (
        Ldate         = log.Ldate
        Ltime         = log.Ltime
        Lmicroseconds = log.Lmicroseconds
        Llongfile     = log.Llongfile
        Lshortfile    = log.Lshortfile
        LUTC          = log.LUTC
        LstdFlags     = log.LstdFlags
)

type Logger struct {
        log *log.Logger
        level LogLevel
}

type LogLevel uint

const (
        LogError   LogLevel = 1
        LogWarning LogLevel = 2
        LogInfo    LogLevel = 3
        LogDebug   LogLevel = 4
)

func New(level LogLevel, out io.Writer, prefix string, flag int) *Logger {
        return &Logger{
                log: log.New(out, prefix, flag),
                level: level }
}

func (l *Logger) Error(v ...interface{}) {
        if l.level >= LogError {
                v = append([]interface{}{"[E]"}, v...)
                l.log.Output(2, fmt.Sprintln(v...))
        }
}

func (l *Logger) Warning(v ...interface{}) {
        if l.level >= LogWarning {
                v = append([]interface{}{"[W]"}, v...)
                l.log.Output(2, fmt.Sprintln(v...))
        }
}

func (l *Logger) Info(v ...interface{}) {
        if l.level >= LogInfo {
                v = append([]interface{}{"[I]"}, v...)
                l.log.Output(2, fmt.Sprintln(v...))
        }
}

func (l *Logger) Debug(v ...interface{}) {
        if l.level >= LogDebug {
                v = append([]interface{}{"[D]"}, v...)
                l.log.Output(2, fmt.Sprintln(v...))
        }
}
