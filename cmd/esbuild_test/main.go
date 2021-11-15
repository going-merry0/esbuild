package main

import (
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"app.js"},
		Bundle:      true,
		Platform:    api.PlatformNode,
		Engines: []api.Engine{
			{Name: api.EngineNode, Version: "10.4"},
		},
		Write: true,
	})

	if len(result.Errors) > 0 {
		os.Exit(1)
	}
}
