package api

import (
	"os/exec"
)

func TTY(h string, u string, p string) (status int) {
	commandAndArgs := []string{
		"-w",
		"sshpass",
		"-p",
		p,
		"ssh",
		h,
		"-l",
		u,
	}
	cmd := exec.Command("gotty", commandAndArgs...)
	err := cmd.Start()
	if err != nil {
		cmd.Process.Kill()
		return -1
	} else {
		return 0
	}
}
