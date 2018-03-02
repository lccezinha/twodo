package environment

import (
	"database/sql"

	repository "github.com/lccezinha/twodo/internal/repository/todo"
	"github.com/lccezinha/twodo/internal/twodo"
	"github.com/lccezinha/twodo/internal/twodo/todo"

	_ "github.com/lib/pq"
)

// Application will hold all services to run up the app
type Application struct {
	CreateService twodo.Creator
	// DestroyService twodo.Destroyer
	// ListService    twodo.List
	// UpdateService  twodo.Updater
}

// EnvVars will hold all env vars
type EnvVars map[string]string

// Init will return a instance of Application with all apps config
func Init() *Application {
	envVars := loadEnvVars()
	dataBase := loadDatabase(envVars)
	repository := loadRepository(dataBase)
	createService := loadCreateService(repository)
	// listService := loadListService(repository)
	// destroyService := loadDestroyService(repository)
	// updateService := loadUpdateService(repository)

	return &Application{
		CreateService: createService,
		// ListService:    listService,
		// DestroyService: destroyService,
		// UpdateService:  updateService,
	}
}

func loadEnvVars() EnvVars {
	envVars := EnvVars{}
	envVars["DATABASE_URL"] = "postgres://postgres@localhost/twodo"

	return envVars
}

func loadDatabase(e EnvVars) *sql.DB {
	db, err := sql.Open("postgres", e["DATABASE_URL"])
	if err != nil {
		panic(err.Error())
	}

	return db
}

func loadRepository(db *sql.DB) twodo.Repository {
	return repository.NewMemoryRepository()

	// return repository.NewPostgresRepository(db)
}

func loadCreateService(r twodo.Repository) twodo.Creator {
	return todo.NewCreateService(r)
}

// func loadListService(r twodo.Repository) twodo.List {
// 	return todo.NewListService(r)
// }
//
// func loadDestroyService(r twodo.Repository) twodo.Destroyer {
// 	return todo.NewDestroyService(r)
// }
//
// func loadUpdateService(r twodo.Repository) twodo.Updater {
// 	return todo.NewUpdateService(r)
// }
