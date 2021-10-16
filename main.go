package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/configs"
	"gitlab.cs.ui.ac.id/ppl-fasilkom-ui/galleryppl/gallery-api/controllers"
)

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", configs.AppConfig.Port),
		Handler:        controllers.InitializeRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("[INIT] failed starting server at %s. %v", s.Addr, err)
	}
}
