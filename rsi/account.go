package rsi

import (
	"bufio"
	"os"
	"path"
	"strings"
)

const ProfileEndpoint = "https://robertsspaceindustries.com/citizens/"

var username = ""

func OpenFile(name string) (*os.File, *bufio.Scanner, error) {
	file, err := os.Open(path.Join(installPath, name))
	if err != nil {
		return nil, nil, err
	}
	return file, bufio.NewScanner(file), nil
}

func Username() (string, error) {
	if username != "" {
		return username, nil
	}
	file, scanner, err := OpenFile("Game.log")
	if err != nil {
		return "", err
	}
	defer file.Close()
	println("Scanning log for username...")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[CIG-net] User Login Success - Handle[") {
			username = strings.Split(strings.Split(line, "[CIG-net] User Login Success - Handle[")[1], "]")[0]
			return username, nil
		}
	}
	return "", nil
}
