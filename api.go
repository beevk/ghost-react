package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	port string
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewAPIServer(listenAddress string) *APIServer {
	fmt.Println("Starting API server ...")
	return &APIServer{
		port: listenAddress,
	}
}

func (server *APIServer) RunServer() {
	router := mux.NewRouter()

	router.HandleFunc("/status", MakeHttpHandleFunc(server.handleHealthCheck))
	router.HandleFunc("/react", MakeHttpHandleFunc(server.handleReaction))

	log.Println("Starting API on Port", server.port)
	if err := http.ListenAndServe(server.port, router); err != nil {
		log.Fatal("API Server failed to run!", err)
		return
	}
}

func (server *APIServer) handleHealthCheck(w http.ResponseWriter, r *http.Request) error {
	// Write the response body
	return WriteJSON(w, http.StatusOK, &StatusResponse{
		Status: "OK",
	})
}

func (server *APIServer) handleReaction(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return server.handleGetReaction(w, r)
	case "POST":
		return server.handleSetReaction(w, r)
	default:
		log.Println("method not allowed: ", r.Method)
		return WriteJSON(w, http.StatusMethodNotAllowed, APIError{Error: "Method not allowed"})
	}
}

func (server *APIServer) handleGetReaction(w http.ResponseWriter, r *http.Request) error {

	return WriteJSON(w, http.StatusOK, "OK")
}

func (server *APIServer) handleSetReaction(w http.ResponseWriter, r *http.Request) error {
	article := NewArticle("https://blog.bivek.info.np")
	return WriteJSON(w, http.StatusCreated, article)
}
