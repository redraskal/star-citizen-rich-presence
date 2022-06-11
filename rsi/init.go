package rsi

import (
	"os"
	"path"
	"strings"

	"github.com/redraskal/star-citizen-rich-presence/win"
)

const LauncherExe = "StarCitizen_Launcher.exe"

var (
	installPath = ""
)

func UpdateInstallPath() {
	println("Detecting Star Citizen installation folder...")
	installPath = win.Path(LauncherExe)
	if installPath == "" {
		println("Could not locate installation folder.")
		return
	}
	println("Installation folder:", installPath)
}

func RequireConsoleCmd() {
	println("Checking for r_DisplayInfo...")
	file, scanner, err := OpenFile("user.cfg")
	cfg := path.Join(installPath, "user.cfg")
	if err != nil {
		err = os.WriteFile(cfg, []byte("r_DisplayInfo = 2"), 0644)
		if err != nil {
			panic(err)
		}
		println("r_DisplayInfo applied to new user.cfg file.")
		return
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "r_DisplayInfo") {
			println("r_DisplayInfo found, skipping.")
			file.Close()
			return
		}
	}
	file.Close()
	file, err = os.OpenFile(cfg, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString("r_DisplayInfo = 2\n")
	if err != nil {
		panic(err)
	}
	println("r_DisplayInfo appended to user.cfg")
}
