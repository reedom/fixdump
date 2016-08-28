package logcat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineParser(t *testing.T) {
	p := newLineParser()

	var tests = []struct {
		timestamp string
		expected  bool
	}{
		{"2014/05/15 19:49:56.659251", true}, // quickfixgo log format
		{"20140515-19:49:56.659", true},      // found in quickfixgo's test code
		{"20140515 194956", true},            // shortest form
		{"2014-05-15 19:49:56", true},        // a basic format
		{"2014-05-15T19:49:56+09:00", true},  // RFC8601

		{" 2014/05/15 19:49:56.659", false},  // SPC at the head
		{"2007-05-15X19:49:56+09:00", false}, // X instead of T
		{"2014/05/15", false},                // date only
		{"19:49:56.659", false},              // time only
		{"2014/05/15 19:49:56,65", false},    // comma instead of period
	}

	for _, test := range tests {
		line := test.timestamp + " 8=FIX.4.2"
		ok := p.parse(line)
		if assert.Equal(t, test.expected, ok, "mismatch result for "+test.timestamp) && ok {
			assert.Equal(t, test.timestamp, p.timestamp)
			if assert.Equal(t, 1, p.nFields) {
				assert.Equal(t, 8, p.fields[0].tag)
			}
		}
	}
}
