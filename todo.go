package main

// Todo ...
type Todo struct {
	ID     int64
	Name   string
	Status bool
}

// Todos ...
type Todos struct {
	Items []Todo
	Index int64
}

func (ts *Todos) add(t *Todo) {
	t.ID = ts.Index + 1
	ts.Items = append(ts.Items, *t)
	ts.Index++
}

func (ts *Todos) remove(todoID int64) {
	todoIndex := find(func(todo Todo) bool { return todo.ID == todoID }, ts.Items)
	if todoIndex > -1 {
		ts.Items = append(ts.Items[:todoIndex], ts.Items[todoIndex+1:]...)
	}
}

func (ts *Todos) markComplete(todoID int64) {
	todoIndex := find(func(todo Todo) bool { return todo.ID == todoID }, ts.Items)
	if todoIndex > -1 {
		ts.Items[todoIndex].Status = true
	}
}

func find(fn func(t Todo) bool, todos []Todo) int {
	for i, v := range todos {
		if todo := fn(v); todo == true {
			return i
		}
	}

	return -1
}
