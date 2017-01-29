package envStore

import (
	"strings"
)

func (e *environment) Get(key string) (value string, ok bool) {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	value, ok = e.data[strings.ToUpper(key)]
	return
}

func (e *environment) Set(key, value string) {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	e.data[strings.ToUpper(key)] = value
}

func (e *environment) Unset(key string) {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	delete(e.data, strings.ToUpper(key))
}

func (e *environment) Clear() {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	e.data = make(dictionary)
}

func (e *environment) Contains(key string) bool {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	_, ok := e.data[strings.ToUpper(key)]
	return ok
}

func (e *environment) Count() int {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	return len(e.data)
}

func (e *environment) Keys() []string {
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

func (e *environment) Values() []string {
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

func (e *environment) Pairs() [][]string {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	i := 0
	pairs := make([][]string, len(e.data))
	for k, v := range e.data {
		pairs[i] = make([]string, 2)
		pairs[i][0], pairs[i][1] = k, v
	}

	return pairs
}
