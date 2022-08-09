package object

import (
	"bufio"
	"fmt"
	"os"
)

type File struct {
	Filename string
	Reader   *bufio.Reader
	Writer   *bufio.Writer
	Handle   *os.File
}

func (f *File) Type() Type { return FILE }
func (f *File) String() string {
	return fmt.Sprintf("<file:%s>", f.Filename)
}
