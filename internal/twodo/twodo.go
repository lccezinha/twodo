package twodo

import (
	"time"
)

// Todo hold todo information
type Todo struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	Done        bool
}

// ValidationError hold the error information
type ValidationError struct {
	Field   string
	Message string
	Type    string
}

// Repository is the basic interface to implement each Repository
type Repository interface {
	Save(Todo) (Todo, error)
	ListAll() ([]Todo, error)
	Destroy(id int) error
	Update(id int, done bool) (Todo, error)
}

// Validator is the interface to create Todos validations
type Validator interface {
	Validate(todo Todo) []ValidationError
}

// Creator define an interface to create new resources
type Creator interface {
	Run(string, string) error
}

//
// // List define an interface to list resources
// type List interface {
// 	Run() ([]*Todo, error)
// }
//
// // Destroyer define an interface to destroy a single resource
// type Destroyer interface {
// 	Run(id int) error
// }
//
// // Updater define an interface to update a single resource
// type Updater interface {
// 	Run(id int, done bool) (*Todo, error)
// }
