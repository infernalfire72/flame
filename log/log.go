package log

import (
	"fmt"
	"runtime"
	"sync"
)

type Color int

const Reset Color = 0

const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Reset iota
const (
	FgHiBlack Color = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

const EscapeCharacter = "\x1b"

var mux sync.Mutex

func PrintColor(c Color, values ...interface{}) {
	fmt.Printf("\x1b[%dm", c)
	for _, v := range values {
		fmt.Print(v, "")
	}
	fmt.Print("\x1b[0m")
}

func PrintlnColor(c Color, values ...interface{}) {
	fmt.Printf("\x1b[%dm", c)
	for _, v := range values {
		fmt.Print(v, "")
	}
	fmt.Println("\x1b[0m")
}

func Chat(sender, target, message string) {
	mux.Lock()
	PrintColor(FgHiBlue, "CHAT    | ")
	fmt.Printf("%s @ %s: %s", sender, target, message)
	mux.Unlock()
}

func Warn(v ...interface{}) {
	mux.Lock()
	PrintColor(FgHiYellow, "WARNING | ")
	fmt.Println(v...)
	mux.Unlock()
}

func Info(v ...interface{}) {
	mux.Lock()
	PrintColor(FgHiGreen, "INFO    | ")
	fmt.Println(v...)
	mux.Unlock()
}

func Infof(format string, v ...interface{}) {
	mux.Lock()
	PrintColor(FgHiGreen, "INFO    | ")
	fmt.Printf(format+"\n", v...)
	mux.Unlock()
}

func Error(err error) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		mux.Lock()
		PrintColor(FgRed, "ERROR   | ")
		fmt.Println("in", file, "at line", no)
		fmt.Println("      └› ", err)
		mux.Unlock()
	}
}
