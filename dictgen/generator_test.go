package dictgen_test

import (
	"testing"

	"github.com/reedom/quickfixgo-logcat/dictgen"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"text/template"
)

func TestGenerator(t *testing.T) {
	xmlpaths, err := filepath.Glob("xml/*.xml")
	require.NoError(t, err)

	tmpl, err := template.ParseGlob("tmpl/*.txt")
	require.NoError(t, err)

	err = dictgen.Generate(xmlpaths, "../dict", tmpl)
	require.NoError(t, err)
}
