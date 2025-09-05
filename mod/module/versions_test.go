package module_test

import (
	"github.com/sahroshan/cue/internal/mod/mvs"
	"github.com/sahroshan/cue/mod/module"
)

var _ mvs.Versions[module.Version] = module.Versions{}
