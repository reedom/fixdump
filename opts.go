package logcat

type opts struct {
	Version bool `short:"v" long:"version"`
	Human   bool `short:"H" long:"human" description:"Additionally print tag/value names"`
	Indent  bool `short:"i" long:"indent" description:"Indent tag entries"`
	Args    struct {
		Paths []string
	} `positional-args:"yes" required:"yes"`
}
