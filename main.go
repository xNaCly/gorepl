package main

import (
	"log/slog"

	"github.com/xnacly/go-repl/repl"
)

func main() {
	if err := HasGo(); err != nil {
		slog.Error("Go executable not found, required!", "err", err)
	}

	r := repl.Repl{}
}
