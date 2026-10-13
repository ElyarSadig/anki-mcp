package main

import (
	_ "embed"
	"flag"
	"os"
)

func main() {
	showVersion := flag.Bool("version", false, "Show build information")
	flag.Parse()

	if *showVersion || (flag.NArg() > 0 && flag.Arg(0) == "version") {
		printVersion(os.Stdout)
		return
	}

	//TODO Start your MCP server here.
}
