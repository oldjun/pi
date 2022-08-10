package object

func NewEnvironment(directory string) *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil, directory: directory}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment(outer.directory)
	env.outer = outer
	return env
}

type Environment struct {
	store     map[string]Object
	outer     *Environment
	directory string
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (e *Environment) Del(name string) bool {
	_, ok := e.store[name]
	if ok {
		delete(e.store, name)
	}
	return true
}

func (e *Environment) All() map[string]Object {
	return e.store
}

func (e *Environment) Has(name string) bool {
	_, ok := e.store[name]
	if !ok && e.outer != nil {
		_, ok = e.outer.Get(name)
	}
	return ok
}

func (e *Environment) SetDirectory(directory string) {
	e.directory = directory
}

func (e *Environment) GetDirectory() string {
	directory := e.directory
	if directory == "" && e.outer != nil {
		directory = e.outer.GetDirectory()
	}
	return directory
}
