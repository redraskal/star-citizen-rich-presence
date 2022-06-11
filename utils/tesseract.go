package utils

import (
	"bytes"
	"os/exec"
	"strings"
	"syscall"
)

var tesseractPath = ""

func init() {
	println("Finding tesseract location...")
	programFiles, err := exec.Command("cmd", "/C", "echo %PROGRAMFILES%").Output()
	if err != nil {
		panic(err)
	}
	tesseractPath = strings.TrimSuffix(string(programFiles), "\r\n") + "\\Tesseract-OCR"
	println("Tesseract:", tesseractPath)
}

func Tesseract(img []byte) (string, error) {
	child := exec.Command("cmd")
	child.SysProcAttr = &syscall.SysProcAttr{CmdLine: `/C tesseract stdin - -l eng`}
	child.Dir = tesseractPath
	child.Stdin = bytes.NewReader(img)
	out, err := child.CombinedOutput()
	if err != nil {
		println(out)
		return "", err
	}
	return string(out), nil
}
