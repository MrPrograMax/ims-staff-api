package models

// UserDto represents a user data transfer object
// @Description Worker model
type UserDto struct {
	Id       int64  `json:"-"`
	FullName string `json:"full_name" binding:"required" db:"full_name"`
	RoleId   int64  `json:"role_id" binding:"required" db:"role_id"`
}
