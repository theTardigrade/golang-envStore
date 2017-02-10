package envStore

import (
	"bytes"
	"testing"

	"./test"
)

var testParseLineData = []struct {
	line, key, value string
	err              error
}{
	{"x=y", "X", "y", nil},
	{"all men=mortal", "ALL MEN", "mortal", nil},
	{"socrates=a man", "SOCRATES", "a man", nil},
	{"Socrates=mortal", "SOCRATES", "mortal", nil},
	{"3=three=tres", "3", "three=tres", nil},
	{"2!=3", "2!", "3", nil},
	{"shell=/bin/bash", "SHELL", "/bin/bash", nil},
	{
		"JS_code=((msg) => { console.log(msg); })('test');",
		"JS_CODE",
		"((msg) => { console.log(msg); })('test');",
		nil,
	},
	{"Oranges=Apples", "ORANGES", "Apples", nil},
	{"=LockedOut", "", "", NoKeyParseErr},
	{"Nihilist=", "", "", NoValParseErr},
	{
		func() string {
			buffer := bytes.NewBuffer([]byte("test="))
			for i, l := 0, MaxLineLen-buffer.Len()+1; i < l; i++ {
				buffer.WriteRune(0)
			}
			return buffer.String()
		}(),
		"",
		"",
		MaxLineLenParseErr,
	},
}

func TestParseLine(t *testing.T) {
	for _, d := range testParseLineData {
		key, value, err := parseLine(d.line)
		if err != nil {
			if d.err == nil {
				t.Error(err)
			}
			test.AssertEqual(t, "error", d.err, err)
			continue
		}

		test.AssertEqual(t, "key", d.key, key)
		test.AssertEqual(t, "value", d.value, value)
	}

}
