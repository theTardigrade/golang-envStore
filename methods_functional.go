package envStore

func (e *Environment) Iterate(callback func(string, string)) {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	for key, value := range e.data {
		callback(key, value)
	}
}

func (e *Environment) Map(callback func(string, string) string) {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	for key, value := range e.data {
		switch newValue := callback(key, value); newValue {
		case value:
			continue
		case "":
			delete(e.data, key)
		default:
			e.data[key] = newValue
		}
	}
}

func (e *Environment) Filter(callback func(string, string) bool) {
	e.lockIfNecessary()
	defer e.unlockIfNecessary()

	for key, value := range e.data {
		if retain := callback(key, value); !retain {
			delete(e.data, key)
		}
	}
}
