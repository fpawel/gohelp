package helpwalk

import
(
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
		CurrentIndex: ComportIndex(IniStr(key)),
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

func ComportIndex(portName string) int {
	ports, _ := comport.Ports()
	return ComboBoxIndex(portName, ports)
}

func ComboBoxIndex(s string, m []string) int {
	for i, x := range m {
		if s == x {
			return i
		}
	}
	return -1
}

func getComports() []string {
	ports, _ := comport.Ports()
	return ports
}