-- name: CreateEmployee :one
INSERT INTO employee(
    code,
    name,
    email,
    phone_number
) VALUES(
    $1, $2, $3, $4
) RETURNING *;

-- name: GetEmployee :one
SELECT * FROM employee
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListEmployees :many
SELECT * FROM employee
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEmployee :one
UPDATE employee
SET name = $2,
email = $3,
phone_number = $4
WHERE id = $1
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employee
WHERE id = $1;
