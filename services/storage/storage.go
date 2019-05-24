package storage

const (
	ServiceName = "Storage"
)

func NewStorage() *Storage {
	return &Storage{}
}

type Storage struct{}
