package main

import "os/exec"

// Discarder discards a and returns err
func Discarder(a any, err error) error {
	return err
}

func HasGo() error {
	return Discarder(exec.LookPath("go"))
}
