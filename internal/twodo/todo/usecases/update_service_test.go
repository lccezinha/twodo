package todo

//
// import (
// 	"testing"
//
// 	"github.com/lccezinha/twodo/internal/repository/todo"
// )
//
// func TestUpdate(t *testing.T) {
// 	t.Run("when todo has done attribute as false, update to true", func(t *testing.T) {
// 		repository := todo.NewMemoryRepository()
// 		service := NewUpdateService(repository)
// 		serviceCreate := NewCreateService(repository)
// 		serviceList := NewListService(repository)
//
// 		serviceCreate.Run("Title", "Description")
// 		todos, _ := serviceList.Run()
//
// 		if todos[0].Done != false {
// 			t.Errorf("Default todo is not false")
// 		}
//
// 		service.Run(todos[0].ID, true)
//
// 		todos, _ = serviceList.Run()
//
// 		if todos[0].Done != true {
// 			t.Errorf("Todo was not updated")
// 		}
//
// 		service.Run(todos[0].ID, false)
//
// 		todos, _ = serviceList.Run()
//
// 		if todos[0].Done != false {
// 			t.Errorf("Todo was not updated")
// 		}
// 	})
// }
