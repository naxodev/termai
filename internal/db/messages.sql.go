// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: messages.sql

package db

import (
	"context"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO messages (
    id,
    session_id,
    role,
    parts,
    created_at,
    updated_at
) VALUES (
    ?, ?, ?, ?, strftime('%s', 'now'), strftime('%s', 'now')
)
RETURNING id, session_id, role, parts, created_at, updated_at
`

type CreateMessageParams struct {
	ID        string `json:"id"`
	SessionID string `json:"session_id"`
	Role      string `json:"role"`
	Parts     string `json:"parts"`
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.queryRow(ctx, q.createMessageStmt, createMessage,
		arg.ID,
		arg.SessionID,
		arg.Role,
		arg.Parts,
	)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.SessionID,
		&i.Role,
		&i.Parts,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteMessage = `-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = ?
`

func (q *Queries) DeleteMessage(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteMessageStmt, deleteMessage, id)
	return err
}

const deleteSessionMessages = `-- name: DeleteSessionMessages :exec
DELETE FROM messages
WHERE session_id = ?
`

func (q *Queries) DeleteSessionMessages(ctx context.Context, sessionID string) error {
	_, err := q.exec(ctx, q.deleteSessionMessagesStmt, deleteSessionMessages, sessionID)
	return err
}

const getMessage = `-- name: GetMessage :one
SELECT id, session_id, role, parts, created_at, updated_at
FROM messages
WHERE id = ? LIMIT 1
`

func (q *Queries) GetMessage(ctx context.Context, id string) (Message, error) {
	row := q.queryRow(ctx, q.getMessageStmt, getMessage, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.SessionID,
		&i.Role,
		&i.Parts,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listMessagesBySession = `-- name: ListMessagesBySession :many
SELECT id, session_id, role, parts, created_at, updated_at
FROM messages
WHERE session_id = ?
ORDER BY created_at ASC
`

func (q *Queries) ListMessagesBySession(ctx context.Context, sessionID string) ([]Message, error) {
	rows, err := q.query(ctx, q.listMessagesBySessionStmt, listMessagesBySession, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Message{}
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.SessionID,
			&i.Role,
			&i.Parts,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateMessage = `-- name: UpdateMessage :exec
UPDATE messages
SET
    parts = ?,
    updated_at = strftime('%s', 'now')
WHERE id = ?
`

type UpdateMessageParams struct {
	Parts string `json:"parts"`
	ID    string `json:"id"`
}

func (q *Queries) UpdateMessage(ctx context.Context, arg UpdateMessageParams) error {
	_, err := q.exec(ctx, q.updateMessageStmt, updateMessage, arg.Parts, arg.ID)
	return err
}
