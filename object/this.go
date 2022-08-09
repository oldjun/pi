package object

import "fmt"

type This struct {
	Instance *Instance
}

func (t *This) Type() Type { return THIS }
func (t *This) String() string {
	return fmt.Sprintf("<this:%s>", t.Instance.Class.Name.Value)
}
