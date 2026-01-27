-- name: CreateUser :one
INSERT INTO users (
    full_name, email, password_hash
) VALUES ( $1, $2, $3)
RETURNING *;

-- name: GetUserByEmail :one
SELECT
    id,
    full_name,
    password_hash,
    profile_picture
FROM users
WHERE email = $1;

-- name: GetUserById :one
SELECT
    id,
    full_name,
    password_hash,
    profile_picture
FROM users
WHERE id = $1;

-- name: UpdateUserProfilePhoto :exec
UPDATE users
SET profile_picture = $2
WHERE id = $1;

-- name: DeleteUserById :exec
DELETE FROM users
WHERE id = $1;