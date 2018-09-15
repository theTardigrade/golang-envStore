package envStore

func (e *Environment) lockIfNecessary() {
	if e.useMutex {
		e.mutex.Lock()
	}
}

func (e *Environment) unlockIfNecessary() {
	if e.useMutex {
		e.mutex.Unlock()
	}
}

func (e *Environment) readLockIfNecessary() {
	if e.useMutex {
		e.mutex.RLock()
	}
}

func (e *Environment) readUnlockIfNecessary() {
	if e.useMutex {
		e.mutex.RUnlock()
	}
}
