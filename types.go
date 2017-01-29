package env

import (
	"sync"
)

type dictionary map[string]string

type environment struct {
	data  dictionary
	mutex sync.Mutex
}

type Config struct {
	FilePaths  []string
	Strings    []string
	FromSystem bool
}
