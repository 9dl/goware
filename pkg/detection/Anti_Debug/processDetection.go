package Anti_Debug

import (
	"github.com/shirou/gopsutil/process"
	"strings"
)

var Blacklisted = []string{
	"processh",
	"debug",
	"debugger",
	"hacker",
	"inject",
	"dump",
	"dumper",
	"deobfs",
}

// Make this in a Loop so its makes more sense
func DetectProcess() bool {
	processList, _ := process.Processes()

	for _, p := range processList {
		name, _ := p.Name()
		for _, blacklisted := range Blacklisted {
			if strings.Contains(strings.ToLower(name), blacklisted) {
				return true
			}
		}
	}
	return false
}
