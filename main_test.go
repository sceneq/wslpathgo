package main

import (
	"testing"
	"testing/quick"
)

func TestWslpath(t *testing.T) {
	cases := []struct {
		win string
		wsl string
	}{{
		win: `C:\Users\Example`,
		wsl: `/mnt/c/Users/Example`,
	}, {
		win: `.`,
		wsl: `.`,
	}, {
		win: `..\f`,
		wsl: `../f`,
		// TODO: it seems to be an absolute path in the original wslpath
		// wsl: `/mnt/c/Users/iiiii/git/f`,
	}, {
		win: `C:\Program Files`,
		wsl: `/mnt/c/Program Files`,
	}}

	for _, tt := range cases {
		if result := WindowsToWSL(tt.win); result != tt.wsl {
			t.Errorf("WindowsToWSL(%q) = %q; want %q", tt.win, result, tt.wsl)
		}

		if result := WSLToWindows(tt.wsl); result != tt.win {
			t.Errorf("WSLToWindows(%q) = %q; want %q", tt.wsl, result, tt.win)
		}
	}
}

func TestWindowsToWSLAndBack(t *testing.T) {
	assert := func(win string) bool {
		wsl := WindowsToWSL(win)
		res := WSLToWindows(wsl)
		return res == win
	}
	if err := quick.Check(assert, nil); err != nil {
		t.Error(err)
	}
}

func TestWSLToWindowsAndBack(t *testing.T) {
	assert := func(wsl string) bool {
		win := WSLToWindows(wsl)
		res := WindowsToWSL(win)
		return res == wsl
	}
	if err := quick.Check(assert, nil); err != nil {
		t.Error(err)
	}
}
