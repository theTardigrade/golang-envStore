package env

func (e *environment) Iterate(callback func(string, string)) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for key, value := range e.data {
		callback(key, value)
	}
}

func (e *environment) Map(callback func(string, string) string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

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

