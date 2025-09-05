package cue_test

import (
	"testing"

	"github.com/go-quicktest/qt"
	"github.com/sahroshan/cue/cue/cuecontext"
)

func TestEmbedFailsWhenNotInModule(t *testing.T) {
	ctx := cuecontext.New()
	v := ctx.CompileString(`
@extern(embed)

package foo
x: _ 	@embed(file="testdata/readme.md",type=text)
`)
	qt.Assert(t, qt.ErrorMatches(v.Err(), `cannot embed files when not in a module`))
}
