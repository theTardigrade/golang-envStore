package envStore

func New(cfg Config) (*Environment, error) {
	env := &Environment{
		data:             make(dictionary),
		useMutex:         cfg.UseMutex,
		ignoreEmptyLines: cfg.IgnoreEmptyLines,
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

	if cfg.FromSystem {
		if err := env.LoadFromSystem(); err != nil {
			return nil, err
		}
	}

	return env, nil
}
