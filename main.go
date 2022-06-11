package main

import (
	"bytes"
	"image"
	"image/png"
	"time"

	"github.com/JamesHovious/w32"
	"github.com/hugolgst/rich-go/client"
	"github.com/oliamb/cutter"
	"github.com/redraskal/star-citizen-rich-presence/rsi"
	"github.com/redraskal/star-citizen-rich-presence/utils"
	"github.com/redraskal/star-citizen-rich-presence/win"
	"github.com/vardius/shutdown"
)

const StarCitizenExe = "StarCitizen.exe"
const DiscordAppID = "983874700440645672"
const CaptureInterval = 10 * time.Second

func main() {
	println("--------------------------------------------------------")
	println("Star Citizen Rich Presence by redraskal.")
	println("https://github.com/redraskal/star-citizen-rich-presence")
	println("--------------------------------------------------------\n")
	go shutdown.GracefulStop(func() {
		println("Disconnecting from Discord...")
		client.Logout()
	})
	loop()
}

func loop() {
	println("Waiting for Star Citizen...")
	hwnd := win.WaitFor(StarCitizenExe, CaptureInterval)
	rsi.UpdateInstallPath()
	rsi.RequireConsoleCmd()
	println("Connecting to Discord...")
	if err := client.Login(DiscordAppID); err != nil {
		panic(err)
	}
	UpdateStartTimestamp(&DefaultActivity)
	if err := UpdateProfile(&DefaultActivity); err != nil {
		panic(err)
	}
	client.SetActivity(DefaultActivity)
	capture_loop(DefaultActivity, hwnd)
}

func capture_loop(a client.Activity, hwnd w32.HWND) {
	s, err := capture(hwnd)
	if err != nil {
		println(err.Error())
		client.Logout()
		time.Sleep(CaptureInterval)
		loop()
		return
	}
	println("\nCurrent Location:", s.Location.Name)
	if err = UpdateActivity(a, s); err != nil {
		println(err.Error())
		client.Logout()
		time.Sleep(CaptureInterval)
		loop()
		return
	}
	time.Sleep(CaptureInterval)
	capture_loop(a, hwnd)
}

func capture(hwnd w32.HWND) (utils.SessionInfo, error) {
	img, err := win.CaptureWindow(hwnd)
	if err != nil {
		return utils.SessionInfo{}, err
	}

	cropped, err := cutter.Crop(img, cutter.Config{
		Width:  430,
		Height: 180,
		Anchor: image.Point{
			X: img.Rect.Dx() - 430,
			Y: 0,
		},
	})
	if err != nil {
		return utils.SessionInfo{}, err
	}

	cropped = utils.PrepareImageForOCR(cropped)

	// file, _ := os.Create("test.png")
	// defer file.Close()
	buf := new(bytes.Buffer)
	if err = png.Encode(buf, cropped); err != nil {
		return utils.SessionInfo{}, err
	}
	// if err = png.Encode(file, cropped); err != nil {
	// 	return SessionInfo{}, err
	// }

	println("Running OCR...\n")

	text, err := utils.Tesseract(buf.Bytes())
	if err != nil {
		return utils.SessionInfo{}, err
	}

	println(text)

	return utils.ParseSessionInfo(text), nil
}
