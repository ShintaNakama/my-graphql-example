package entity

type NewTodo struct {
	ID     string
	Text   string
	UserID string
}

type Todo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}

type User struct {
	ID   string
	Name string
}
