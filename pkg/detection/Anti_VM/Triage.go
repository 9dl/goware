package Anti_VM

import (
	"os/exec"
	"strings"
)

const (
	MouseID   = "ACPI\\PNP0F13\\4&22F5829E&0"
	MouseName = "PS2/2 Compatible Mouse"

	InputID   = "USB\\VID_0627&PID_0001\\28754-0000:00:04.0-1"
	InputName = "USB Input Device"

	MonitorID   = "DISPLAY\\RHT1234\\4&22F5829E&0"
	MonitorName = "Generic PnP Monitor"

	DiskID   = "232138804165"
	DiskName = "WDC WDS100T2B0A"

	KeyboardID   = "ACPI\\PNP0303\\4&22F5829E&0"
	KeyboardName = "Standard PS/2 Keyboard"

	CPU = "Intel Core Processor (Broadwell)"
)

func checkDiskdrive() (bool, int) {
	var Index int
	cmd := exec.Command("wmic", "diskdrive", "get", "model,serialnumber")
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), DiskName) {
		Index++
		if strings.Contains(string(output), DiskID) {
			Index++
		}
	}

	return Index >= 1, Index
}

func checkKeyboard() (bool, int) {
	var Index int
	cmd := exec.Command("wmic", "path", "Win32_Keyboard", "get", "Description,DeviceID")
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), KeyboardID) {
		Index++
		if strings.Contains(string(output), KeyboardName) {
			Index++
		}
	}

	return Index >= 1, Index
}

func checkMouse() (bool, int) {
	var Index int
	cmd := exec.Command("wmic", "path", "Win32_PointingDevice", "get", "Description,PNPDeviceID")
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), MouseID) {
		Index++
		if strings.Contains(string(output), MouseName) {
			Index++
		}
	}

	return Index >= 1, Index
}

func checkInput() (bool, int) {
	var Index int
	cmd := exec.Command("wmic", "path", "Win32_PointingDevice", "get", "Description,PNPDeviceID")
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), InputID) {
		Index++
		if strings.Contains(string(output), InputName) {
			Index++
		}
	}

	return Index >= 1, Index
}

func checkMonitor() (bool, int) {
	var Index int
	cmd := exec.Command("wmic", "path", "Win32_DesktopMonitor", "get", "Description,PNPDeviceID")
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), MonitorID) {
		Index++
		if strings.Contains(string(output), MonitorName) {
			Index++
		}
	}

	return Index >= 1, Index
}

func checkCPU() (bool, int) {
	var Index int
	cmd := exec.Command("wmic", "get", "name")
	output, _ := cmd.CombinedOutput()

	if strings.Contains(string(output), CPU) {
		Index++
	}

	return Index >= 1, Index
}

func Triage() bool {
	checks := []func() (bool, int){
		checkDiskdrive,
		checkKeyboard,
		checkMouse,
		checkInput,
		checkMonitor,
		checkCPU,
	}

	totalIndex := 0
	for _, check := range checks {
		_, index := check()
		totalIndex += index
	}

	return totalIndex >= 5
}
