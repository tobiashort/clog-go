package clog

import (
	"strings"

	"github.com/tobiashort/cfmt-go"
)

//nofmt:enable
const LevelDebug = 0
const LevelInfo  = 1
const LevelWarn  = 2
const LevelError = 3

var DebugString = func() string { return cfmt.Sprint("#B{DEBUG}") }
var InfoString  = func() string { return cfmt.Sprint("#bB{INFO}") }
var WarnString  = func() string { return cfmt.Sprint("#yB{WARN}") }
var ErrorString = func() string { return cfmt.Sprint("#rB{ERROR}") }
//nofmt:disable

var Level = LevelInfo

func keyValues(args ...any) string {
	sb := strings.Builder{}
	isVal := false
	for _, arg := range args {
		if isVal {
			cfmt.Fprintf(&sb, "%v ", arg)
			isVal = false
		} else {
			cfmt.Cfprintf(&sb, "c", "%s=", arg)
			isVal = true
		}
	}
	return sb.String()
}

func Debug(msg string) {
	if Level != LevelDebug {
		return
	}
	cfmt.Println(DebugString(), msg)
}

func Debugf(format string, args ...any) {
	if Level != LevelDebug {
		return
	}
	cfmt.Println(DebugString(), cfmt.Sprintf(format, args...))
}

func Debugs(msg string, args ...any) {
	if Level != LevelDebug {
		return
	}
	cfmt.Println(DebugString(), msg, keyValues(args...))
}

func Info(msg string) {
	if Level != LevelInfo && Level != LevelDebug {
		return
	}
	cfmt.Println(InfoString(), msg)
}

func Infof(format string, args ...any) {
	if Level != LevelInfo && Level != LevelDebug {
		return
	}
	cfmt.Println(InfoString(), cfmt.Sprintf(format, args...))
}

func Infos(msg string, args ...any) {
	if Level != LevelInfo && Level != LevelDebug {
		return
	}
	cfmt.Println(InfoString(), msg, keyValues(args...))
}

func Warn(msg string) {
	if Level != LevelWarn && Level != LevelInfo && Level != LevelDebug {
		return
	}
	cfmt.Println(WarnString(), msg)
}

func Warnf(format string, args ...any) {
	if Level != LevelWarn && Level != LevelInfo && Level != LevelDebug {
		return
	}
	cfmt.Println(WarnString(), cfmt.Sprintf(format, args...))
}

func Warns(msg string, args ...any) {
	if Level != LevelWarn && Level != LevelInfo && Level != LevelDebug {
		return
	}
	cfmt.Println(WarnString(), keyValues(args...))
}

func Error(msg string) {
	cfmt.Println(ErrorString(), msg)
}

func Errorf(format string, args ...any) {
	cfmt.Println(ErrorString(), cfmt.Sprintf(format, args...))
}

func Errors(msg string, args ...any) {
	cfmt.Println(ErrorString(), keyValues(args...))
}
