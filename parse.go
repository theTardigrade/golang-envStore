package envStore

import (
	"bytes"
	"errors"
)

const (
	MaxLineLen = 1 << 12
)

var (
	NoKeyParseErr      error = errors.New("Key must be provided.")
	NoValParseErr      error = errors.New("Value must be provided.")
	MaxLineLenParseErr error = errors.New("Line length is greater than the maximum allowed.")
)

func parseLine(line string) (key, value string, err error) {
	var buffer bytes.Buffer
	var metEquals bool

	if len(line) > MaxLineLen {
		err = MaxLineLenParseErr
		return
	}

	for _, r := range line {
		if r == '=' && !metEquals {
			metEquals = true
			if buffer.Len() == 0 {
				break
			}
			key = string(bytes.ToUpper(buffer.Bytes()))
			buffer.Reset()
			continue
		}

		buffer.WriteRune(r)
	}

	if key == "" {
		err = NoKeyParseErr
	} else if buffer.Len() == 0 {
		err = NoValParseErr
	} else {
		value = buffer.String()
	}

	return
}
