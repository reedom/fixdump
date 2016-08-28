package logcat

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"path"
)

const version = "0.1.0"

type App struct {
	opts opts
}

func NewApp(args []string) (*App, error) {
	app := App{}
	_, err := flags.ParseArgs(&app.opts, args)
	if err != nil {
		return nil, err
	}

	if len(app.opts.Args.Paths) == 0 {
		app.opts.Args.Paths = []string{"-"}
	}

	return &app, nil
}

func (app *App) Run() {
	if app.opts.Version {
		cmdname := path.Base(os.Args[0])
		fmt.Fprintf(os.Stdout, "%v version %s\n", cmdname, version)
		os.Exit(1)
	}

	dumper := app.newDumper()

	for _, path := range app.opts.Args.Paths {
		if path == "-" {
			dumper.dump(os.Stdin)
			continue
		}

		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open file: %s\n", err.Error())
			continue
		}

		dumper.dump(file)
		file.Close()
	}
}
