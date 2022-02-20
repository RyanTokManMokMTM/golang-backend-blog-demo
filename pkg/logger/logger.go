package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

//Fields All leve case
type Fields map[string]interface{}

type Level uint8
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

//total 6 level 0-5
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarning:
		return "warning"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return "" //not such case
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	log := log.New(w, prefix, flag)
	return &Logger{
		newLogger: log,
	}
}

func (log *Logger) Clone() *Logger {
	l := *log //dereference
	return &l
}

//WithFields Global field info
func (log *Logger) WithFields(f Fields) *Logger {
	l := log.Clone() //clone itself(not reference to same logger)
	//create the fields
	if l.fields == nil {
		l.fields = make(Fields) //creak a empty map
	}

	//set the fields to logger
	for key, val := range f {
		l.fields[key] = val
	}

	return l
}

//WithContext Content info
func (log *Logger) WithContext(ctx context.Context) *Logger {
	l := log.Clone()
	l.ctx = ctx
	return l
}

//WithCaller specific caller info in stack
//caller info by skip argument
func (log *Logger) WithCaller(skip int) *Logger {
	l := log.Clone()
	//reports file and line number information about function invocations
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc) //return the func info for current pc
		l.callers = []string{fmt.Sprintf("%s:%d %s", file, line, f.Name())}
	}
	return l
}

//WithCallerFrame Whole Caller Stack Info
//All caller info
func (log *Logger) WithCallerFrame() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string

	pcs := make([]uintptr, maxCallerDepth)        //total callers
	depth := runtime.Callers(minCallerDepth, pcs) //current callers in stack?
	frames := runtime.CallersFrames(pcs[:depth])  //current callers frames?
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}

	l := log.Clone()
	l.callers = callers
	return l
}

//JSONFormat JSON Formatting the info with fields info,time,message and caller info
func (log *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(log.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = log.callers
	if len(log.fields) > 0 {
		//storing field info
		for key, value := range log.fields {
			//check data isn't in data
			if _, ok := data[key]; !ok {
				data[key] = value
			}
		}
	}
	return data
}

//Output print info with logger
func (log *Logger) Output(level Level, message string) {
	//encoding formatted info to body
	body, _ := json.Marshal(log.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		log.newLogger.Println(content)
	case LevelWarning:
		log.newLogger.Println(content)
	case LevelError:
		log.newLogger.Println(content)
	case LevelInfo:
		log.newLogger.Println(content)
	case LevelFatal:
		log.newLogger.Println(content)
	case LevelPanic:
		log.newLogger.Println(content)
	}
}

//OUTPUT Method

//Info without format
func (log *Logger) Info(v ...interface{}) {
	log.Output(LevelInfo, fmt.Sprintln(v...))
}

//Infof with format
func (log *Logger) Infof(format string, v ...interface{}) {
	log.Output(LevelInfo, fmt.Sprintf(format, v...))
}

//Debug without format
func (log *Logger) Debug(v ...interface{}) {
	log.Output(LevelDebug, fmt.Sprintln(v...))
}

//Debugf with format
func (log *Logger) Debugf(format string, v ...interface{}) {
	log.Output(LevelDebug, fmt.Sprintf(format, v...))
}

//Error without format
func (log *Logger) Error(v ...interface{}) {
	log.Output(LevelError, fmt.Sprintln(v...))
}

//Errorf with format
func (log *Logger) Errorf(format string, v ...interface{}) {
	log.Output(LevelError, fmt.Sprintf(format, v...))
}

//Warning without format
func (log *Logger) Warning(v ...interface{}) {
	log.Output(LevelWarning, fmt.Sprintln(v...))
}

//Warningf with format
func (log *Logger) Warningf(format string, v ...interface{}) {
	log.Output(LevelWarning, fmt.Sprintf(format, v...))
}

//v without format
func (log *Logger) Fatal(v ...interface{}) {
	log.Output(LevelFatal, fmt.Sprintln(v...))
}

//Fatalf with format
func (log *Logger) Fatalf(format string, v ...interface{}) {
	log.Output(LevelFatal, fmt.Sprintf(format, v...))
}

//Panic without format
func (log *Logger) Panic(v ...interface{}) {
	log.Output(LevelWarning, fmt.Sprintln(v...))
}

//Panicf with format
func (log *Logger) Panicf(format string, v ...interface{}) {
	log.Output(LevelWarning, fmt.Sprintf(format, v...))
}
