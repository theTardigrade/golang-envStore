package envStore

import (
	"sync"
)

type dictionary map[string]string

type environment struct {
	data     dictionary
	mutex    sync.Mutex
	useMutex bool
}

type Config struct {
	FromFilePaths []string
	FromStrings   []string
	FromSystem    bool
	UseMutex      bool
}
