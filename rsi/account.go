package rsi

import (
	"bufio"
	"os"
	"path"
	"strings"
)

const ProfileEndpoint = "https://robertsspaceindustries.com/citizens/"

func Path() string {
	return "D:\\Roberts Space Industries\\StarCitizen\\LIVE"
}

func OpenLog() (*os.File, *bufio.Scanner, error) {
	file, err := os.Open(path.Join(Path(), "Game.log"))
	if err != nil {
		return nil, nil, err
	}
	return file, bufio.NewScanner(file), nil
}

func Username() (string, error) {
	file, scanner, err := OpenLog()
	if err != nil {
		return "", err
	}
	defer file.Close()
	println("Scanning log for username...")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[CIG-net] User Login Success - Handle[") {
			return strings.Split(strings.Split(line, "[CIG-net] User Login Success - Handle[")[1], "]")[0], nil
		}
	}
	return "", nil
}
