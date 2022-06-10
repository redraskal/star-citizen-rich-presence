package main

import (
	"bytes"
	"os/exec"
	"syscall"
)

func Tesseract(img []byte) (string, error) {
	child := exec.Command("cmd")
	child.SysProcAttr = &syscall.SysProcAttr{CmdLine: `/C tesseract stdin - -l eng`}
	child.Stdin = bytes.NewReader(img)
	out, err := child.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
