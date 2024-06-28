package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"ims-staff-api/models"
	"reflect"
	"strings"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (t *TaskPostgres) GetTasksList() ([]models.Task, error) {
	var tasks []models.Task

	query := fmt.Sprintf(`SELECT id, title, description, status_id, COALESCE(user_id, 0) as user_id FROM %s`, taskTable)

	err := t.db.Select(&tasks, query)

	return tasks, err
}

func (t *TaskPostgres) GetTasksByStatus(statusName string) ([]models.Task, error) {
	var task []models.Task

	query := fmt.Sprintf("SELECT t.id, t.title, t.description, t.status_id, COALESCE(t.user_id, 0) AS user_id FROM %s t JOIN %s ts ON t.status_id = ts.id WHERE ts.name = $1", taskTable, taskStatusTable)
	err := t.db.Select(&task, query, statusName)

	return task, err
}

func (t *TaskPostgres) GetTaskByID(id int64) (models.Task, error) {
	var task models.Task

	query := fmt.Sprintf("SELECT id, title, description, status_id, COALESCE(user_id, 0) as user_id FROM %s WHERE id = $1", taskTable)
	err := t.db.Get(&task, query, id)

	return task, err
}

func (t *TaskPostgres) CreateTask(task models.Task, statusId int64) (int64, error) {
	logrus.Println(task, statusId)
	var id int64
	createQuery := fmt.Sprintf(`INSERT INTO %s (title, description, status_id) VALUES ($1, $2, $3) RETURNING id`, taskTable)

	row := t.db.QueryRow(createQuery, task.Title, task.Description, statusId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TaskPostgres) UpdateTask(taskId int64, input models.UpdateTask) error {
	placeholders := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	prodValue := reflect.ValueOf(input)

	for i := 0; i < prodValue.NumField(); i++ {
		valueField := prodValue.Field(i)
		typeField := prodValue.Type().Field(i)
		if valueField.Kind() == reflect.Ptr && !valueField.IsNil() {
			dbName := typeField.Tag.Get("db")
			if dbName == "" {
				dbName = typeField.Name
			}

			placeholders = append(placeholders, fmt.Sprintf("%s=$%d", dbName, argId))
			args = append(args, valueField.Elem().Interface())
			argId++
		}
	}

	if argId == 1 {
		return nil
	}

	updateString := strings.Join(placeholders, ", ")
	updateQuery := fmt.Sprintf(`UPDATE %s SET %s WHERE id=$%d`, taskTable, updateString, argId)
	args = append(args, taskId)

	_, err := t.db.Exec(updateQuery, args...)
	return err
}

func (t *TaskPostgres) DeleteTask() error {
	return nil
}
