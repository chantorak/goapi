package main

import (
	"myapi/internals/handlers"

	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API sercice...")
	fmt.Println(`GO API`)

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		fmt.Println("Hahaha error")
		log.Error(err)
	}
}
