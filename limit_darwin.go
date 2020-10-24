package limit

import (
	"syscall"
)

//Limit 解除IO限制
func Limit() error {
	var rlim syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim); err != nil {
		return err
	}
	rlim.Cur = 0x2048
	rlim.Max = 0x2048
	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
}
