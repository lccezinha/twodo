package env

import (
	"database/sql"
	"log"
	"time"

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
	v["DATABASE_URL"] = "postgres://postgres@localhost/twodo_app"

	return v
}

func loadDatabase(e vars) *sql.DB {
	db, err := sql.Open("postgres", e["DATABASE_URL"])
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(i) * time.Second)

		if err = db.Ping(); err == nil {
			break
		}
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func loadRepository(db *sql.DB) twodo.Repository {
	// return todo.NewMemoryRepository()

	return todo.NewPostgresRepository(db)
}
