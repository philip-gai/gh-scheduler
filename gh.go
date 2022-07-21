package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/cli/safeexec"
)

// gh shells out to gh, returning STDOUT/STDERR and any error
func gh(args ...string) (sout, eout bytes.Buffer, err error) {
	ghBin, err := safeexec.LookPath("gh")
	if err != nil {
		err = fmt.Errorf("could not find gh. Is it installed? error: %w", err)
		return
	}
	cmd := exec.Command(ghBin, args...)
	cmd.Stderr = &eout
	cmd.Stdout = &sout
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to run gh. error: %w, stderr: %s", err, eout.String())
		return
	}
	return
}

// gh shells out to gh, connecting IO handles for user input
func ghWithInput(args ...string) error {
	ghBin, err := safeexec.LookPath("gh")
	if err != nil {
		return fmt.Errorf("could not find gh. Is it installed? error: %w", err)
	}
	cmd := exec.Command(ghBin, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run gh. error: %w", err)
	}
	return nil
}
