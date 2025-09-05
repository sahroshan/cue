package workspace

import (
	"testing"

	"github.com/go-quicktest/qt"
	"github.com/sahroshan/cue/internal/golangorgx/gopls/protocol"
	. "github.com/sahroshan/cue/internal/golangorgx/gopls/test/integration"
)

func TestDelete(t *testing.T) {
	const files = `
-- cue.mod/module.cue --
module: "mod.example/x"
language: version: "v0.11.0"

-- a.cue --
package a
`
	WithOptions(RootURIAsDefaultFolder()).Run(t, files, func(t *testing.T, env *Env) {
		rootURI := env.Sandbox.Workdir.RootURI()

		env.OpenFile("a.cue")
		env.Await(
			env.DoneWithOpen(),
			LogExactf(protocol.Debug, 1, false, "Module dir=%v module=mod.example/x@v0 For file %v/a.cue found [Package dirs=[%v] importPath=mod.example/x@v0:a]", rootURI, rootURI, rootURI),
		)

		err := env.Sandbox.Workdir.RemoveFile(env.Ctx, "a.cue")
		qt.Assert(t, qt.IsNil(err))

		// Closing a.cue will now cause a re-read of the file from disk
		// (as the overlay has gone), which tests that the LSP can cope
		// with the re-read finding the file has really been deleted.
		env.CloseBuffer("a.cue")
		env.Await(
			env.DoneWithClose(),
		)
	})
}
