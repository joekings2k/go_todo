-- name: CreateTodo :one
INSERT INTO todos (
    title,
    description,
    user_id,
    status
) VALUES (
    $1, $2, $3 , $4
)
RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListTodosByUserID :many
SELECT * FROM todos
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: UpdateTodo :one
UPDATE todos
SET title = $2, description = $3, status = $4
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;