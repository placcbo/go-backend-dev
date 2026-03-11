package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
)

func main() {
    server := NewServer()

    http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/todos" {
            // Exact match: /todos
            switch r.Method {
            case http.MethodGet:
                server.handleGetTodos(w, r)
            case http.MethodPost:
                server.handleCreateTodo(w, r)
            default:
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            }
        } else if strings.HasPrefix(r.URL.Path, "/todos/") {
            // Prefix match: /todos/1, /todos/2, etc.
            switch r.Method {
            case http.MethodGet:
                server.handleGetTodo(w, r)
            case http.MethodPut:
                server.handleUpdateTodo(w, r)
            case http.MethodDelete:
                server.handleDeleteTodo(w, r)
            default:
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            }
        } else {
            http.NotFound(w, r)
        }
    })


    port := ":8080"
    fmt.Printf("Server starting on http://localhost%s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))

}