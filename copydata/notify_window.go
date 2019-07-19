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
	log.Info(x.peerWindowClassName + ": peer closed")
}

func (x *NotifyWindow) InitPeer() {
	x.mu.Lock()
	defer x.mu.Unlock()
	x.initPeer()
}

func (x *NotifyWindow) initPeer() {
	x.hWndPeer = winapp.FindWindow(x.peerWindowClassName)
	if !winapp.IsWindow(x.hWndPeer) {
		log.PrintErr(x.peerWindowClassName + ": init peer: window class not found")
		return
	}
	log.Info(x.peerWindowClassName + ": init peer")
}

func (x *NotifyWindow) sendMsg(msg uintptr, b []byte) {
	x.mu.Lock()
	if x.hWndPeer == 0 {
		x.initPeer()
	}
	hWndPeer := x.hWndPeer
	x.mu.Unlock()

	if hWndPeer != 0 && SendMessage(x.hWnd, hWndPeer, msg, b) == 0 {
		log.PrintErr(fmt.Sprintf("SendMessage failed: %s: %d, %+v, %+v", x.peerWindowClassName, msg, x.hWnd, hWndPeer))
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
