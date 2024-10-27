package Anti_VM

import (
	"bytes"
	"github.com/yusufpapurcu/wmi"
	"os/exec"
	"strings"
)

type DiskInfo struct {
	DeviceID    string
	PNPDeviceID string
	Description string
	MediaType   string
}

type BIOSInfo struct {
	SerialNumber string
}

func checkModelname() bool {
	cmds := []string{
		"Get-WmiObject -Class Win32_ComputerSystem | Select-Object -ExpandProperty Model",
		"Get-WmiObject -Class Win32_ComputerSystem | Select-Object -ExpandProperty Manufacturer",
	}
	var outputs []string

	for _, cmdStr := range cmds {
		cmd := exec.Command("powershell", "-Command", cmdStr)
		var out bytes.Buffer
		cmd.Stdout = &out
		_ = cmd.Run()
		outputs = append(outputs, strings.TrimSpace(out.String()))
	}

	return strings.Contains(strings.ToLower(outputs[0]), "vmware") || strings.Contains(strings.ToLower(outputs[1]), "vmware")
}

func checkBIOS() bool {
	var bios []BIOSInfo
	query := "SELECT SerialNumber FROM Win32_BIOS"
	_ = wmi.Query(query, &bios)
	return strings.Contains(strings.ToLower(bios[0].SerialNumber), "vmware")
}

func checkDiscdrive() bool {
	var disks []DiskInfo
	query := "SELECT DeviceID, PNPDeviceID, Description, MediaType FROM Win32_DiskDrive"
	_ = wmi.Query(query, &disks)

	for _, disk := range disks {
		if strings.Contains(strings.ToLower(disk.PNPDeviceID), "vmware") {
			return true
		}
	}
	return false
}

func VMWareBIOS() bool {
	return checkBIOS()
}

func VMWareDisk() bool {
	return checkDiscdrive()
}

func VMWareModel() bool {
	return checkModelname()
}
