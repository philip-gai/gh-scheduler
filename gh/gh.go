package gh

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/cli/safeexec"
	"github.com/gizak/termui/v3/widgets"
	"github.com/philip-gai/gh-scheduler/utils"
)

// gh shells out to gh, connecting IO handles for user input
func Exec(logs *widgets.List, args ...string) (err error) {
	ghBin, err := safeexec.LookPath("gh")
	if err != nil {
		err = fmt.Errorf("could not find gh. Is it installed? error: %w", err)
		return
	}
	humanReadableArgs := fmt.Sprintf("gh %s", strings.Join(args, " "))
	utils.PushListRow(humanReadableArgs, logs)
	output, err := exec.Command(ghBin, args...).CombinedOutput()

	outputString := string(output)
	utils.PushListRow(outputString, logs)

	if err != nil {
		utils.PushListRow("Err is not nil", logs)
		utils.PushListRow(fmt.Sprint("Error: ", err.Error()), logs)
		return err
	}
	return nil
}
