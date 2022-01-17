package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"lastimplementation.com/pkg/db"
	projectsAPI "lastimplementation.com/pkg/services/projects/transport"
)

const (
	dbhost     = "localhost"
	dbport     = 5432
	dbuser     = "postgres"
	dbpassword = "pass123" // temporary
	dbname     = "projects"
	dbreset    = false

	srvhost = "0.0.0.0"
	srvport = 8081
)

func main() {
	ctx := context.Background()
	l := log.New(os.Stdout, "api ", log.LstdFlags)

	// Get db connection.
	db, err := db.GetConnection(&db.DBConfig{
		Host:     dbhost,
		Port:     dbport,
		User:     dbuser,
		Password: dbpassword,
		DBname:   dbname,
	})
	if err != nil {
		l.Printf("unable to establish a database connection: %v", err)
		os.Exit(1)
		return
	}
	defer db.Close()

	// Setup server.
	r := mux.NewRouter()
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", srvhost, srvport),
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	l.Printf("Running server on port %d\n", srvport)

	// Setup services
	projectsAPI.Activate(ctx, r, db, dbreset)

	go func() {
		// Initiate the server listening.
		err := srv.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// Wait for a signal that ends the execution.
	signc := make(chan os.Signal, 1)
	signal.Notify(signc, os.Interrupt)

	sign := <-signc
	l.Println("Receive terminate, grateful shutdown", sign)

	// Set a timeout to end the server gracefully.
	tc, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// End gracefully.
	srv.Shutdown(tc)
	os.Exit(0)
}
