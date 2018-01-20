package todo

import (
	"errors"
	"sort"

	"github.com/lccezinha/twodo/internal/twodo"
)

var errNotFound = errors.New("not found")

// MemoryRepository holds a fake repository to be used in tests
type MemoryRepository struct {
	data   map[int]*twodo.Todo
	lastID int
}

// sortedTodos implements the sort interface to order Todos from most recent to most old
type sortedTodos []*twodo.Todo

func (st sortedTodos) Len() int           { return len(st) }
func (st sortedTodos) Swap(i, j int)      { st[i], st[j] = st[j], st[i] }
func (st sortedTodos) Less(i, j int) bool { return st[i].ID > st[j].ID }

// func (st sortedTodos) Less(i, j int) bool { return st[i].CreatedAt.After(st[j].CreatedAt) }

// Save will save a Todo in a MemoryRepository
func (m *MemoryRepository) Save(t *twodo.Todo) error {
	m.lastID++
	t.ID = m.lastID
	m.data[t.ID] = t
	return nil
}

// ListAll will list all Todos in a MemoryRepository
func (m *MemoryRepository) ListAll() ([]*twodo.Todo, error) {
	todos := make(sortedTodos, 0)

	for _, todo := range m.data {
		todos = append(todos, todo)
	}

	sort.Sort(todos)

	return todos, nil
}

// Destroy will destroy a single Todo
func (m *MemoryRepository) Destroy(id int) error {
	if _, ok := m.data[id]; ok {
		delete(m.data, id)
		return nil
	}

	return errNotFound
}

// Update will update a single Todo
func (m *MemoryRepository) Update(id int, done bool) error {
	if todo, ok := m.data[id]; ok {
		todo.Done = done
		return nil
	}

	return errNotFound
}

// NewMemoryRepository will return a new instance of MemoryRepository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make(map[int]*twodo.Todo),
	}
}
