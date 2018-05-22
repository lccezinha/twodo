package todo

import (
	"strconv"
	"testing"

	"github.com/lccezinha/twodo/internal/twodo"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSave(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	todoID := 1
	todo := twodo.Todo{
		Description: "Description",
		Done:        false,
	}

	rows := sqlmock.NewRows([]string{"id"}).FromCSVString(strconv.Itoa(todoID))

	mock.ExpectQuery("INSERT INTO todos (.+) VALUES (.+) RETURNING id").
		WithArgs(todo.Description, todo.Done).
		WillReturnRows(rows)

	repository := NewPostgresRepository(db)
	createdTodo, err := repository.Save(todo)

	if createdTodo.Description != "Description" {
		t.Errorf("Todo.Description does not match. Expect %s, got %s", todo.Description, createdTodo.Description)
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

	rows := sqlmock.NewRows([]string{"id", "description", "done"}).
		AddRow(2, "Description #2", false).
		AddRow(1, "Description #1", false)

	mock.ExpectQuery("SELECT (.+) FROM todos ORDER BY id DESC").WillReturnRows(rows)

	repository := NewPostgresRepository(db)
	todos, err := repository.ListAll()

	if err != nil {
		t.Error("Error: ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expections: %s", err)
	}

	if todos[0].Description != "Description #2" {
		t.Error("Does not fetch the most recent todo first")
	}

	if todos[1].Description != "Description #1" {
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

		row := sqlmock.NewRows([]string{"id", "description", "done"}).
			AddRow(1, "Description #1", true)

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

		row := sqlmock.NewRows([]string{"id", "description", "done"}).
			AddRow(1, "Description #1", false)

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
