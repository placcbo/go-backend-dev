package main

import "sync"

type TodoStore struct {
	todos  map[int]Todo
	NextID int
	mu     sync.RWMutex
}

func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos:  make(map[int]Todo),
		NextID: 1,
	}
}

// Create adds a new todo and returns it with its assigned ID.
func (s *TodoStore) Create(title string) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	todo := Todo{
		ID:        s.NextID,
		Title:     title,
		Completed: false,
	}
	s.todos[s.NextID] = todo
	s.NextID++
	return todo
}
