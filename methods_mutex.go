package env

func (e *environment) lockIfNecessary() {
	if !e.useMutex {
		return
	}

	e.mutex.Lock()
}

func (e *environment) unlockIfNecessary() {
	if !e.useMutex {
		return
	}

	e.mutex.Unlock()
}
