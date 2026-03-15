package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Server holds our store. All handlers are methods on Server.
// This is how all handlers share the same in-memory data.
type Server struct {
	store *TodoStore
}

func NewServer() *Server {
	return &Server{store: &TodoStore{}}
}

// ─── GET /todos ────────────────────────────────────────────
// Returns all todos as a JSON array.
func (s *Server) handleGetTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	todos := s.store.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// ─── POST /todos ───────────────────────────────────────────
// Creates a new todo. Client sends {"title":"..."} in the body.
func (s *Server) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body into a CreateTodoRequest struct
	var req CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Validate: title must not be empty
	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	todo := s.store.Create(req.Title)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 = something was created
	json.NewEncoder(w).Encode(todo)
}

// ─── GET /todos/:id ────────────────────────────────────────
// Returns one specific todo by ID.
func (s *Server) handleGetTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL: "/todos/3" → split → ["", "todos", "3"]
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[2]) // "3" → 3
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, exists := s.store.GetByID(id)
	if !exists {
		http.Error(w, "Todo not found", http.StatusNotFound) // 404
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// ─── PUT /todos/:id ────────────────────────────────────────
// Updates a todo's completed status.
func (s *Server) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if !s.store.Update(id, req.Completed) {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK) // 200 — updated, no body needed
}

// ─── DELETE /todos/:id ─────────────────────────────────────
// Deletes a todo permanently.
func (s *Server) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if !s.store.Delete(id) {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 — deleted, nothing to return
}
