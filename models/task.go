package models

// Task represents a task entity
// @Description Task model
type Task struct {
	Id          int64  `json:"-"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
	StatusId    int64  `json:"status_id" db:"status_id"`
	UserId      int64  `json:"user_id" db:"user_id"`
}

// TaskStatus represents a task status entity
// @Description TaskStatus model
type TaskStatus struct {
	Id   int64  `json:"-"`
	Name string `json:"name" db:"name"`
}
