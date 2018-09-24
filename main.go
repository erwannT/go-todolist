package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/erwannT/go-todolist/connection"
	"github.com/erwannT/go-todolist/task"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Status struct {
	Status string `json:"status"`
}

func main() {

	fmt.Println("Starting application")

	viper.SetDefault("PORT", "8080")
	viper.SetDefault("MYSQL_HOST", "localhost")
	viper.SetDefault("MYSQL_PORT", "3306")
	viper.AutomaticEnv()

	connection.InitDB()

	initGracefullyShutdown()

	var serverPort = viper.GetString("PORT")

	// http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodGet {
	// 		response, _ := json.Marshal(Status{"OK"})
	// 		fmt.Fprintf(w, "%s", response)
	// 	} else {
	// 		w.WriteHeader(http.StatusMethodNotAllowed)
	// 	}
	// })

	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(Status{"OK"}); err != nil {
			panic(err)
		}
	}

	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/tasks", task.PostTaskHandler).Methods(http.MethodPost)
	r.HandleFunc("/tasks", task.GetTasksHandler).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", task.UpdateTaskHandler).Methods(http.MethodPut)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}

func initGracefullyShutdown() {
	var gracefulStop = make(chan os.Signal)

	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for closing DB")
		connection.Close()
		os.Exit(0)
	}()
}
