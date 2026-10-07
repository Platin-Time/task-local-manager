package service

import (
	"context"
	"prod/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type DataStore struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *DataStore {
	return &DataStore{db: db}
}

func (s *DataStore) Create(ctx context.Context, task string, description string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx, err := s.db.BeginTxx(ctx, nil)

	if err != nil {
		return err
	}
	defer tx.Rollback()

	qwery := `INSERT INTO tasks (task, description, done) VALUES ($1, $2, $3) RETURNING id`

	var lastInsertID int64

	err = tx.QueryRowContext(ctx, qwery, task, description, false).Scan(&lastInsertID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *DataStore) GetAll(ctx context.Context) ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx, err := s.db.BeginTxx(ctx, nil)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var tasks []models.Task
	qwery := `SELECT * FROM tasks`
	err = s.db.SelectContext(ctx, &tasks, qwery)
	if err != nil {
		return nil, err
	}

	return tasks, tx.Commit()
}

func (s *DataStore) GetByID(ctx context.Context, id int) (*models.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx, err := s.db.BeginTxx(ctx, nil)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var task models.Task
	qwery := `SELECT * FROM tasks WHERE id = $1`
	err = s.db.GetContext(ctx, &task, qwery, id)
	if err != nil {
		return nil, err
	}

	return &task, tx.Commit()
}

func (s *DataStore) Update(ctx context.Context, id int, task string, description string, done bool) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx, err := s.db.BeginTxx(ctx, nil)

	if err != nil {
		return err
	}
	defer tx.Rollback()

	qwery := `UPDATE tasks SET task = $1, description = $2, done = $3, updated_at = NOW() WHERE id = $4`
	_, err = tx.ExecContext(ctx, qwery, task, description, done, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *DataStore) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx, err := s.db.BeginTxx(ctx, nil)

	if err != nil {
		return err
	}
	defer tx.Rollback()

	qwery := `DELETE FROM tasks WHERE id = $1`
	_, err = tx.ExecContext(ctx, qwery, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
