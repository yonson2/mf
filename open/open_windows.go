// +build windows

package open

import (
	"os/exec"
	"strings"
)

func cleaninput(input string) string {
	r := strings.NewReplacer("&", "^&")
	return r.Replace(input)
}

func RunWith(input string, appName string) error {
	cmd := exec.Command(appName, cleaninput(input))
	return cmd.Run()
}
