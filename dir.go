package gohelp

import "os"

func EnsuredDir(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) { // создать каталог если его нет
		err = os.MkdirAll(dir, os.ModePerm)
	}
	return err
}
