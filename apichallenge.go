package main

import (
	"encoding/json"
	_ "fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/1set/todotxt"
	"github.com/julienschmidt/httprouter"
)

// Use this function to output the list of todos.  This function should accept
// query params that allow parameterization of the search
func ListTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos, e := todotxt.LoadFromPath("todo.txt")
	if e != nil {
		//handle error
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}

// Use this function to get a specific todo
func GetTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	tdid := ps.ByName("id")
	todo_id, e := strconv.Atoi(tdid)
	if e != nil {
		//If the parameter is not parseable as an integer, there will be an error here.  You probably don't need to worry about this.
	}

	todos, e := todotxt.LoadFromPath("todo.txt")
	if e != nil {
		//If it cannot find the file, there will be an error here. You probably don't need to worry about this.
	}

	task, e := todos.GetTask(todo_id)
	if e != nil {
		//If the task is not found, there will be an error here
		//fmt.Println(e.Error())
		//handle error
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}

func UpdateTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	tdid := ps.ByName("id")
	todo_id, e := strconv.Atoi(tdid)
	if e != nil {
		//If the parameter is not parseable as an integer, there will be an error here.  You probably don't need to worry about this.
	}

	tasks, e := todotxt.LoadFromPath("todo.txt")
	if e != nil {
		//If it cannot find the file, there will be an error here. You probably don't need to worry about this.
	}

	task, e := tasks.GetTask(todo_id)
	if e != nil {
		//If the task is not found, there will be an error here
		//fmt.Println(e.Error())
		//handle error
	}

	decoder := json.NewDecoder(r.Body)
	e = decoder.Decode(&task)

	todotxt.WriteToPath(&tasks, "todo.txt")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}

func CreateTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	todos, _ := todotxt.LoadFromPath("todo.txt")
	task, _ := todos.GetTask(2)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)

}

func main() {

	router := httprouter.New()
	router.GET("/todos", ListTodos)
	router.GET("/todo/:id", GetTodo)
	router.PUT("/todo/:id", UpdateTodo)
	router.POST("/todo", CreateTodo)

	router.NotFound = http.FileServer(http.Dir("./static"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
