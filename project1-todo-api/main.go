package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
)

func main() {
    server := NewServer()

    // Register ONE handler for all /todos routes.
    // Go's HandleFunc does prefix matching, so "/todos" catches
    // "/todos", "/todos/1", "/todos/99" — all of them.
    http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {

        if r.URL.Path == "/todos" {
            // Exact match: /todos (no ID)
            switch r.Method {
            case http.MethodGet:
                server.handleGetTodos(w, r)
            case http.MethodPost:
                server.handleCreateTodo(w, r)
            default:
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            }

        } else if strings.HasPrefix(r.URL.Path, "/todos/") {
            // Prefix match: /todos/1, /todos/99, etc. (has an ID)
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
    fmt.Printf("✓ Server running at http://localhost%s\n", port)
    fmt.Println("  Press Ctrl+C to stop")

    // ListenAndServe blocks forever — it runs until the server stops.
    // log.Fatal logs the error and exits if something goes wrong at startup.
    log.Fatal(http.ListenAndServe(port, nil))

}