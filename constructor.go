package envStore

import (
	"sync"
)

func New(cfg Config) (*Environment, error) {
	env := &Environment{
		data:             make(dictionary),
		useMutex:         cfg.UseMutex,
		ignoreEmptyLines: cfg.IgnoreEmptyLines,
		acceptComments:   cfg.AcceptComments,
	}

	if cfg.UseMutex {
		env.mutex = &sync.RWMutex{}
	}

	if cfg.FromFilePaths != nil {
		for _, path := range cfg.FromFilePaths {
			if err := env.LoadFromFile(path); err != nil {
				return nil, err
			}
		}
	}

	if cfg.FromStrings != nil {
		for _, str := range cfg.FromStrings {
			if err := env.LoadFromString(str); err != nil {
				return nil, err
			}
		}
	}

	if cfg.FromEnvironments != nil {
		for _, env2 := range cfg.FromEnvironments {
			env.LoadFromEnviroment(env2)
		}
	}

	if cfg.FromSystem {
		if err := env.LoadFromSystem(); err != nil {
			return nil, err
		}
	}

	return env, nil
}
