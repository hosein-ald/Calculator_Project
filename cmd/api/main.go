package main

import (
	"calculatorGo/internal/handlers"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	db, err := sql.Open("sqlite3", "./log.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.SetReportCaller(true)

	r := chi.NewRouter()

	// Register handlers
	r.Post("/subtract", handlers.HandleSubtract)
	r.Post("/add", handlers.HandleAdd)
	r.Post("/multiply", handlers.HandleMultiply)
	r.Post("/divide", handlers.HandleDivide)

	fmt.Println("Server is running on http://localhost:9000")

	if err := http.ListenAndServe("localhost:9000", r); err != nil {
		log.Fatal(err)
	}
}
