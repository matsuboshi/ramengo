package error

import "sync"

type SharedError struct {
	mu  sync.Mutex
	err error
}

func (e *SharedError) Update(err error) {
	e.mu.Lock()
	e.err = err
	e.mu.Unlock()
}

func (e *SharedError) Read() error {
	e.mu.Lock()
	err := e.err
	e.mu.Unlock()
	return err
}
