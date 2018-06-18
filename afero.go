package afero

import (
	"errors"
	"io"
	"os"
	"time"
)

// File represents a file in the filesystem.
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	io.WriterAt

	Name() string
	Readdir(count int) ([]os.FileInfo, error)
	Readdirnames(n int) ([]string, error)
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error
	WriteString(s string) (ret int, err error)
}

// Fs is the filesystem interface.
//
// Any simulated or real filesystem should implement this interface.
type Fs interface {
	// Create creates a file in the filesystem, returning the file and an
	// error, if any happens.
	Create(name string) (File, error)

	// Mkdir creates a directory in the filesystem, return an error if any
	// happens.
	Mkdir(name string, perm os.FileMode) error

	// MkdirAll creates a directory path and all parents that does not exist
	// yet.
	MkdirAll(path string, perm os.FileMode) error

	// Open opens a file, returning it or an error, if any happens.
	Open(name string) (File, error)

	// OpenFile opens a file using the given flags and the given mode.
	OpenFile(name string, flag int, perm os.FileMode) (File, error)

	// Remove removes a file identified by name, returning an error, if any
	// happens.
	Remove(name string) error

	// RemoveAll removes a directory path and any children it contains. It
	// does not fail if the path does not exist (return nil).
	RemoveAll(path string) error

	// Rename renames a file.
	Rename(oldname, newname string) error

	// Stat returns a FileInfo describing the named file, or an error, if any
	// happens.
	Stat(name string) (os.FileInfo, error)

	// The name of this FileSystem
	Name() string

	//Chmod changes the mode of the named file to mode.
	Chmod(name string, mode os.FileMode) error

	//Chtimes changes the access and modification times of the named file
	Chtimes(name string, atime time.Time, mtime time.Time) error
}

var (
	ErrFileClosed        = errors.New("File is closed")
	ErrOutOfRange        = errors.New("Out of range")
	ErrTooLarge          = errors.New("Too large")
	ErrFileNotFound      = os.ErrNotExist
	ErrFileExists        = os.ErrExist
	ErrDestinationExists = os.ErrExist
)

var fs = NewMemMapFs()

func Name() string { return fs.Name() }

func Create(name string) (File, error) {
	return fs.Create(name)
}

func Mkdir(name string, perm os.FileMode) error {
	return fs.Mkdir(name, perm)
}

func MkdirAll(path string, perm os.FileMode) error {
	return fs.MkdirAll(path, perm)
}

func Open(name string) (File, error) {
	return fs.Open(name)
}

func OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	return fs.OpenFile(name, flag, perm)
}

func Remove(name string) error {
	return fs.Remove(name)
}

func RemoveAll(path string) error {
	return fs.RemoveAll(path)
}

func Rename(oldname, newname string) error {
	return fs.Rename(oldname, newname)
}

func Stat(name string) (os.FileInfo, error) {
	return fs.Stat(name)
}

func Chmod(name string, mode os.FileMode) error {
	return fs.Chmod(name, mode)
}

func Chtimes(name string, atime time.Time, mtime time.Time) error {
	return fs.Chtimes(name, atime, mtime)
}
