package Anti_VM

import (
	"os/exec"
	"strings"
	"syscall"
)

func VMWare() bool {
	cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	gpu, _ := cmd.Output()
	detected := strings.Contains(strings.ToLower(string(gpu)), "vmware")
	if detected {
		return true
	}
	return false
}
