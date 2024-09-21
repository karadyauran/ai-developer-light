package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "io/ioutil"
    "strconv"
    "../models"
    "../utils"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/tasks", GetTasks).Methods("GET")
    r.HandleFunc("/tasks", CreateTask).Methods("POST")
    r.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
    r.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
    tasks := models.GetAllTasks()
    json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &task)
    models.AddTask(task)
    json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    var task models.Task
    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &task)
    models.UpdateTask(id, task)
    json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])
    models.DeleteTask(id)
    utils.RespondWithJSON(w, http.StatusNoContent, nil)
}