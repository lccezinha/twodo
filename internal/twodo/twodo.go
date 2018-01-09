// package tuiter
//
// import "time"
//
// // User hold user information
// type User struct {
// 	ID        int
// 	Username  string
// 	Password  string
// 	CreatedAt time.Time
// }
//
// // Repository is the basic user repository interface
// type Repository interface {
// 	Save(*User) error
// 	FindByID(id int) (*User, error)
// 	FindByUsername(username string) (*User, error)
// }
//
// // Creator defined the interface for user service
// type Creator interface {
// 	Run(string, string) error
// }
