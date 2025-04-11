package logger

import (
	"fmt"
)

const (
	info    = "[INFO]"
	warning = "[WARNING]"
	error   = "[ERROR]"
)

func Println(level, msg string) {
	fmt.Printf("%v %v\n", level, msg)
}

func Printf(level, msg string, args ...any) {
	formatted := fmt.Sprintf(msg, args...)

	fmt.Printf("%v %s\n", level, formatted)
}

func Info(msg string) {
	Println(info, "Hello world")
}

func Infof(msg string, args ...any) {
	Printf(info, msg, args...)
}

func Warn(msg string) {
	Println(warning, "Hello world")
}

func Warnf(msg string, args ...any) {
	Printf(warning, msg, args...)
}

func Error(msg string) {
	Println(error, "Hello world")
}

func Errorf(msg string, args ...any) {
	Printf(error, msg, args...)
}
