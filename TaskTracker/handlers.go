package handlers

import (
    "encoding/json"
    "net/http"
    "project/models"
    "project/utils"
    "github.com/gorilla/mux"
    "strconv"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
    tasks := models.GetAllTasks()
    json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)
    models.AddTask(task)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)
    updatedTask, err := models.UpdateTask(id, task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(updatedTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    err := models.DeleteTask(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}