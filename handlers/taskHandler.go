package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"practice/server/store"
	"practice/server/utils"
	"practice/server/validators"
	"strconv"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []store.Task
	tasks, err := store.TasksStore.GetTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error getting tasks list"))
		return
	}

	if len(tasks) == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
	}

	data, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error getting tasks list"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	if idString == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("provide a proper id"))
		return
	}

	var task *store.Task
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("provide a proper id"))
		return
	}

	task, err = store.TasksStore.GetTask(id)
	if err != nil {
		if err.Error() == "no Row Found" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("please provide a valid id"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error fetching task"))
		return
	}

	data, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error getting task"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Logger.Debug("error reading request body ", "error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("malformed message"))
		return
	}

	var newTask store.Task
	err = json.Unmarshal(body, &newTask)
	if err != nil {
		utils.Logger.Debug("error unmarshalling request body ", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	validationError := validators.ValidateTask(newTask)

	if validationError != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("invalid request body: %s", validationError.Error())))
		return
	}

	taskId, err := store.TasksStore.AddTask(newTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("{\"id\": \"%d\"}", taskId)))
	return
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	if idString == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("provide a proper id"))
		return
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("provide a proper id"))
		return
	}

	err = store.TasksStore.DeleteTask(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to delete task"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("task deleted successfully"))
	return
}

func GetTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetTagName")
}

func GetTasksDue(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in GetTaskDue")
}
