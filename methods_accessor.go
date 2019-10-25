package envStore

import (
	"bytes"
	"encoding/json"
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
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	e.readLockIfNecessary()
	value, ok := e.data[key]
	e.readUnlockIfNecessary()

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

func (e *Environment) GetInt8(key string) (value int8, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value64, err := strconv.ParseInt(rawValue, 10, 8)
	if err != nil {
		return
	}

	value = int8(value64)
	return
}

func (e *Environment) GetInt16(key string) (value int16, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value64, err := strconv.ParseInt(rawValue, 10, 16)
	if err != nil {
		return
	}

	value = int16(value64)
	return
}

func (e *Environment) GetInt32(key string) (value int32, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value64, err := strconv.ParseInt(rawValue, 10, 32)
	if err != nil {
		return
	}

	value = int32(value64)
	return
}

func (e *Environment) GetInt64(key string) (value int64, err error) {
	rawValue, err := e.Get(key)
	if err != nil {
		return
	}

	value, err = strconv.ParseInt(rawValue, 10, 64)
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

func (e *Environment) MustGetInt8(key string) (value int8) {
	value, err := e.GetInt8(key)
	if err != nil {
		mustGetPanic(err, key)
	}

	return
}

func (e *Environment) MustGetInt16(key string) (value int16) {
	value, err := e.GetInt16(key)
	if err != nil {
		mustGetPanic(err, key)
	}

	return
}

func (e *Environment) MustGetInt32(key string) (value int32) {
	value, err := e.GetInt32(key)
	if err != nil {
		mustGetPanic(err, key)
	}

	return
}

func (e *Environment) MustGetInt64(key string) (value int64) {
	value, err := e.GetInt64(key)
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

func (e *Environment) LazyGet(key string) (value string) {
	if prospectiveValue, err := e.Get(key); err == nil {
		value = prospectiveValue
	}

	return
}

func (e *Environment) LazyGetByteSlice(key string) (value []byte) {
	if prospectiveValue, err := e.GetByteSlice(key); err == nil {
		value = prospectiveValue
	}

	return
}

func (e *Environment) LazyGetInt(key string) (value int) {
	if prospectiveValue, err := e.GetInt(key); err == nil {
		value = prospectiveValue
	}

	return
}

func (e *Environment) LazyGetInt8(key string) (value int8) {
	if prospectiveValue, err := e.GetInt8(key); err == nil {
		value = prospectiveValue
	}

	return
}

func (e *Environment) LazyGetInt16(key string) (value int16) {
	if prospectiveValue, err := e.GetInt16(key); err == nil {
		value = prospectiveValue
	}

	return
}

func (e *Environment) LazyGetInt32(key string) (value int32) {
	if prospectiveValue, err := e.GetInt32(key); err == nil {
		value = prospectiveValue
	}

	return
}

func (e *Environment) LazyGetInt64(key string) (value int64) {
	if prospectiveValue, err := e.GetInt64(key); err == nil {
		value = prospectiveValue
	}

	return
}

func (e *Environment) LazyGetUint(key string) (value uint) {
	if prospectiveValue, err := e.GetUint(key); err == nil {
		value = prospectiveValue
	}

	return
}

func (e *Environment) LazyGetFloat(key string) (value float64) {
	if prospectiveValue, err := e.GetFloat(key); err == nil {
		value = prospectiveValue
	}

	return
}

func (e *Environment) LazyGetBool(key string) (value bool) {
	if prospectiveValue, err := e.GetBool(key); err == nil {
		value = prospectiveValue
	}

	return
}

func (e *Environment) LazyGetDuration(key string) (value time.Duration) {
	if prospectiveValue, err := e.GetDuration(key); err == nil {
		value = prospectiveValue
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
	e.data[key] = value
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetInt(key string, value int) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.Itoa(value)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) MustSet(key, value string) {
	if err := e.Set(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) Unset(key string) (err error) {
	if key, err = e.formatKey(key); err != nil {
		return
	}

	e.writeLockIfNecessary()
	delete(e.data, key)
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) MustUnset(key string) {
	if err := e.Unset(key); err != nil {
		panic(err)
	}
}

func (e *Environment) Clear() {
	e.writeLockIfNecessary()
	e.data = make(dictionary)
	e.writeUnlockIfNecessary()
}

func (e *Environment) Contains(key string) (found bool, err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	e.readLockIfNecessary()
	_, found = e.data[key]
	e.readUnlockIfNecessary()

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

	l := len(e.data)
	pairs := make([][]string, l)

	for k, v := range e.data {
		l--
		pairs[l] = make([]string, 2)
		pairs[l][0], pairs[l][1] = k, v
	}

	return pairs
}

func (e *Environment) String() string {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	var buffer bytes.Buffer

	for k, v := range e.data {
		buffer.WriteString(k)
		buffer.WriteByte('=')
		buffer.WriteString(v)
		buffer.WriteByte('\n')
	}

	return buffer.String()
}

func (e *Environment) JSON() ([]byte, error) {
	e.readLockIfNecessary()
	defer e.readUnlockIfNecessary()

	return json.Marshal(e.data)
}
