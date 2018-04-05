package twodo

// Todo hold todo information
type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// CreatedAt   time.Time `json:"created_at"`
	Done bool `json:"done"`
}

// ValidationError hold the error information
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

// Repository is the basic interface to implement each Repository
type Repository interface {
	Save(Todo) (Todo, error)
	// ListAll() ([]Todo, error)
	// Destroy(id int) error
	// Update(id int, done bool) (Todo, error)
}

// Validator is the interface to create Todos validations
type Validator interface {
	Validate(todo Todo) []ValidationError
}

// Presenter is the interface to present Todos
type Presenter interface {
	PresentCreatedTodo(todo Todo)
	PresentErrors(errs []ValidationError)
}

// CreateUseCase is the interface to create a use case that create a todo
type CreateUseCase interface {
	Run(title string, description string, presenter Presenter) error
}
