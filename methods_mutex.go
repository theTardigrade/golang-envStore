package envStore

func (e *Environment) writeLockIfNecessary() {
	if e.useMutex {
		e.mutex.Lock()
	}
}

func (e *Environment) writeUnlockIfNecessary() {
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
