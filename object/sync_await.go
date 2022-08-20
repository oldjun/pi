package object

import (
	"fmt"
	"sync"
)

type SyncAwait struct {
	Value *sync.WaitGroup
}

func (s *SyncAwait) Type() Type { return SYNC_AWAIT }
func (s *SyncAwait) String() string {
	return fmt.Sprintf("<sync_await:%v>", s.Value)
}

func (s *SyncAwait) Method(method string, args []Object) Object {
	switch method {
	case "add":
		return s.add(args)
	case "done":
		return s.done(args)
	case "wait":
		return s.wait(args)
	}
	return NewError("sync await undefined method: %s", method)
}

func (s *SyncAwait) add(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. await.add() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		s.Value.Add(int(arg.Value))
	default:
		return NewError("wrong type of arguments. await.add() got=%s", arg.Type())
	}
	return &Null{}
}

func (s *SyncAwait) done(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. await.done() got=%d", len(args))
	}
	s.Value.Done()
	return &Null{}
}

func (s *SyncAwait) wait(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. await.wait() got=%d", len(args))
	}
	s.Value.Wait()
	return &Null{}
}
