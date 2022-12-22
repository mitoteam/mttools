package mttools

import (
	"sync"
)

type LockableI interface {
	Lock()
	TryLock() bool
	Unlock()
}

type Lockable struct {
	lockMutex sync.Mutex
}

func (o *Lockable) Lock() {
	o.lockMutex.Lock()
}

func (o *Lockable) TryLock() bool {
	return o.lockMutex.TryLock()
}

func (o *Lockable) UnLock() {
	o.lockMutex.Unlock()
}
