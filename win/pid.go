package win

import (
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func PID(name string) int {
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
