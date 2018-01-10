package twodo

import "time"

// Todo hold todo information
type Todo struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
}

// Todos is a simple slice of Todo
type Todos []Todo

// Repository is the basic interface to implement each Repository
type Repository interface {
	Save(*Todo) error
	ListAll() (*Todos, error)
}

// Creator define an interface to create new resources
type Creator interface {
	Run(string, string) error
}

// Destroyer define an interface to destroy a single resource
type Destroyer interface {
	Run(id int) error
}
