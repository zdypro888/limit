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
	rlim.Cur = 50000
	rlim.Max = 50000
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim); err != nil {
		return err
	}
	return nil
}
