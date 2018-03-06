package todo

import (
	"strconv"
	"testing"
	"time"

	"github.com/lccezinha/twodo/internal/twodo"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSave(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	todoID := 1
	createdAt, _ := time.Parse("2006-01-02", "2013-02-03")
	todo := twodo.Todo{
		Title:       "Title",
		Description: "Description",
		CreatedAt:   createdAt,
		Done:        false,
	}

	rows := sqlmock.NewRows([]string{"id"}).FromCSVString(strconv.Itoa(todoID))

	mock.ExpectQuery("INSERT INTO todos (.+) VALUES (.+) RETURNING id").
		WithArgs(todo.Title, todo.Description, todo.CreatedAt, todo.Done).
		WillReturnRows(rows)

	repository := NewPostgresRepository(db)
	createdTodo, err := repository.Save(todo)

	if createdTodo.Title != "Title" {
		t.Errorf("Todo.title does not match. Expect %s, got %s", todo.Title, createdTodo.Title)
	}

	if err != nil {
		t.Error("Error: ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expections: %s", err)
	}

	if createdTodo.ID != todoID {
		t.Error("Does not assign value to todoID")
	}

	if todo.Done != false {
		t.Error("Todo must be false by default")
	}
}

func TestListAll(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	createdAtTodo, _ := time.Parse("2006-01-02", "2013-02-03")
	createdAtOtherTodo, _ := time.Parse("2006-01-02", "2013-02-04")

	rows := sqlmock.NewRows([]string{"id", "title", "description", "created_at", "done"}).
		AddRow(2, "Title #2", "Description #2", createdAtOtherTodo, false).
		AddRow(1, "Title #1", "Description #1", createdAtTodo, false)

	mock.ExpectQuery("SELECT (.+) FROM todos ORDER BY id DESC").WillReturnRows(rows)

	repository := NewPostgresRepository(db)
	todos, err := repository.ListAll()

	if err != nil {
		t.Error("Error: ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expections: %s", err)
	}

	if todos[0].Title != "Title #2" {
		t.Error("Does not fetch the most recent todo first")
	}

	if todos[1].Title != "Title #1" {
		t.Error("Does not fetch the most old todo last")
	}
}

func TestDestroy(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	todoID := 1

	mock.ExpectPrepare("DELETE FROM todos WHERE id = (.+)")

	repository := NewPostgresRepository(db)
	err := repository.Destroy(todoID)

	if err != nil {
		t.Error("Error: ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expections: %s", err)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("Update todo.Done from false to true", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer db.Close()

		createdAtTodo, _ := time.Parse("2006-01-02", "2013-02-03")
		row := sqlmock.NewRows([]string{"id", "title", "description", "created_at", "done"}).
			AddRow(1, "Title #1", "Description #1", createdAtTodo, true)

		mock.ExpectQuery("UPDATE todos SET done = (.+) WHERE id = (.+);").
			WithArgs(true, 1).
			WillReturnRows(row)

		repository := NewPostgresRepository(db)
		todo, err := repository.Update(1, true)

		if err != nil {
			t.Error("Error: ", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expections: %s", err)
		}

		if todo.Done != true {
			t.Errorf("Todo was not update correctly. todo.Done eq %v, expected %v", todo.Done, true)
		}
	})

	t.Run("Update todo.Done from true to false", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer db.Close()

		createdAtTodo, _ := time.Parse("2006-01-02", "2013-02-03")
		row := sqlmock.NewRows([]string{"id", "title", "description", "created_at", "done"}).
			AddRow(1, "Title #1", "Description #1", createdAtTodo, false)

		mock.ExpectQuery("UPDATE todos SET done = (.+) WHERE id = (.+);").
			WithArgs(false, 1).
			WillReturnRows(row)

		repository := NewPostgresRepository(db)
		todo, err := repository.Update(1, false)

		if err != nil {
			t.Error("Error: ", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expections: %s", err)
		}

		if todo.Done != false {
			t.Errorf("Todo was not update correctly. todo.Done eq %v, expected %v", todo.Done, false)
		}
	})
}
