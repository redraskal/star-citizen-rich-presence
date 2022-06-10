package main

import (
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/JamesHovious/w32"
)

func WaitForProcess(name string, interval time.Duration) w32.HWND {
	pid := FindPID(name)
	if pid == -1 {
		return sleep(name, interval)
	}
	println("PID:", pid)
	hwnd := FindHWND(pid)
	if hwnd == 0 {
		return sleep(name, interval)
	}
	println("HWND:", hwnd)
	return hwnd
}

func sleep(name string, interval time.Duration) w32.HWND {
	time.Sleep(interval)
	return WaitForProcess(name, interval)
}

func FindPID(name string) int {
	child := exec.Command("cmd")
	child.SysProcAttr = &syscall.SysProcAttr{CmdLine: `/C tasklist /fi "IMAGENAME eq ` + name + `" /fi "STATUS eq RUNNING" /fo list | findstr PID:`}
	out, err := child.Output()
	if err != nil {
		return -1
	}
	split := strings.Split(strings.ReplaceAll(string(out), " ", ""), "PID:")
	if len(split) < 2 {
		return -1
	}
	n, err := strconv.Atoi(strings.TrimRight(split[1], "\r\n"))
	if err != nil {
		return -1
	}
	return n
}

func FindHWND(pid int) w32.HWND {
	var hwnd w32.HWND
	for ok := true; ok; ok = (hwnd != 0) {
		hwnd = w32.FindWindowExW(w32.HWND_DESKTOP, hwnd, nil, nil)
		if _, x := w32.GetWindowThreadProcessId(hwnd); x == pid && w32.IsWindowVisible(hwnd) {
			return hwnd
		}
	}
	return hwnd
}
