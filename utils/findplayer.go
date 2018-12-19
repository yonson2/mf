package utils

import (
	"os/exec"
)

func GetPlayer() (string, error) {
	path, err := exec.LookPath("mpv")
	if err != nil {
		path, err = exec.LookPath("mplayer")
		if err != nil {
			path, err = exec.LookPath("vlc")
		}
	}
	return path, err
}
