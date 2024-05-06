package main

import (
	"log/slog"
	"os/exec"
)

// Discarder discards a and returns err
func Discarder(a any, err error) error {
	return err
}

func HasGo() error {
	slog.Debug("Checking for go compiler executable in path")
	return Discarder(exec.LookPath("go"))
}
