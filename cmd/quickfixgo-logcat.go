package main

import "github.com/reedom/quickfixgo-logcat"
import "os"

func main() {
	if app, err := logcat.NewApp(os.Args[1:]); err == nil {
		app.Run()
	}
}
