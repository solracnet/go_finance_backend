// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
insert into users (username, password, email) values ($1, $2, $3) returning id, username, password, email, created_at, updated_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.Username, arg.Password, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
select id, username, password, email, created_at, updated_at from users where username ilike $1 limit 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
select id, username, password, email, created_at, updated_at from users where id = $1 limit 1
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
