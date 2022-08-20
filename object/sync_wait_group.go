package object

import (
	"fmt"
	"sync"
)

type SyncWaitGroup struct {
	Value *sync.WaitGroup
}

func (s *SyncWaitGroup) Type() Type { return SYNC_WAIT_GROUP }
func (s *SyncWaitGroup) String() string {
	return fmt.Sprintf("<sync_wait_group:%v>", s.Value)
}

func (s *SyncWaitGroup) Method(method string, args []Object) Object {
	switch method {
	case "add":
		return s.add(args)
	case "done":
		return s.done(args)
	case "wait":
		return s.wait(args)
	}
	return nil
}

func (s *SyncWaitGroup) add(args []Object) Object {
	if len(args) != 1 {
		return NewError("wrong number of arguments. wait_group.add() got=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *Integer:
		s.Value.Add(int(arg.Value))
	default:
		return NewError("wrong type of arguments. wait_group.add() got=%s", arg.Type())
	}
	return &Null{}
}

func (s *SyncWaitGroup) done(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. wait_group.done() got=%d", len(args))
	}
	s.Value.Done()
	return &Null{}
}

func (s *SyncWaitGroup) wait(args []Object) Object {
	if len(args) != 0 {
		return NewError("wrong number of arguments. wait_group.wait() got=%d", len(args))
	}
	s.Value.Wait()
	return &Null{}
}
