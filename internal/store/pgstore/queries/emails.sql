-- name: CreateEmail :one
INSERT INTO emails (
    title, content, id_receiver, id_sender
) VALUES ( $1, $2, $3, $4)
RETURNING *;

-- name: GetEmailById :one
SELECT *
FROM emails
WHERE id = $1;

-- name: DeleteEmail :execrows
DELETE FROM emails
WHERE id = $1;

-- name: GetMySentEmails :many
SELECT *
FROM emails e
INNER JOIN users u ON e.id_receiver = u.id
WHERE id_sender = $1
ORDER BY e.id;

-- name: GetMyReceivedEmails :many
SELECT *
FROM emails e
INNER JOIN users u ON e.id_sender = u.id
WHERE id_receiver = $1
ORDER BY e.id;

-- name: UpdateAndGetEmailByID :one
UPDATE emails
SET wasSeen = TRUE
WHERE id = $1
RETURNING *;