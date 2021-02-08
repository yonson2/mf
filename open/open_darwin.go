// +build darwin

package open

import (
	"os/exec"
)

func RunWith(input string, appName string) error {
	return exec.Command("open", "-a", appName, input).Run()
}
