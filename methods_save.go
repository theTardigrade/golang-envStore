package envStore

import (
	"bufio"
	"os"
)

func (e *Environment) SaveToSystem() error {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	for key, value := range e.data {
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}

func (e *Environment) SaveToFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(e.String())
	if err != nil {
		return err
	}
	return w.Flush()
}
