package storage

type Storage interface {
	Read(offset int64, size int) ([]byte, error)
	Write(data []byte) (int64, error)
	Sync() error
}
