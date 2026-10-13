package main

import (
	"anki-mcp/anki"
	"context"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
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
	if _, err := fmt.Fprintf(os.Stdout, "Anki version: %s\nAPI major: %d\n", description.AnkiVersion, description.ApiMajor); err != nil {
		return err
	}

	keys := make([]string, 0, len(description.Limits))
	for key := range description.Limits {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	if len(keys) == 0 {
		return nil
	}

	width := 0
	for _, key := range keys {
		if len(key) > width {
			width = len(key)
		}
	}

	if _, err := fmt.Fprintln(os.Stdout, "\nLimits:"); err != nil {
		return err
	}
	for _, key := range keys {
		if _, err := fmt.Fprintf(os.Stdout, "  %-*s %s\n", width, key, formatLimit(key, description.Limits[key])); err != nil {
			return err
		}
	}

	return nil
}

func formatLimit(key string, value any) string {
	if strings.HasSuffix(key, "Bytes") {
		if bytes, ok := value.(float64); ok {
			return formatBytes(bytes)
		}
	}
	return fmt.Sprintf("%v", value)
}

func formatBytes(bytes float64) string {
	const unit = int64(1024)

	n := int64(bytes)
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}

	div := unit
	exp := 0
	for next := n / unit; next >= unit; next /= unit {
		div *= unit
		exp++
	}

	suffixes := []string{"KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	return fmt.Sprintf("%.2g %s", float64(n)/float64(div), suffixes[exp])
}
