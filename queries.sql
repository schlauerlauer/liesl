-- name: GetNode :one
SELECT * FROM nodes
WHERE id = ? LIMIT 1;
