package env

import (
	"database/sql"

	"github.com/lccezinha/twodo/cmd/repository/todo"
	"github.com/lccezinha/twodo/internal/twodo"

	_ "github.com/lib/pq"
)

// Application will hold all services to run up the app
type Application struct {
	Repository twodo.Repository
}

// EnvVars will hold all env vars
type vars map[string]string

// Init will return a instance of Application with all apps config
func Init() *Application {
	envVars := loadEnvVars()
	database := loadDatabase(envVars)
	repository := loadRepository(database)

	return &Application{Repository: repository}
}

func loadEnvVars() vars {
	v := vars{}
	v["DATABASE_URL"] = "postgres://postgres@localhost/twodo"

	return v
}

func loadDatabase(e vars) *sql.DB {
	db, err := sql.Open("postgres", e["DATABASE_URL"])
	if err != nil {
		panic(err.Error())
	}

	return db
}

func loadRepository(db *sql.DB) twodo.Repository {
	// return todo.NewMemoryRepository()

	return todo.NewPostgresRepository(db)
}
