package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/xnacly/gorepl/repl"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug logs")
	cmd := flag.String("c", "", "Run go code from the cli")
	flag.Parse()

	if *debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Starting with Debug log level")
	}

	if err := HasGo(); err != nil {
		slog.Error("Go executable not found, required!", "err", err)
		os.Exit(1)
	}

	if *cmd != "" {
		r := repl.Repl{
			Instructions: []string{*cmd},
		}
		if err := r.Exec(); err != nil {
			slog.Error("Failed to run, clearing previous instructions", "err", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	r := repl.Repl{}
	if err := r.Wait(); err != nil {
		slog.Info("Got interrupt, exiting...")
		os.Exit(0)
	} else {
		slog.Info("Got exit cmd, exiting...")
	}
}
