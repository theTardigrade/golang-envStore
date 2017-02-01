package envStore

func (e *Environment) lockIfNecessary() {
	if !e.useMutex {
		return
	}

	e.mutex.Lock()
}

func (e *Environment) unlockIfNecessary() {
	if !e.useMutex {
		return
	}

	e.mutex.Unlock()
}
