package store

import (
	"database/sql"
	"errors"
	"fmt"

	"practice/server/utils"

	_ "github.com/mattn/go-sqlite3"
)

type TaskDBStore struct {
	db *sql.DB
}

func (taskStore TaskDBStore) AddTask(task Task) (int, error) {
	statement := `
		INSERT into tasks (text, due, tags) values (?,?,?);
	`
	result, err := taskStore.db.Exec(statement, task.Text, task.Due, stringifyTags(task.Tags))
	if err != nil {
		utils.Logger.Debug("error storing task in db ", "error", err.Error())
		return -1, err
	}

	insertedId, err := result.LastInsertId()
	return int(insertedId), nil
}

func (taskStore TaskDBStore) GetTask(id int) (*Task, error) {
	statement := `
		SELECT * from tasks where id = ?;
	`

	row := taskStore.db.QueryRow(statement, id)

	var task Task
	var rawTags string
	err := row.Scan(&task.Id, &task.Text, &task.Due, &rawTags)
	tags := slicifyTags(rawTags)
	task.Tags = tags
	switch {
	case errors.Is(err, sql.ErrNoRows):
		utils.Logger.Debug("no rows found in DB for the given ID ", "error", err.Error())
		return nil, errors.New("no Row Found")
	case err != nil:
		utils.Logger.Debug("error retrieving row from db ", "error", err.Error())
		return nil, errors.New("error getting row from DB")
	default:
		return &task, nil
	}
}

func (taskStore TaskDBStore) GetTasks() ([]Task, error) {
	statement := `
		SELECT * from tasks;	
	`

	rows, err := taskStore.db.Query(statement)
	if err != nil {
		utils.Logger.Debug("error reading tasks from db ", "error", err.Error())
		return nil, errors.New("")
	}

	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		var rawTags string
		err := rows.Scan(&task.Id, &task.Text, &task.Due, &rawTags)
		tags := slicifyTags(rawTags)
		task.Tags = tags
		if err != nil {
			utils.Logger.Debug("error scanning row ", "error", err.Error())
			return nil, errors.New("")
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (taskStore TaskDBStore) DeleteTask(id int) error {
	statement := `
		Delete from tasks
		where id = ?; 
	`

	_, err := taskStore.db.Exec(statement, id)
	if err != nil {
		utils.Logger.Debug(fmt.Sprintf("error deleting row with id %d from DB ", id), "error", err.Error())
		return errors.New("unable to delete row")
	}

	return nil
}

func newDBStore() (TaskActions, error) {
	connectionString := "/Users/a.gonugunta/Downloads/tasks.db"
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return nil, err
	}
	return TaskDBStore{
		db: db,
	}, nil
}
