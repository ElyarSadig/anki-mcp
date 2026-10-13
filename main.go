package main

import (
	"anki-mcp/anki"
	"context"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	showVersion := flag.Bool("version", false, "Show build information")
	checkAnki := flag.Bool("check-anki", false, "check the Anki connection and print its version")
	flag.Parse()

	if *showVersion || (flag.NArg() > 0 && flag.Arg(0) == "version") {
		printVersion(os.Stdout)
		return
	}

	anki, err := anki.New(os.Getenv("ANKI_CONNECT_URL"), os.Getenv("ANKI_API_KEY"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid Anki Connect configuration.")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *checkAnki {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		description, err := anki.Describe(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to connect to Anki. Make sure Anki is running with AnkiConnect enabled.")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := printAnkiDescription(description); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to print Anki information.")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func printAnkiDescription(description anki.Description) error {
	_, err := fmt.Fprintf(os.Stdout, "Anki version: %s\nAPI major: %d\n", description.AnkiVersion, description.ApiMajor)
	return err
}
