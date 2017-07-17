package envStore

import (
	"bytes"
	"errors"
	"strings"
)

const (
	MaxLineLen = 1 << 12
)

var (
	CommentParseErr    error = errors.New("Line is commented out.")
	NoKeyParseErr      error = errors.New("Key must be provided.")
	NoValParseErr      error = errors.New("Value must be provided.")
	MaxLineLenParseErr error = errors.New("Line length is greater than the maximum allowed.")
)

func parseLine(line string) (key, value string, err error) {
	line = strings.TrimSpace(line)

	if len(line) > MaxLineLen {
		err = MaxLineLenParseErr
		return
	}

	var buffer bytes.Buffer
	var metEquals bool

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
	} else if key[0] == '#' {
		err = CommentParseErr
	} else if buffer.Len() == 0 {
		err = NoValParseErr
	} else {
		value = buffer.String()
	}

	return
}
