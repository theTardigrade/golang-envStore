package envStore

import (
	"bytes"
)

var (
//	MultipleEqualsErr error = errors.New("Env: only one equals sign allowed per line.")
)

func parseLine(line string) (key, value string, err error) {
	var buffer bytes.Buffer
	var metEquals bool

	for _, r := range line {
		if r == '=' && !metEquals {
			metEquals = true
			key = string(bytes.ToUpper(buffer.Bytes()))
			buffer.Reset()
			continue
		}

		buffer.WriteRune(r)
	}
	value = buffer.String()

	return
}
