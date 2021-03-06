// Code generated by sqlc. DO NOT EDIT.
// source: task_group.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTaskGroup = `-- name: CreateTaskGroup :one
INSERT INTO task_group (project_id, created_at, name, position)
  VALUES($1, $2, $3, $4) RETURNING task_group_id, project_id, created_at, name, position
`

type CreateTaskGroupParams struct {
	ProjectID uuid.UUID `json:"project_id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Position  float64   `json:"position"`
}

func (q *Queries) CreateTaskGroup(ctx context.Context, arg CreateTaskGroupParams) (TaskGroup, error) {
	row := q.db.QueryRowContext(ctx, createTaskGroup,
		arg.ProjectID,
		arg.CreatedAt,
		arg.Name,
		arg.Position,
	)
	var i TaskGroup
	err := row.Scan(
		&i.TaskGroupID,
		&i.ProjectID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
	)
	return i, err
}

const deleteTaskGroupByID = `-- name: DeleteTaskGroupByID :execrows
DELETE FROM task_group WHERE task_group_id = $1
`

func (q *Queries) DeleteTaskGroupByID(ctx context.Context, taskGroupID uuid.UUID) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteTaskGroupByID, taskGroupID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getAllTaskGroups = `-- name: GetAllTaskGroups :many
SELECT task_group_id, project_id, created_at, name, position FROM task_group
`

func (q *Queries) GetAllTaskGroups(ctx context.Context) ([]TaskGroup, error) {
	rows, err := q.db.QueryContext(ctx, getAllTaskGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskGroup
	for rows.Next() {
		var i TaskGroup
		if err := rows.Scan(
			&i.TaskGroupID,
			&i.ProjectID,
			&i.CreatedAt,
			&i.Name,
			&i.Position,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTaskGroupByID = `-- name: GetTaskGroupByID :one
SELECT task_group_id, project_id, created_at, name, position FROM task_group WHERE task_group_id = $1
`

func (q *Queries) GetTaskGroupByID(ctx context.Context, taskGroupID uuid.UUID) (TaskGroup, error) {
	row := q.db.QueryRowContext(ctx, getTaskGroupByID, taskGroupID)
	var i TaskGroup
	err := row.Scan(
		&i.TaskGroupID,
		&i.ProjectID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
	)
	return i, err
}

const getTaskGroupsForProject = `-- name: GetTaskGroupsForProject :many
SELECT task_group_id, project_id, created_at, name, position FROM task_group WHERE project_id = $1
`

func (q *Queries) GetTaskGroupsForProject(ctx context.Context, projectID uuid.UUID) ([]TaskGroup, error) {
	rows, err := q.db.QueryContext(ctx, getTaskGroupsForProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskGroup
	for rows.Next() {
		var i TaskGroup
		if err := rows.Scan(
			&i.TaskGroupID,
			&i.ProjectID,
			&i.CreatedAt,
			&i.Name,
			&i.Position,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setTaskGroupName = `-- name: SetTaskGroupName :one
UPDATE task_group SET name = $2 WHERE task_group_id = $1 RETURNING task_group_id, project_id, created_at, name, position
`

type SetTaskGroupNameParams struct {
	TaskGroupID uuid.UUID `json:"task_group_id"`
	Name        string    `json:"name"`
}

func (q *Queries) SetTaskGroupName(ctx context.Context, arg SetTaskGroupNameParams) (TaskGroup, error) {
	row := q.db.QueryRowContext(ctx, setTaskGroupName, arg.TaskGroupID, arg.Name)
	var i TaskGroup
	err := row.Scan(
		&i.TaskGroupID,
		&i.ProjectID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
	)
	return i, err
}

const updateTaskGroupLocation = `-- name: UpdateTaskGroupLocation :one
UPDATE task_group SET position = $2 WHERE task_group_id = $1 RETURNING task_group_id, project_id, created_at, name, position
`

type UpdateTaskGroupLocationParams struct {
	TaskGroupID uuid.UUID `json:"task_group_id"`
	Position    float64   `json:"position"`
}

func (q *Queries) UpdateTaskGroupLocation(ctx context.Context, arg UpdateTaskGroupLocationParams) (TaskGroup, error) {
	row := q.db.QueryRowContext(ctx, updateTaskGroupLocation, arg.TaskGroupID, arg.Position)
	var i TaskGroup
	err := row.Scan(
		&i.TaskGroupID,
		&i.ProjectID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
	)
	return i, err
}
