package main

import (
	"fmt"

	"github.com/evanw/esbuild/internal/config"
	"github.com/evanw/esbuild/internal/js_parser"
	"github.com/evanw/esbuild/internal/logger"
	"github.com/evanw/esbuild/internal/test"
)

func main() {
	log := logger.NewDeferLog(logger.DeferLogNoVerboseOrDebug)
	ast, ok := js_parser.Parse(log, test.SourceForTest(`
	// comment
	let a = false;
	let x = [];
	x = x.concat(x);
	`), js_parser.OptionsFromConfig(&config.Options{}))
	fmt.Printf("test %v %v\n", ast, ok)
}
