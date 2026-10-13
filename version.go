package main

import (
	_ "embed"
	"fmt"
	"io"
	"runtime/debug"
	"strings"
)

//go:embed version.txt
var serverVersion string

type BuildInfo struct {
	Version string
	Commit  string
	Dirty   bool
}

func getBuildInfo() BuildInfo {
	build := BuildInfo{
		Version: strings.TrimSpace(serverVersion),
		Commit:  "unknown",
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return build
	}

	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			build.Commit = setting.Value
		case "vcs.modified":
			build.Dirty = setting.Value == "true"
		}
	}

	return build
}

func printVersion(w io.Writer) {
	build := getBuildInfo()

	commit := build.Commit
	if len(commit) > 12 {
		commit = commit[:12]
	}

	if build.Dirty {
		commit += "-dirty"
	}

	fmt.Fprintf(w, "%s version %s (commit: %s)\n", serverName, build.Version, commit)
}
