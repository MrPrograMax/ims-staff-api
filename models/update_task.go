package models

import (
	"fmt"
	"reflect"
)

type UpdateTask struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	StatusId    *int64  `json:"status_id" db:"status_id"`
	UserId      *int64  `json:"user_id" db:"user_id"`
}

func Verify(item interface{}) error {
	prodType := reflect.ValueOf(item).Elem()
	for i := 0; i < prodType.NumField(); i++ {
		valueField := prodType.Field(i)
		if valueField.Kind() == reflect.Interface {
			continue
		}

		if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
			return nil
		}
	}
	return fmt.Errorf("updating entity is empty")
}
