package logcat

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

const version = "0.1.0"

type App struct {
	opts opts
	tags map[string]string
}

func NewApp(args []string) (*App, error) {
	app := App{
		tags: make(map[string]string),
	}
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
		fmt.Fprintf(os.Stdout, "%v version %s\n", os.Args[0], version)
		os.Exit(1)
	}

	if app.opts.Human {
		app.populateTags()
	}

	for _, path := range app.opts.Args.Paths {
		if path == "-" {
			app.dump(os.Stdin)
			continue
		}

		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open file: %s\n", err.Error())
			continue
		}
		defer file.Close()

		app.dump(file)
	}
}
