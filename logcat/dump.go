package logcat

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/reedom/quickfixgo-logcat/logcat/dict"
)

const tagBeginString int = 8

var reTimestamp *regexp.Regexp
var reScreenLog *regexp.Regexp

func init() {
	reTimestamp = regexp.MustCompile(`^([\d/-]{8,12}[ T-][\d:]{6,8}(\.\d+)?(Z|[+-]\S+)?)\s+`)
	reScreenLog = regexp.MustCompile(`(\d+\S+\x01.*\x01)`)
}

type field struct {
	tag   int
	value string
}

type lineParser struct {
	timestamp string
	fields    []*field
	nFields   int
}

func newLineParser() *lineParser {
	return &lineParser{
		fields: make([]*field, 0),
	}
}

func (p *lineParser) parse(line string) bool {
	p.nFields = 0

	if m := reTimestamp.FindStringSubmatch(line); m != nil {
		// timestamp
		p.timestamp = m[1]
		line = line[len(m[0]):]
	} else if m := reScreenLog.FindStringSubmatch(line); m != nil {
		// screen log
		p.timestamp = ""
		line = m[1]
	} else {
		return false
	}

	// fields
	fields := p.fields
	n := 0
	for {
		pos := strings.Index(line, "\x01")
		if pos < 2 {
			break
		}

		field := p.parseField(line[0:pos])
		line = line[pos+1:]
		if field == nil {
			continue
		}

		if len(fields) <= n {
			fields = append(fields, field)
		} else {
			fields[n] = field
		}
		n++
	}

	p.fields = fields
	p.nFields = n
	return 0 < n
}

func (p *lineParser) parseField(s string) *field {
	pos := strings.Index(s, "=")
	if pos < 0 {
		return nil
	}

	tag, err := strconv.Atoi(s[0:pos])
	if err != nil {
		return nil
	}

	value := s[pos+1:]
	return &field{tag, value}
}

func (p *lineParser) getFields() []*field {
	return p.fields[0:p.nFields]
}

func (p *lineParser) findField(tag int) *field {
	for _, f := range p.getFields() {
		if f.tag == tag {
			return f
		}
	}
	return nil
}

type logPrinter interface {
	print(parser *lineParser)
}

type simpleLogPrinter struct {
	indent string
}

func (p simpleLogPrinter) print(parser *lineParser) {
	if parser.timestamp != "" {
		fmt.Printf("[%s]\n", parser.timestamp)
	}

	for _, f := range parser.getFields() {
		fmt.Printf("%s%d=%s\n", p.indent, f.tag, f.value)
	}
}

type humanLogPrinter struct {
	indent  string
	fixdict dict.FixDict
}

func (p *humanLogPrinter) print(parser *lineParser) {
	// prepare FixDict
	f := parser.findField(tagBeginString)
	if f == nil {
		return
	}
	if p.fixdict == nil || p.fixdict.Version() != f.value {
		p.fixdict = dict.Get(f.value)
	}

	// print
	if parser.timestamp != "" {
		fmt.Printf("[%s]\n", parser.timestamp)
	}

	for _, f := range parser.getFields() {
		fmt.Printf("%s%s=%s\n", p.indent, p.formatTag(f), p.formatTagValue(f))
	}
}

func (p *humanLogPrinter) formatTag(f *field) string {
	if name, ok := p.fixdict.TagName(f.tag); ok {
		return fmt.Sprintf("%d(%s)", f.tag, name)
	}
	return fmt.Sprintf("%d", f.tag)
}

func (p *humanLogPrinter) formatTagValue(f *field) string {
	if name, ok := p.fixdict.ValueName(f.tag, f.value); ok {
		return fmt.Sprintf("%s(%s)", f.value, name)
	}
	return f.value
}

type dumper struct {
	parser  *lineParser
	printer logPrinter
}

func (app *App) newDumper() dumper {
	d := dumper{
		parser: newLineParser(),
	}

	var indent string
	if app.opts.Indent {
		indent = "  "
	}

	if app.opts.Human {
		d.printer = &humanLogPrinter{
			indent: indent,
		}
	} else {
		d.printer = simpleLogPrinter{
			indent: indent,
		}
	}

	return d
}

func (d dumper) dump(reader io.Reader) {
	r := bufio.NewReader(reader)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "Read error: %s\n", err.Error())
			}
			return
		}

		if d.parser.parse(line) {
			d.printer.print(d.parser)
		} else {
			fmt.Printf(line)
		}
	}
}
