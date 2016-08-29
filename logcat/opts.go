package logcat

type opts struct {
	Version      bool `short:"v" long:"version"`
	Human        bool `short:"H" long:"human" description:"Additionally print tag/value names"`
	Indent       bool `short:"i" long:"indent" description:"Indent tag entries"`
	NoHeartBeats bool `long:"no-heartbeats" description:"Omit heartbeat log from output"`
	Args         struct {
		Paths []string `description:"Log file path(s) to read"`
	} `positional-args:"yes"`
}
