package win

import (
	"time"

	"github.com/JamesHovious/w32"
)

func WaitFor(name string, interval time.Duration) w32.HWND {
	pid := PID(name)
	if pid == -1 {
		return sleep(name, interval)
	}
	println("PID:", pid)
	hwnd := HWND(pid)
	if hwnd == 0 {
		return sleep(name, interval)
	}
	println("HWND:", hwnd)
	return hwnd
}

func sleep(name string, interval time.Duration) w32.HWND {
	time.Sleep(interval)
	return WaitFor(name, interval)
}
