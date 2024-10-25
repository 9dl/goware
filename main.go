package main

import (
	"fmt"
	"goware/pkg/detection/Anti_Debug"
	"goware/pkg/detection/Anti_VM"
	"goware/pkg/modules"
)

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

	fmt.Scanln()
}
