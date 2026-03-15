package main

import "sync"

type TodoStore struct {
	todos  map[int]Todo
	NextId int
	mu     sync.RWMutex
}

// NewTodoStore creates an empty store, ready to use.
func NewTodoStore() TodoStore {
	return TodoStore{
		todos:  make(map[int]Todo),
		NextId: 1,
	}
}

// Create adds a new todo and returns it with its assigned ID.
func (s *TodoStore) Create(title string) Todo {
	s.mu.Lock() //exclusive write lock
	defer s.mu.Unlock()
	todo := Todo{
		ID:        s.NextId,
		Title:     title,
		Completed: true,
	}

	s.todos[s.NextId] = todo // save it
	s.NextId++
	return todo
}

// GetAll returns every todo as a slice.

func (s *TodoStore) GetAll() []Todo {
	s.mu.Lock()
	s.mu.RLocker().Unlock()
	todos := make([]Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (s *TodoStore) GetByID(id int) (Todo, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, exists := s.todos[id]
	return todo, exists
}

func (s *TodoStore) Update(id int, completed bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, exists := s.todos[id]
	if !exists {
		return false
	}
	todo.Completed = completed
	s.todos[id] = todo // ← CRITICAL: must write the copy back! (explained below)

	return true
}
func (s *TodoStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.todos[id]
	if !exists {
		return false
	}
	delete(s.todos, id) // built-in function to remove a map key
	return true
}
