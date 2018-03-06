package todo

// import (
// 	"testing"
//
// 	"github.com/lccezinha/twodo/internal/repository/todo"
// )
//
// func TestCreateService(t *testing.T) {
// 	t.Run("Given invalid args return status 400 and invalid fields", func(t *testing.T) {
// 		repository := todo.NewMemoryRepository()
// 		service := NewCreateService(repository)
//
// 		err := service.Run("", "Description")
// 		expectedErrs := []map[string]string{
// 			"Field":     "title",
// 			"Message":   "Can't be blank",
// 			"ErrorType": "REQUIRED",
// 		}
// 	})
//
// t.Run("when all fields are filled must save", func(t *testing.T) {
// 	repository := todo.NewMemoryRepository()
// 	service := NewCreateService(repository)
//
// 	title := "Some title"
// 	description := "Some Description"
//
// 	err := service.Run(title, description)
//
// 	if err != nil {
// 		t.Errorf("Error: %s", err.Error())
// 	}
//
// 	err = service.Run("Other Title", "")
//
// 	if err != nil {
// 		t.Errorf("Error: %s", err.Error())
// 	}
// })
// }
