package main

import (
	"encoding/base64"
	"fmt"
	"goware/pkg/detection/Anti_Debug"
	"goware/pkg/detection/Anti_VM"
	"goware/pkg/modules"
	"io/ioutil"
	"os/exec"
)

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

func main() {
	fmt.Println()
	fmt.Println("=== Anti-VM Detection ===")
	fmt.Println(fmt.Sprintf("Windows Sandbox: %v", Anti_VM.WindowsSandbox()))
	fmt.Println(fmt.Sprintf("VMWare (BIOS Detection): %v", Anti_VM.VMWareBIOS()))
	fmt.Println(fmt.Sprintf("VMWare (Discdrive Detection): %v", Anti_VM.VMWareDisk()))
	fmt.Println(fmt.Sprintf("VMWare (Model Detection): %v", Anti_VM.VMWareDisk()))
	fmt.Println(fmt.Sprintf("VirtualBox (Motherboard Detection): %v", Anti_VM.VirtualBoxMotherboard()))
	fmt.Println(fmt.Sprintf("VirtualBox (Disc Detection): %v", Anti_VM.VirtualBoxDisc()))
	fmt.Println(fmt.Sprintf("VirtualBox (BIOS Detection): %v", Anti_VM.VirtualBoxBIOS()))
	fmt.Println(fmt.Sprintf("Triage: %v", Anti_VM.Triage()))
	fmt.Println(fmt.Sprintf("Triage (Background Detection): %v", Anti_VM.TriageV2()))
	fmt.Println()
	fmt.Println("=== Anti-Debug Detection ===")
	fmt.Println(fmt.Sprintf("IsDebuggerPresent: %v", Anti_Debug.IsDebuggerPresent()))
	fmt.Println(fmt.Sprintf("DetectProcess: %v", Anti_Debug.DetectProcess()))
	fmt.Println()
	fmt.Println("=== Modules ===")
	fmt.Println(fmt.Sprintf("IsAdmin: %v", modules.IsAdmin()))

	fmt.Scanln()
}
