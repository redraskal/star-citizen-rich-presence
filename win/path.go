package win

import (
	"os/exec"
	"strings"
	"syscall"
)

func Path(name string) string {
	child := exec.Command("cmd")
	child.SysProcAttr = &syscall.SysProcAttr{CmdLine: `/C wmic process get ExecutablePath | findstr ` + name}
	out, err := child.Output()
	if err != nil {
		return ""
	}
	line := strings.TrimSpace(string(out))
	return line[0 : len(line)-len(name)-1]
}
