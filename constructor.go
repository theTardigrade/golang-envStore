package env

func New(cfg Config) (*environment, error) {
	env := &environment{
		data: make(dictionary),
	}

	if cfg.FilePaths != nil {
		for _, path := range cfg.FilePaths {
			if err := env.LoadFromFile(path); err != nil {
				return nil, err
			}
		}
	}

	if cfg.Strings != nil {
		for _, str := range cfg.Strings {
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
