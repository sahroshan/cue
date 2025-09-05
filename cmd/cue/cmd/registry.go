package cmd

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/sahroshan/cue/internal/cuedebug"
	"github.com/sahroshan/cue/internal/httplog"
	"github.com/sahroshan/cue/internal/mod/modload"
	"github.com/sahroshan/cue/mod/modconfig"
)

// getRegistryResolver returns an implementation of [modregistry.Resolver]
// that resolves to registries as specified in the configuration.
func getRegistryResolver() (*modconfig.Resolver, error) {
	return modconfig.NewResolver(newModConfig(""))
}

func getCachedRegistry() (modload.Registry, error) {
	return modconfig.NewRegistry(newModConfig(""))
}

func newModConfig(registry string) *modconfig.Config {
	return &modconfig.Config{
		Transport:   httpTransport(),
		ClientType:  "cmd/cue",
		CUERegistry: registry,
	}
}

func httpTransport() http.RoundTripper {
	cuedebug.Init()
	if !cuedebug.Flags.HTTP {
		return http.DefaultTransport
	}
	return httplog.Transport(&httplog.TransportConfig{
		// It would be nice to use the default slog logger,
		// but that does a terrible job of printing structured
		// values, so use JSON output instead.
		Logger: httplog.SlogLogger{
			Logger: slog.New(slog.NewJSONHandler(os.Stderr, nil)),
		},
	})
}
