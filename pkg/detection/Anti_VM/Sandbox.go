package Anti_VM

import (
	"os"
	"strings"
)

func WindowsSandbox() bool {
	username := os.Getenv("USERNAME")
	return strings.ToLower(username) == "wdagutilityaccount"
}
