package logcat

type opts struct {
	Version bool   `short:"v" long:"version"`
	TagPath string `short:"t" long:"tagfile" description:"Path to TAG definition file"`
	Human   bool   `short:"H" long:"human" description:"Additionally print tag definition"`
	Args    struct {
		Paths []string
	} `positional-args:"yes" required:"yes"`
}
