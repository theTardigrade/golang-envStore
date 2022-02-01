package envStore

import (
	"strconv"
	"strings"
	"time"
)

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

func (e *Environment) SetFields(key string, value []string) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strings.Join(value, " ")

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetByteSlice(key string, value []byte) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := string(value)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
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

func (e *Environment) SetInt8(key string, value int8) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatInt(int64(value), 10)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetInt16(key string, value int16) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatInt(int64(value), 10)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetInt32(key string, value int32) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatInt(int64(value), 10)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetInt64(key string, value int64) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatInt(value, 10)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetUint(key string, value uint) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatUint(uint64(value), 10)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetUint8(key string, value uint8) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatUint(uint64(value), 10)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetUint16(key string, value uint16) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatUint(uint64(value), 10)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetUint32(key string, value uint32) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatUint(uint64(value), 10)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetUint64(key string, value uint64) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatUint(value, 10)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetFloat32(key string, value float32) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatFloat(float64(value), 'e', -1, 32)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetFloat64(key string, value float64) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatFloat(value, 'e', -1, 32)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetBool(key string, value bool) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := strconv.FormatBool(value)

	e.writeLockIfNecessary()
	e.data[key] = trueValue
	e.writeUnlockIfNecessary()

	return
}

func (e *Environment) SetDuration(key string, value time.Duration) (err error) {
	key, err = e.formatKey(key)
	if err != nil {
		return
	}

	trueValue := value.String()

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

func (e *Environment) MustSetFields(key string, value []string) {
	if err := e.SetFields(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetByteSlice(key string, value []byte) {
	if err := e.SetByteSlice(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetInt(key string, value int) {
	if err := e.SetInt(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetInt8(key string, value int8) {
	if err := e.SetInt8(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetInt16(key string, value int16) {
	if err := e.SetInt16(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetInt32(key string, value int32) {
	if err := e.SetInt32(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetInt64(key string, value int64) {
	if err := e.SetInt64(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetUint(key string, value uint) {
	if err := e.SetUint(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetUint8(key string, value uint8) {
	if err := e.SetUint8(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetUint16(key string, value uint16) {
	if err := e.SetUint16(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetUint32(key string, value uint32) {
	if err := e.SetUint32(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetUint64(key string, value uint64) {
	if err := e.SetUint64(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetFloat32(key string, value float32) {
	if err := e.SetFloat32(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetFloat64(key string, value float64) {
	if err := e.SetFloat64(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetBool(key string, value bool) {
	if err := e.SetBool(key, value); err != nil {
		panic(err)
	}
}

func (e *Environment) MustSetDuration(key string, value time.Duration) {
	if err := e.SetDuration(key, value); err != nil {
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
