package logging

import (
	"fmt"
	"github.com/svenschaper/goproperties"
	"time"
	"test123
)

type Logger struct {
	Class    string
	Loglevel int
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

const (
	ERROR       = 1
	COMMUNICATE = 2
	REPORT      = 3
	INFO        = 4
	WARN        = 5
	DEBUG       = 6
)

var logger Logger
var prop *properties.Propertie

func init() {
	properties.SetPropertiePath("config.yml")
	prop2, _ := properties.LoadProperty()
	prop = prop2
}

func GeneralInitLogger(class string) Logger {
	var logger Logger
	s := prop.GetProperty("log.level." + class)
	if s != "" {
		return logger.initialize(class, s)
	}
	return logger.initialize(class, prop.GetProperty("log.level"))
}

func (l Logger) initialize(class string, level string) Logger {
	var loglevel int
	switch level {
	case "ERROR":
		loglevel = 1
	case "COMMUNICATE":
		loglevel = 2
	case "REPORT":
		loglevel = 3
	case "INFO":
		loglevel = 4
	case "WARN":
		loglevel = 5
	case "DEBUG":
		loglevel = 6
	}
	return Logger{Class: class, Loglevel: loglevel}
}

func (l Logger) Info(s string, args ...interface{}) {

	if l.Loglevel >= INFO {
		fmt.Println(Green + "[" + time.Now().Format(time.RFC850) + "]" + "   " + "   INFO          [" + l.Class + "]   " + fmt.Sprintf(s, args...) + Reset)
	}
}

func (l Logger) Error(s string, args ...interface{}) {
	if l.Loglevel >= ERROR {
		fmt.Println(Red + "[" + time.Now().Format(time.RFC850) + "]" + "   " + "   ERROR         [" + l.Class + "]   " + fmt.Sprintf(s, args...) + Reset)
	}
}

func (l Logger) Warn(s string, args ...interface{}) {
	if l.Loglevel >= WARN {
		fmt.Println(Purple + "[" + time.Now().Format(time.RFC850) + "]" + "   " + "   WARN          [" + l.Class + "]   " + fmt.Sprintf(s, args...) + Reset)
	}
}

func (l Logger) Debug(s string, args ...interface{}) {
	if l.Loglevel >= DEBUG {
		fmt.Println(Blue + "[" + time.Now().Format(time.RFC850) + "]" + "   " + "   DEBUG         [" + l.Class + "]   " + fmt.Sprintf(s, args...) + Reset)
	}
}

func (l Logger) Report(s string, args ...interface{}) {
	if l.Loglevel >= REPORT {
		fmt.Println(Cyan + "[" + time.Now().Format(time.RFC850) + "]" + "   " + "   REPORT        [" + l.Class + "]   " + fmt.Sprintf(s, args...) + Reset)
	}
}

func (l Logger) Communicate(s string, args ...interface{}) {
	if l.Loglevel >= COMMUNICATE {
		fmt.Println(Yellow + "[" + time.Now().Format(time.RFC850) + "]" + "   " + "   COMMUNICATE   [" + l.Class + "]   " + fmt.Sprintf(s, args...) + Reset)
	}
}
