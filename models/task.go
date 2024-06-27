package models

type Task struct {
	Id          int64  `json:"-"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
	StatusId    int64  `json:"status_id" db:"status_id"`
	UserId      int64  `json:"user_id" db:"user_id"`
}

// TaskStatus
// Выполнена
// В работе
// Ожидает
// Отменена

// Приме
type TaskStatus struct {
	Id   int64  `json:"-"`
	Name string `json:"name" db:"name"`
}
