package helpwalk

import (
	"github.com/fpawel/comm/comport"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func ComboBoxComport(comboBox **walk.ComboBox, key string) ComboBox {
	var comboBox2 *walk.ComboBox
	if comboBox == nil {
		comboBox = &comboBox2
	}
	return ComboBox{
		AssignTo:     comboBox,
		Model:        getComports(),
		CurrentIndex: comportIndex(IniStr(key)),
		OnMouseDown: func(_, _ int, _ walk.MouseButton) {
			cb := *comboBox
			n := cb.CurrentIndex()
			if err := cb.SetModel(getComports()); err != nil {
				panic(err)
			}
			if err := cb.SetCurrentIndex(n); err != nil {
				panic(err)
			}
		},
		OnCurrentIndexChanged: func() {
			IniPutStr(key, (*comboBox).Text())
		},
	}
}

func ComboBoxWithStringList(comboBox **walk.ComboBox, key string, model []string) ComboBox {
	return ComboBox{
		AssignTo:     comboBox,
		Model:        model,
		CurrentIndex: comboBoxIndex(IniStr(key), model),
		OnCurrentIndexChanged: func() {
			IniPutStr(key, (*comboBox).Text())
		},
	}
}

func comboBoxIndex(s string, m []string) int {
	for i, x := range m {
		if s == x {
			return i
		}
	}
	return -1
}

func comportIndex(portName string) int {
	ports, _ := comport.Ports()
	return comboBoxIndex(portName, ports)
}

func getComports() []string {
	ports, _ := comport.Ports()
	return ports
}
