package interfaces

import "os"

type Config interface {
	Load(f *os.File) error
}
