package object

import (
	"fmt"
	"sync"
)

type SyncMutex struct {
	Value *sync.Mutex
}

func (s *SyncMutex) Type() Type { return SYNC_MUTEX }
func (s *SyncMutex) String() string {
	return fmt.Sprintf("<sync_mutex:%v>", s.Value)
}

func (s *SyncMutex) Method(method string, args []Object) Object {
	switch method {
	case "lock":
		return s.lock(args)
	case "unlock":
		return s.unlock(args)
	}
	return NewError("sync mutex undefined method: %s", method)
}

func (s *SyncMutex) lock(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. mutex.lock() got=%d", len(args))
	}
	s.Value.Lock()
	return &Null{}
}

func (s *SyncMutex) unlock(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. mutex.unlock() got=%d", len(args))
	}
	s.Value.Unlock()
	return &Null{}
}
