package main

import (
	"github.com/reedom/fixdump/logcat"
	"os"
)

func main() {
	if app, err := logcat.NewApp(os.Args[1:]); err == nil {
		app.Run()
	}
}
