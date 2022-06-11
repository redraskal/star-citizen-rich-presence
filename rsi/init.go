package rsi

import "github.com/redraskal/star-citizen-rich-presence/win"

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
