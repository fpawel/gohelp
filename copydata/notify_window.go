package copydata

import (
	"encoding/json"
	"fmt"
	"github.com/fpawel/gohelp/winapp"
	"github.com/fpawel/goutils"
	"github.com/lxn/win"
)

type NotifyWindow struct {
	hWnd, hWndPeer      win.HWND
	peerWindowClassName string
}

func NewNotifyWindow(windowClassName, peerWindowClassName string) *NotifyWindow {
	hWnd := winapp.FindWindow(windowClassName)
	if !winapp.IsWindow(hWnd) {
		hWnd = winapp.NewWindowWithClassName(windowClassName, win.DefWindowProc)
	}
	if !winapp.IsWindow(hWnd) {
		panic(windowClassName)
	}
	return &NotifyWindow{
		peerWindowClassName: peerWindowClassName,
		hWnd:                hWnd,
	}
}

func (x *NotifyWindow) TryClose() bool {
	return win.SendMessage(x.hWnd, win.WM_CLOSE, 0, 0) == 0
}

func (x *NotifyWindow) Close() {
	x.TryClose()
}

func (x *NotifyWindow) init() {
	x.hWndPeer = winapp.FindWindow(x.peerWindowClassName)
}

func (x *NotifyWindow) sendMsg(msg uintptr, b []byte) bool {
	if x.hWndPeer == 0 {
		x.init()
	}
	if x.hWndPeer == 0 {
		return false
	}
	if SendMessage(x.hWnd, x.hWndPeer, msg, b) == 0 {
		x.hWndPeer = 0
		return false
	}
	return true
}

func (x *NotifyWindow) Notify(msg uintptr, a ...interface{}) bool {
	return x.NotifyStr(msg, fmt.Sprint(a...))
}

func (x *NotifyWindow) NotifyStr(msg uintptr, s string) bool {
	return x.sendMsg(msg, goutils.UTF16FromString(s))
}

func (x *NotifyWindow) NotifyJson(msg uintptr, param interface{}) bool {
	b, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}
	return x.Notify(msg, string(b))
}

func (x *NotifyWindow) Notifyf(msg uintptr, format string, a ...interface{}) bool {
	return x.NotifyStr(msg, fmt.Sprintf(format, a...))
}
