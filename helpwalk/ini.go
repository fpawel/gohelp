package helpwalk

import "github.com/lxn/walk"

func IniStr(key string) string {

	s, _ := walk.App().Settings().Get(key)
	return s
}

func IniPutStr(key, value string) {
	if err := walk.App().Settings().Put(key, value); err != nil {
		panic(err)
	}
}
