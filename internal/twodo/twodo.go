package twodo

import "time"

// Todo hold todo information
type Todo struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
}

// Repository is the basic interface to implement each Repository
type Repository interface {
	Save(*Todo) error
	ListAll() ([]*Todo, error)
}

// Creator define an interface to create new resources
type Creator interface {
	Run(string, string) error
}

// List define an interface to list resources
type List interface {
	Run() (*[]Todo, error)
}

// Destroyer define an interface to destroy a single resource
type Destroyer interface {
	Run(id int) error
}
