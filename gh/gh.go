package gh

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
)

// gh shells out to gh, connecting IO handles for user input
func Exec(args ...string) (err error) {
	ghBin, err := safeexec.LookPath("gh")
	if err != nil {
		err = fmt.Errorf("could not find gh. Is it installed? error: %w", err)
		return
	}
	fmt.Printf("gh %s\n", strings.Join(args, " "))
	cmd := exec.Command(ghBin, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
