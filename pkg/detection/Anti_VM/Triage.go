package Anti_VM

import (
	"os/exec"
	"strings"
	"syscall"
)

func Triage() bool {
	cmd := exec.Command("wmic", "diskdrive", "get", "model")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmdOutput, err := cmd.Output()
	if err != nil {
		return false
	}

	models := []string{"DADY HARDDISK", "QEMU HARDDISK"}
	for _, model := range models {
		if strings.Contains(string(cmdOutput), model) {
			return true
		}
	}
	return false
}
