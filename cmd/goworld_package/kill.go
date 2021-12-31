package goworld_package

import "syscall"

func kill(sid ServerID) {
	stopWithSignal(sid, syscall.SIGKILL)
}
