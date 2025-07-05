package store

import "context"

type Store struct {
}

func New() *Store {
	return &Store{}
}

func (s *Store) GetFeeds(ctx context.Context) error {
	return nil
}
