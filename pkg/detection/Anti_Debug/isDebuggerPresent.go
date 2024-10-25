package Anti_Debug

import "syscall"

var (
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")
	isDebugger  = kernel32DLL.NewProc("IsDebuggerPresent")
)

func IsDebuggerPresent() bool {
	flag, _, _ := isDebugger.Call()
	return flag != 0
}
