package Anti_VM

import (
	"github.com/yusufpapurcu/wmi"
	"strings"
)

// BIOS; VirtualBox

type BaseBoard struct {
	Product      string
	Manufacturer string
	SerialNumber string
	Version      string
}

func checkMotherboard() bool {
	var result []BaseBoard
	_ = wmi.Query("SELECT Product, Manufacturer, SerialNumber, Version FROM Win32_BaseBoard", &result)

	for _, board := range result {
		if strings.Contains(strings.ToLower(board.Product), "virtualbox") {
			return true
		}
	}
	return false
}

func checkDisc() bool {
	var disks []DiskInfo
	query := "SELECT DeviceID, PNPDeviceID, Description, MediaType FROM Win32_DiskDrive"
	_ = wmi.Query(query, &disks)

	for _, disk := range disks {
		if strings.Contains(strings.ToLower(disk.PNPDeviceID), "ven_vbox") {
			return true
		}
	}
	return false
}

func checkBios() bool {
	var bios []BIOSInfo
	query := "SELECT SerialNumber FROM Win32_BIOS"
	_ = wmi.Query(query, &bios)
	return strings.Contains(strings.ToLower(bios[0].SerialNumber), "virtualbox")
}

func VirtualBoxMotherboard() bool {
	return checkMotherboard()
}

func VirtualBoxDisc() bool {
	return checkDisc()
}

func VirtualBoxBIOS() bool {
	return checkBios()
}
