package filesystem

import "os"

// Opener is the interface that wraps the file opening methods.
type Opener interface {
	Open(name string) (*os.File, error)
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
}

type opener struct {
	open     func(name string) (*os.File, error)
	openFile func(name string, flag int, perm os.FileMode) (*os.File, error)
}

// NewDefaultOpener constructs an Opener with the default dependencies.
func NewDefaultOpener() Opener {
	return NewOpener(
		os.Open,
		os.OpenFile,
	)
}

// NewOpener constructs a struct that satisfies the Opener interface.
func NewOpener(
	open func(name string) (*os.File, error),
	openFile func(name string, flag int, perm os.FileMode) (*os.File, error),
) Opener {
	return &opener{
		open:     open,
		openFile: openFile,
	}
}

func (o *opener) Open(name string) (*os.File, error) {
	return o.open(name)
}

func (o *opener) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return o.openFile(name, flag, perm)
}
