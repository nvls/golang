package repository

import (
	"context"
	"database/sql"
	"demo09/internal/model"
	"errors"
	"fmt"
)

type TaskRepository interface {
	GetAll(ctx context.Context) ([]model.Task, error)
	GetById(ctx context.Context, id string) (model.Task, error)
	Create(ctx context.Context, task model.Task) (model.Task, error)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) GetAll(ctx context.Context) ([]model.Task, error) {
	const query = `SELECT id, title, created_at, modified_at FROM tasks`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Error query Tasks: %v", err)
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.CreatedAt, &task.ModifiedAt); err != nil {
			return nil, fmt.Errorf("Error scanning Task: %v", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *taskRepository) GetById(ctx context.Context, id string) (model.Task, error) {
	var task model.Task
	const query = `SELECT id, title, created_at, modified_at FROM tasks where id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&task.Id, &task.Title, &task.CreatedAt, &task.ModifiedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Task{}, sql.ErrNoRows
		}
		return model.Task{}, fmt.Errorf("Error query Tasks by ID: %v", err)
	}

	return task, nil
}

func (r *taskRepository) Create(ctx context.Context, task model.Task) (model.Task, error) {
	const query = `INSERT INTO tasks (title, created_at, modified_at) VALUES ($1, current_timestamp, current_timestamp) RETURNING id, title, created_at, modified_at`
	var ntask model.Task
	err := r.db.QueryRowContext(ctx, query, task.Title).Scan(&ntask.Id, &ntask.Title, &ntask.CreatedAt, &ntask.ModifiedAt)
	if err != nil {
		return model.Task{}, fmt.Errorf("Error creating Task: %v", err)
	}
	return ntask, nil
}
