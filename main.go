package main

import (
	"aye-robot/api"
	"aye-robot/services"
	"aye-robot/types"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	listenAddr := flag.String("listenAddr", ":9090", "The Address the Server should listen on.")
	flag.Parse()

	ghToken := os.Getenv("GH_TOKEN")
	aiToken := os.Getenv("AI_TOKEN")

	ctx := context.Background()
	ai := services.NewChatGPTClient(ctx, aiToken)
	git := services.NewGithubClientWithAuth(ctx, ghToken)
	logger := log.New(os.Stdout, "aye-robot ", log.LstdFlags)
	validation := types.NewPullRequestValidation()

	handler := api.NewPullRequestReviewHandler(ai, git, logger, validation)
	sm := mux.NewRouter()

	// handlers for API
	postPullRequest := sm.Methods(http.MethodPost).Subrouter()
	postPullRequest.HandleFunc("/", handler.ReviewPR)
	postPullRequest.Use(handler.MiddlewareValidatePullRequestReview)

	// create a new server
	server := http.Server{
		Addr:         *listenAddr,       // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     logger,            // set the logger for the server
		ReadTimeout:  10 * time.Second,  // max time to read request from the client
		WriteTimeout: 120 * time.Second, // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logger.Println("Starting server on port: ", listenAddr)

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("%s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	logger.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
	defer cancel()
}
