package copydata

import (
	"encoding/json"
	"fmt"
	"github.com/fpawel/gohelp/winapp"
	"github.com/fpawel/goutils"
	"github.com/lxn/win"
	"github.com/powerman/structlog"
	"sync"
)

type NotifyWindow struct {
	hWnd, hWndPeer      win.HWND
	peerWindowClassName string
	mu                  sync.Mutex
}

func NewNotifyWindow(windowClassName, peerWindowClassName string) *NotifyWindow {
	return &NotifyWindow{
		peerWindowClassName: peerWindowClassName,
		hWnd:                winapp.NewWindowWithClassName(windowClassName, win.DefWindowProc),
	}
}

func (x *NotifyWindow) TryClose() bool {
	return win.SendMessage(x.hWnd, win.WM_CLOSE, 0, 0) == 0
}

func (x *NotifyWindow) Close() {
	x.TryClose()
}

func (x *NotifyWindow) ResetPeer() {
	x.mu.Lock()
	defer x.mu.Unlock()
	x.hWndPeer = 0
	log.Debug(x.peerWindowClassName + ": peer closed")
}

func (x *NotifyWindow) InitPeer() {
	x.mu.Lock()
	defer x.mu.Unlock()
	x.initPeer()
}

func (x *NotifyWindow) initPeer() {
	x.hWndPeer = winapp.FindWindow(x.peerWindowClassName)
}

func (x *NotifyWindow) sendMsg(msg uintptr, b []byte) {
	x.mu.Lock()
	defer x.mu.Unlock()
	if x.hWndPeer == 0 {
		x.initPeer()
	}
	if SendMessage(x.hWnd, x.hWndPeer, msg, b) == 0 {
		x.hWndPeer = 0
	}
}

func (x *NotifyWindow) Notify(msg uintptr, a ...interface{}) {
	x.NotifyStr(msg, fmt.Sprint(a...))
}

func (x *NotifyWindow) NotifyStr(msg uintptr, s string) {
	x.sendMsg(msg, goutils.UTF16FromString(s))
}

func (x *NotifyWindow) NotifyJson(msg uintptr, param interface{}) {
	b, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}
	x.Notify(msg, string(b))
}

func (x *NotifyWindow) Notifyf(msg uintptr, format string, a ...interface{}) {
	x.NotifyStr(msg, fmt.Sprintf(format, a...))
}

var (
	log = structlog.New()
)
