package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	// DB     *gorm.DB
}

func (a *App) Initialize() {

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/", a.Home).Methods("GET")
	a.Router.HandleFunc("/todos", a.Todos).Methods("GET")
	a.Router.HandleFunc("/todos", a.TodoCreate).Methods("POST")
	a.Router.HandleFunc("/todos/{id}", a.Todo).Methods("GET")
}

func (a *App) Run(addr string) {

	fmt.Println("listening at 8080")
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		log.Fatal(err)
	}

}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")

}
func (a *App) Todos(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Kaan"},
		Todo{Name: "Atilla"},
	}
	respondWithJSON(w, http.StatusOK, todos)

}
func (a *App) TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondWithError(w, http.StatusOK, " :((( ")

	}
	// respondWithJSON(w, http.StatusOK, todos)

}
func (a *App) Todo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "todo page %s", vars["id"])
}

// ---------------------------------------------------------------------------

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})

}
