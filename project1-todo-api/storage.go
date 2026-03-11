package main

import "sync"

// TodoStore holds all todos in memory.
// It's safe for concurrent use because every method acquires
// the appropriate mutex lock before touching the map.
type TodoStore struct {
    todos  map[int]Todo  // the data: id → todo

    nextID int           // auto-increment counter

    mu     sync.RWMutex  // protects todos and nextID

}

// NewTodoStore constructs a store with an empty map and ID starting at 1.
func NewTodoStore() *TodoStore {
    return &TodoStore{
        todos:  make(map[int]Todo),
        nextID: 1,
    }
}

// Create adds a new todo and returns it with its assigned ID.
func (s *TodoStore) Create(title string) Todo {
    s.mu.Lock()         // exclusive: nothing else can read OR write

    defer s.mu.Unlock() // always release, even if we panic


    todo := Todo{
        ID:        s.nextID,
        Title:     title,
        Completed: false,
    }
    s.todos[s.nextID] = todo
    s.nextID++
    return todo
}

// GetAll returns a slice containing every todo.
func (s *TodoStore) GetAll() []Todo {
    s.mu.RLock()         // shared: many readers can hold this at once

    defer s.mu.RUnlock()


    todos := make([]Todo, 0, len(s.todos))
    for _, todo := range s.todos {
        todos = append(todos, todo)
    }
    return todos
}

// GetByID returns the todo with the given ID.
// The second return value is false if no todo with that ID exists.
func (s *TodoStore) GetByID(id int) (Todo, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()


    todo, exists := s.todos[id]
    return todo, exists
}

// Update sets the completed status on an existing todo.
// Returns false if the todo doesn't exist.
func (s *TodoStore) Update(id int, completed bool) bool {
    s.mu.Lock()
    defer s.mu.Unlock()


    todo, exists := s.todos[id]
    if !exists {
        return false
    }
    todo.Completed = completed
    s.todos[id] = todo // must reassign — maps hold copies, not pointers

    return true
}

// Delete removes the todo with the given ID.
// Returns false if no todo with that ID exists.
func (s *TodoStore) Delete(id int) bool {
    s.mu.Lock()
    defer s.mu.Unlock()


    _, exists := s.todos[id]
    if !exists {
        return false
    }
    delete(s.todos, id)
    return true
}