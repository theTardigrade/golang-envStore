package envStore

import (
	"sync"
)

func New(cfg *Config) (*Environment, error) {
	env := &Environment{
		data:             make(dictionary),
		useMutex:         cfg.UseMutex,
		ignoreEmptyLines: cfg.IgnoreEmptyLines,
		acceptComments:   cfg.AcceptComments,
		maxKeyLength:     cfg.MaxKeyLength,
	}

	if cfg != nil {
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

		if cfg.FromJSONSlices != nil {
			for _, data := range cfg.FromJSONSlices {
				if err := env.LoadFromJSON(data); err != nil {
					return nil, err
				}
			}
		}

		if cfg.FromEnvironments != nil {
			for _, env2 := range cfg.FromEnvironments {
				if err := env.LoadFromEnvironment(env2); err != nil {
					return nil, err
				}
			}
		}

		if cfg.FromSystem {
			if err := env.LoadFromSystem(); err != nil {
				return nil, err
			}
		}
	}

	return env, nil
}
