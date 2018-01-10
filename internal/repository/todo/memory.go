package todo

import "github.com/lccezinha/twodo/internal/twodo"

// MemoryRepository holds a fake repository to be used in tests
type MemoryRepository struct {
	data   map[int]*twodo.Todo
	lastID int
}

// Save will save a Todo in a MemoryRepository
func (m *MemoryRepository) Save(t *twodo.Todo) error {
	m.lastID++
	t.ID = m.lastID
	m.data[t.ID] = t
	return nil
}

// ListAll will list all Todos in a MemoryRepository
func (m *MemoryRepository) ListAll() ([]*twodo.Todo, error) {
	var todos []*twodo.Todo

	for _, todo := range m.data {
		todos = append(todos, todo)
	}

	return todos, nil
}

// NewMemoryRepository will return a new instance of MemoryRepository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make(map[int]*twodo.Todo),
	}
}
