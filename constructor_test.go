package envStore

import (
	"sync"
	"testing"

	internalTest "github.com/theTardigrade/envStore/internal/test"
)

var testNewData = []struct {
	cfg *Config
	env *Environment
	err error
}{
	{
		&Config{},
		&Environment{
			data:             make(dictionary),
			useMutex:         false,
			ignoreEmptyLines: false,
		},
		nil,
	},
	{
		&Config{
			UseMutex:         true,
			IgnoreEmptyLines: true,
			FromStrings:      []string{"key=value"},
		},
		&Environment{
			data:             dictionary{"KEY": "value"},
			mutex:            &sync.RWMutex{},
			useMutex:         true,
			ignoreEmptyLines: true,
		},
		nil,
	},
}

func TestNew(t *testing.T) {
	for _, d := range testNewData {
		env, err := New(d.cfg)
		if err != nil {
			if d.err == nil {
				t.Error(err)
			}

			internalTest.AssertEqual(t, "error", d.err, err)
			continue
		}

		internalTest.AssertEqual(t, "environment", d.env, env)
	}
}
