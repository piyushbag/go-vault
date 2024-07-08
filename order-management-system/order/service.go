package main

import "context"

// OrderService provides order service
type service struct {
	store OrderStore
}

// NewService creates a new order service
func NewService(store OrderStore) OrderService {
	return &service{store}
}

// CreateOrder creates a new order
func (s *service) CreateOrder(ctx context.Context) error {
	return s.store.Create(ctx)
}
