package environment

type Application struct{}

func Init() *Application {
	return &Application{}
}
