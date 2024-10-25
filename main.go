package main

import (
	"fmt"
	Anti_Debug "goware/pkg/detection/Anti_Debug"
	"goware/pkg/detection/Anti_VM"
	"goware/pkg/modules"
	"log"
	"os"
	"path/filepath"
)

func CheckForKVM() (bool, error) {
	badDriversList := []string{"balloon.sys", "netkvm.sys", "vioinput", "viofs.sys", "vioser.sys"}
	for _, driver := range badDriversList {
		files, err := filepath.Glob(filepath.Join(os.Getenv("SystemRoot"), "System32", driver))
		if err != nil {
			log.Printf("Error accessing system files for %s: %v", driver, err)
			continue
		}
		if len(files) > 0 {
			return true, nil
		}
	}
	return false, nil
}

func main() {
	fmt.Println()
	fmt.Println("=== Anti-VM Detection ===")
	fmt.Println(fmt.Sprintf("Windows Sandbox: %v", Anti_VM.WindowsSandbox()))
	fmt.Println(fmt.Sprintf("VMWare: %v", Anti_VM.VMWare()))
	fmt.Println(fmt.Sprintf("VirtualBox: %v", Anti_VM.VirtualBox()))
	fmt.Println(fmt.Sprintf("Triage: %v", Anti_VM.Triage()))
	fmt.Println()
	fmt.Println("=== Anti-Debug Detection ===")
	fmt.Println(fmt.Sprintf("IsDebuggerPresent: %v", Anti_Debug.IsDebuggerPresent()))
	fmt.Println(fmt.Sprintf("DetectProcess: %v", Anti_Debug.DetectProcess()))
	fmt.Println()
	fmt.Println("=== Modules ===")
	fmt.Println(fmt.Sprintf("IsAdmin: %v", modules.IsAdmin()))
}
