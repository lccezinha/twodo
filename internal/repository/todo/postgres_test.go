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
	todo := &twodo.Todo{
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
	err := repository.Save(todo)

	if err != nil {
		t.Error("Error: ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expections: %s", err)
	}

	if todo.ID != todoID {
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
	// db, mock, _ := sqlmock.New()
	// defer db.Close()
	//
	// // createdAtTodo, _ := time.Parse("2006-01-02", "2013-02-03")
	// // row := sqlmock.NewRows([]string{"id", "title", "description", "created_at", "done"}).
	// // 	AddRow(1, "Title #1", "Description #1", createdAtTodo, true)
	//
	// mock.ExpectQuery("SELECT done FROM todos WHERE id = (.+) RETURNING done").
	// 	WithArgs(1)
	//
	// // mock.ExpectQuery("UPDATE FROM todos SET done = (.+) WHERE id = (.+)").
	// // 	WithArgs(1).
	// // 	WillReturnRows(row)
	//
	// repository := NewPostgresRepository(db)
	// _, err := repository.Update(1)
	//
	// if err != nil {
	// 	t.Error("Error: ", err)
	// }
	//
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("There were unfulfilled expections: %s", err)
	// }
}

func TestFindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	createdAtTodo, _ := time.Parse("2006-01-02", "2013-02-03")
	rows := sqlmock.NewRows([]string{"id", "title", "description", "created_at", "done"}).
		AddRow(1, "Title #1", "Description #1", createdAtTodo, false)

	mock.ExpectQuery("SELECT id, title, description, done, created_at FROM todos WHERE id = (.+)").
		WithArgs(1).
		WillReturnRows(rows)

	repository := NewPostgresRepository(db)
	todo, err := repository.FindByID(1)

	if err != nil {
		t.Error("Error: ", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expections: %s", err)
	}

	if todo.Title != "Title #1" {
		t.Error("Todo's title does not match")
	}
}
