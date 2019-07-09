package envStore

import (
	"sync"
)

type dictionary map[string]string

type Environment struct {
	data             dictionary
	mutex            *sync.RWMutex
	useMutex         bool
	ignoreEmptyLines bool
	acceptComments   bool
	maxKeyLength     int
}

type Config struct {
	FromFilePaths    []string
	FromStrings      []string
	FromJSONSlices   [][]byte
	FromEnvironments []*Environment
	FromSystem       bool
	UseMutex         bool
	IgnoreEmptyLines bool
	AcceptComments   bool
	MaxKeyLength     int
}
