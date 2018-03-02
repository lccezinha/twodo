package todo

import (
	"database/sql"

	"github.com/lccezinha/twodo/internal/twodo"
)

const (
	insertStmt    = "INSERT INTO todos (title, description, created_at, done) VALUES ($1, $2, $3, $4) RETURNING id;"
	selectAllStmt = "SELECT id, title, description, created_at, done FROM todos ORDER BY id DESC;"
	destroyStmt   = "DELETE FROM todos WHERE id = $1;"
	updateStmt    = "UPDATE todos SET done = $1 WHERE id = $2;"
)

// PostgresRepository holds repository that deal with postgres
type PostgresRepository struct {
	DB *sql.DB
}

// Save will save the user in postgres repository
func (p *PostgresRepository) Save(t twodo.Todo) (twodo.Todo, error) {
	err := p.DB.QueryRow(
		insertStmt, t.Title, t.Description, t.CreatedAt, t.Done,
	).Scan(&t.ID)

	return t, err
}

// ListAll will list all resources
func (p *PostgresRepository) ListAll() ([]twodo.Todo, error) {
	var todos []twodo.Todo

	rows, err := p.DB.Query(selectAllStmt)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		todo := twodo.Todo{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.Done); err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// Destroy will destroy a single resource
func (p *PostgresRepository) Destroy(id int) error {
	stmt, err := p.DB.Prepare(destroyStmt)

	if err != nil {
		return err
	}

	stmt.Exec(id)

	return nil
}

// Update will update the done status
func (p *PostgresRepository) Update(id int, done bool) (twodo.Todo, error) {
	todo := twodo.Todo{}

	err := p.DB.QueryRow(updateStmt, done, id).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.Done)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

// NewPostgresRepository will return a new instance of PostgresRepository
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		DB: db,
	}
}
