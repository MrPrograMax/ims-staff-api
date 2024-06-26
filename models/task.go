package models

type Task struct {
	Id          int64
	Title       string
	Description string
	StatusId    int64
	UserId      int64
}

// TaskStatus
// Выполнена
// В работе
// Ожидает
// Отменена

// Приме
type TaskStatus struct {
	Id   int64
	Name string
}
