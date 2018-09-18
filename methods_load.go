package envStore

import (
	"bufio"
	"bytes"
	"os"
)

func (e *Environment) conditionalSet(key, value string, err error) error {
	if err != nil && (err != NoKeyParseErr || !e.ignoreEmptyLines) && (err != CommentParseErr && e.acceptComments) {
		return err
	}

	return e.Set(key, value)
}

func (e *Environment) LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key, value, err := parseLine(scanner.Text())
		err = e.conditionalSet(key, value, err)
		if err != nil {
			return err
		}
	}
	if scanner.Err() != nil {
		return err
	}

	return nil
}

func (e *Environment) LoadFromString(text string) error {
	for _, line := range bytes.Split([]byte(text), []byte{'\n'}) {
		key, value, err := parseLine(string(line))
		err = e.conditionalSet(key, value, err)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Environment) LoadFromEnviroment(e2 *Environment) error {
	e2.readLockIfNecessary()
	defer e2.readUnlockIfNecessary()

	for key, value := range e2.data {
		if err := e.Set(key, value); err != nil {
			return err
		}
	}

	return nil
}

func (e *Environment) LoadFromSystem() error {
	for _, pair := range os.Environ() {
		key, value, err := parseLine(pair)
		if err != nil {
			return err
		}
		if err = e.Set(key, value); err != nil {
			return err
		}
	}

	return nil
}
