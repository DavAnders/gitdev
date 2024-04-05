// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
        id,
        access_token,
        name,
        username,
        github_id,
        email,
        panel_body,
        avatar_url
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8
    )
RETURNING id
`

type CreateUserParams struct {
	ID          uuid.UUID
	AccessToken string
	Name        string
	Username    string
	GithubID    int32
	Email       string
	PanelBody   sql.NullString
	AvatarUrl   string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.AccessToken,
		arg.Name,
		arg.Username,
		arg.GithubID,
		arg.Email,
		arg.PanelBody,
		arg.AvatarUrl,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, created_at, updated_at, access_token, name, username, github_id, email, followers, following, panel_body, role, avatar_url
FROM users
LIMIT 20
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AccessToken,
			&i.Name,
			&i.Username,
			&i.GithubID,
			&i.Email,
			&i.Followers,
			&i.Following,
			&i.PanelBody,
			&i.Role,
			&i.AvatarUrl,
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

const getUserByGitHubID = `-- name: GetUserByGitHubID :one
SELECT id, created_at, updated_at, access_token, name, username, github_id, email, followers, following, panel_body, role, avatar_url
FROM users
WHERE github_id = $1
`

func (q *Queries) GetUserByGitHubID(ctx context.Context, githubID int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByGitHubID, githubID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccessToken,
		&i.Name,
		&i.Username,
		&i.GithubID,
		&i.Email,
		&i.Followers,
		&i.Following,
		&i.PanelBody,
		&i.Role,
		&i.AvatarUrl,
	)
	return i, err
}

const updateUserToken = `-- name: UpdateUserToken :exec
UPDATE users
SET access_token = $1,
    updated_at = CURRENT_TIMESTAMP
`

func (q *Queries) UpdateUserToken(ctx context.Context, accessToken string) error {
	_, err := q.db.ExecContext(ctx, updateUserToken, accessToken)
	return err
}
