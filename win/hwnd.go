package win

import "github.com/JamesHovious/w32"

func HWND(pid int) w32.HWND {
	var hwnd w32.HWND
	for ok := true; ok; ok = (hwnd != 0) {
		hwnd = w32.FindWindowExW(w32.HWND_DESKTOP, hwnd, nil, nil)
		if _, x := w32.GetWindowThreadProcessId(hwnd); x == pid && w32.IsWindowVisible(hwnd) {
			return hwnd
		}
	}
	return hwnd
}
