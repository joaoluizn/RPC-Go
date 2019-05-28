package storage

const (
	StorageServiceName = "Storage"
)

func NewStorage() *Storage {
	return &Storage{}
}

type Storage struct{}

func (s Storage) HelloStorage() string {
	return "A Hello message from Storage!\n"
}

func (s Storage) Create() string {
	return "Create Function\n"
}

func (s Storage) Read() string {
	return "Read Function\n"
}

func (s Storage) Update() string {
	return "Update Function\n"
}

func (s Storage) Delete() string {
	return "Delete Function\n"
}
