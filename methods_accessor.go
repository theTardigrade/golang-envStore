package envStore

import (
	"bytes"
	"strings"
)

func (e *Environment) Get(key string) (value string, ok bool) {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	value, ok = e.data[strings.ToUpper(key)]
	return
}

func (e *Environment) MustGet(key string) (value string) {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	value, ok := e.data[strings.ToUpper(key)]
	if !ok {
		panic("not found")
	}

	return
}

func (e *Environment) Set(key, value string) {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	e.data[strings.ToUpper(key)] = value
}

func (e *Environment) Unset(key string) {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	delete(e.data, strings.ToUpper(key))
}

func (e *Environment) Clear() {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	e.data = make(dictionary)
}

func (e *Environment) Contains(key string) bool {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	_, ok := e.data[strings.ToUpper(key)]
	return ok
}

func (e *Environment) Count() int {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	return len(e.data)
}

func (e *Environment) Keys() []string {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	i := 0
	keys := make([]string, len(e.data))
	for k := range e.data {
		keys[i] = k
		i++
	}

	return keys
}

func (e *Environment) Values() []string {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	i := 0
	values := make([]string, len(e.data))
	for _, v := range e.data {
		values[i] = v
		i++
	}

	return values
}

func (e *Environment) Pairs() [][]string {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	i := 0
	pairs := make([][]string, len(e.data))
	for k, v := range e.data {
		pairs[i] = make([]string, 2)
		pairs[i][0], pairs[i][1] = k, v
		i++
	}

	return pairs
}

func (e *Environment) String() string {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	var buffer bytes.Buffer
	var passedFirstIteration bool

	for k, v := range e.data {
		if passedFirstIteration {
			buffer.WriteRune('\n')
		} else {
			passedFirstIteration = true
		}
		buffer.WriteString(k)
		buffer.WriteRune('=')
		buffer.WriteString(v)
	}

	return buffer.String()
}
