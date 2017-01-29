package env

import (
	"bufio"
	"bytes"
	"os"
)

func (e *environment) LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key, value, err := parseLine(scanner.Text())
		if err != nil {
			return err
		}
		e.Set(key, value)
	}
	if scanner.Err() != nil {
		return err
	}

	return nil
}

func (e *environment) LoadFromString(text string) error {
	for _, line := range bytes.Split([]byte(text), []byte{'\n'}) {
		key, value, err := parseLine(string(line))
		if err != nil {
			return err
		}
		e.Set(key, value)
	}

	return nil
}

func (e *environment) LoadFromSystem() error {
	for _, pair := range os.Environ() {
		key, value, err := parseLine(pair)
		if err != nil {
			return err
		}
		e.Set(key, value)
	}

	return nil
}
