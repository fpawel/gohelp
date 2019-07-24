package gohelp

import (
	"fmt"
	"os"
	"sync"
)

func GetEnvWithLog(envvar string) string {
	return getEnvWithLog(envvar)
}

var getEnvWithLog = func() func(string) string {
	var (
		mu sync.Mutex
		m  = make(map[string]struct{})
	)
	return func(envvar string) string {
		s := os.Getenv(envvar)
		mu.Lock()
		defer mu.Unlock()
		if _, f := m[envvar]; !f {
			m[envvar] = struct{}{}
			fmt.Printf("%s=%q\n", envvar, s)
		}
		return s
	}
}()
