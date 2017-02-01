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
		switch returnee := callback(key, value); returnee {
		case value:
			continue
		case "":
			delete(e.data, key)
		default:
			e.data[key] = returnee
		}
	}
}
