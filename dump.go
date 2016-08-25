package logcat

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func (app *App) dump(r io.Reader) {
	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Read error: %s\n", err.Error())
			}
			return
		}

		// chomp
		line = line[0 : len(line)-2]

		// timestamp
		if pos := strings.Index(line, " "); pos == 0 {
			continue
		} else {
			fmt.Printf("%s\n", line[0:pos])
			line = line[pos+1:]
		}

		fields := strings.Split(line, "\x01")
		for _, field := range fields {
			pair := strings.Split(field, "=")
			if len(pair) != 2 {
				continue
			}

			if !app.opts.Human {
				fmt.Printf("  %s=%s\n", pair[0], pair[1])
				continue
			}

			meaning, ok := app.tags[pair[0]]
			if ok {
				fmt.Printf("  %s(%s)=%s\n", pair[0], meaning, pair[1])
			} else {
				fmt.Printf("  %s=%s\n", pair[0], pair[1])
			}
		}
	}
}
