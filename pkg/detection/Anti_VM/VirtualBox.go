package Anti_VM

import (
	"os/exec"
	"strings"
	"syscall"
)

func VirtualBox() bool {
	cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	gpu, _ := cmd.Output()
	detected := strings.Contains(strings.ToLower(string(gpu)), "virtualbox")
	if detected {
		return true
	}
	return false
}
