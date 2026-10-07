package service

import "context"

func (s *DataStore) CreateTask(ctx context.Context, task string, description string) error {
	return s.Create(ctx, task, description)
}

func (s *DataStore) UpdateTask(ctx context.Context, id int, task string, description string, done bool) error {
	return s.Update(ctx, id, task, description, done)
}
