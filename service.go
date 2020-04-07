package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	datasourceName := os.Getenv("SQL_DATASOURCE_NAME")
	_, err := sql.Open("cloudsqlpostgres", datasourceName)
	log.Print(datasourceName)
	log.Print(err)
	http.ListenAndServe(":3000", r)
	log.Print(postgres.Driver{})
}
