package task

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t Task
	json.NewDecoder(r.Body).Decode(&t)

	t = AddTask(t)

	w.Header()["Location"] = []string{strconv.Itoa(t.ID)}
	w.WriteHeader(http.StatusNoContent)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {

	ts := GetTasks()
	if err := json.NewEncoder(w).Encode(ts); err != nil {
		panic(err)
	}
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var t Task
	t.ID, _ = strconv.Atoi(vars["id"])
	json.NewDecoder(r.Body).Decode(&t)

	UpdateTask(t)

	w.WriteHeader(http.StatusNoContent)
}
