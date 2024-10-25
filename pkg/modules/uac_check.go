package modules

import (
	"syscall"
)

func IsAdmin() bool {
	isAdmin, _, _ := syscall.NewLazyDLL("shell32.dll").NewProc("IsUserAnAdmin").Call()
	return isAdmin != 0
}
