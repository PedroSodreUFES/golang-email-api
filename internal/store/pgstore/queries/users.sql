-- name: CreateUser :one
INSERT INTO users (
    full_name, email, password_hash
) VALUES ( $1, $2, $3)
RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: UpdateUserProfilePhoto :exec
UPDATE users
SET profile_picture = $2
WHERE id = $1;

-- name: DeleteUserById :execrows
DELETE FROM users
WHERE id = $1;