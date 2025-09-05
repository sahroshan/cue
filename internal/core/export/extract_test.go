// Copyright 2020 CUE Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package export_test

import (
	"fmt"
	"testing"

	"github.com/sahroshan/cue/internal/core/adt"
	"github.com/sahroshan/cue/internal/core/compile"
	"github.com/sahroshan/cue/internal/core/eval"
	"github.com/sahroshan/cue/internal/core/export"
	"github.com/sahroshan/cue/internal/cuetdtest"
	"github.com/sahroshan/cue/internal/cuetxtar"
)

func TestExtract(t *testing.T) {
	test := cuetxtar.TxTarTest{
		Root: "./testdata/main",
		Name: "doc",

		// TODO: use FullMatrix when enough tests pass.
		Matrix: cuetdtest.SmallMatrix,
	}

	test.Run(t, func(t *cuetxtar.Test) {
		r := t.Runtime()
		a := t.Instance()

		v, err := compile.Files(nil, r, "", a.Files...)
		if err != nil {
			t.Fatal(err)
		}

		ctx := eval.NewContext(r, v)
		v.Finalize(ctx)

		writeDocs(t, r, v, nil)
	})
}

func writeDocs(t *cuetxtar.Test, r adt.Runtime, v *adt.Vertex, path []string) {
	fmt.Fprintln(t, path)
	for _, c := range export.ExtractDoc(v) {
		fmt.Fprintln(t, "-", c.Text())
	}

	// Simulate the dereference behavior that is implemented in the CUE api.
	v = v.DerefValue()
	for _, a := range v.Arcs {
		writeDocs(t, r, a, append(path, a.Label.SelectorString(r)))
	}
}
