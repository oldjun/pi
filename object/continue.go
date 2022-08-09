package object

type Continue struct{}

func (c *Continue) Type() Type     { return CONTINUE }
func (c *Continue) String() string { return "continue" }
