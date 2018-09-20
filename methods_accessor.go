package envStore

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	ErrKeyNotFound     = errors.New("environment variable not found")
	ErrKeyLenBeyondMax = errors.New("key length is greater than the maximum allowed")
)

func (e *Environment) Get(key string) (value string, err error) {
	if e.maxKeyLength > 0 && len(key) > e.maxKeyLength {
		err = ErrKeyLenBeyondMax
		return
	}

	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	value, ok := e.data[strings.ToUpper(key)]
	if !ok {
		err = ErrKeyNotFound
	}

	return
}

func (e *Environment) GetByteSlice(key string) (value []byte, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value = []byte(rawValue)
	return
}

func (e *Environment) GetInt(key string) (value int, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value, err = strconv.Atoi(rawValue)
	return
}

func (e *Environment) GetUint(key string) (value uint, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value64, err := strconv.ParseUint(rawValue, 10, 0)
	value = uint(value64)
	return
}

func (e *Environment) GetFloat(key string) (value float64, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value, err = strconv.ParseFloat(rawValue, 64)
	return
}

func (e *Environment) GetBool(key string) (value bool, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value, err = strconv.ParseBool(rawValue)
	return
}

func (e *Environment) GetDuration(key string) (value time.Duration, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value, err = time.ParseDuration(rawValue)
	return
}

func mustGetPanic(err error, key string) {
	msg := err.Error()

	if err == ErrKeyNotFound {
		msg += " [" + strings.ToUpper(key) + "]"
	}

	panic(msg)
}

func (e *Environment) MustGet(key string) (value string) {
	value, err := e.Get(key)
	if err != nil {
		mustGetPanic(err, key)
	}

	return
}

func (e *Environment) MustGetByteSlice(key string) (value []byte) {
	value, err := e.GetByteSlice(key)
	if err != nil {
		mustGetPanic(err, key)
	}

	return
}

func (e *Environment) MustGetInt(key string) (value int) {
	value, err := e.GetInt(key)
	if err != nil {
		mustGetPanic(err, key)
	}

	return
}

func (e *Environment) MustGetUint(key string) (value uint) {
	value, err := e.GetUint(key)
	if err != nil {
		mustGetPanic(err, key)
	}

	return
}

func (e *Environment) MustGetFloat(key string) (value float64) {
	value, err := e.GetFloat(key)
	if err != nil {
		mustGetPanic(err, key)
	}

	return
}

func (e *Environment) MustGetBool(key string) (value bool) {
	value, err := e.GetBool(key)
	if err != nil {
		mustGetPanic(err, key)
	}

	return
}

func (e *Environment) MustGetDuration(key string) (value time.Duration) {
	value, err := e.GetDuration(key)
	if err != nil {
		mustGetPanic(err, key)
	}

	return
}

func (e *Environment) formatKey(key string) (string, error) {
	key = strings.ToUpper(key)

	if e.maxKeyLength > 0 && len(key) > e.maxKeyLength {
		return "", ErrKeyLenBeyondMax
	}

	return key, nil
}

func (e *Environment) Set(key, value string) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	e.writeLockIfNecessary()
	defer e.writeUnlockIfNecessary()

	e.data[key] = value
	return
}

func (e *Environment) MustSet(key, value string) {
	if err := e.Set(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) Unset(key string) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	e.writeLockIfNecessary()
	defer e.writeUnlockIfNecessary()

	delete(e.data, key)
	return
}

func (e *Environment) MustUnset(key string) {
	if err := e.Unset(key); err != nil {
		panic(err)
	}
}

func (e *Environment) Clear() {
	e.writeLockIfNecessary()
	defer e.writeUnlockIfNecessary()

	e.data = make(dictionary)
}

func (e *Environment) Contains(key string) (found bool, err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	_, found = e.data[key]
	return
}

func (e *Environment) MustContains(key string) (found bool) {
	found, err := e.Contains(key)
	if err != nil {
		panic(err)
	}

	return
}

func (e *Environment) Count() int {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	return len(e.data)
}

func (e *Environment) Keys() []string {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	i := 0
	keys := make([]string, len(e.data))
	for k := range e.data {
		keys[i] = k
		i++
	}

	return keys
}

func (e *Environment) Values() []string {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	i := 0
	values := make([]string, len(e.data))
	for _, v := range e.data {
		values[i] = v
		i++
	}

	return values
}

func (e *Environment) Pairs() [][]string {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

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
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

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
