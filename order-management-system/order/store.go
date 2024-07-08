package main

import "context"

// OrderStore provides order store
type store struct{}

// NewStore creates a new order store
func NewStore() OrderStore {
	return &store{}
}

// Create creates a new order
func (s *store) Create(ctx context.Context) error {
	return nil
}
