package models

type UserDto struct {
	Id       int64  `json:"-"`
	FullName string `json:"full_name" binding:"required" db:"full_name"`
	RoleId   int64  `json:"role_id" binding:"required" db:"role_id"`
}
