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
