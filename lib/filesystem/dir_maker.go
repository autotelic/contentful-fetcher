package filesystem

import "os"

// DirMaker is the interface that wraps the directory making methods.
type DirMaker interface {
	MkdirAll(path string, perm os.FileMode) error
}

type dirMaker struct {
	mkdirAll func(path string, perm os.FileMode) error
}

// NewDefaultDirMaker constructs a DirMaker with the default dependencies.
func NewDefaultDirMaker() DirMaker {
	return NewDirMaker(
		os.MkdirAll,
	)
}

// NewDirMaker constructs a struct that satisfies the DirMaker interface.
func NewDirMaker(
	mkdirAll func(path string, perm os.FileMode) error,
) DirMaker {
	return &dirMaker{
		mkdirAll: mkdirAll,
	}
}

func (d *dirMaker) MkdirAll(path string, perm os.FileMode) error {
	return d.mkdirAll(path, perm)
}
