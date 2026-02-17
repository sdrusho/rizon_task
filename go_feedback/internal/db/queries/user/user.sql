-- name: CreateUser :one
INSERT INTO rizon_db.user (email,name, status, deviceId)
VALUES ($1, $2,$3, $4) RETURNING *;


-- name: GetUserByID :one
SELECT id, userId, name, email, status, createdAt, isDeleted
FROM rizon_db.user
WHERE userId = $1 AND isDeleted = false;

-- name: GetUserByEmail :one
SELECT id, userId, name, email, status, createdAt, isDeleted
FROM rizon_db.user
WHERE email = $1 AND isDeleted = false;

-- name: GetUserById :one
SELECT id, userId, name, email, status, createdAt, isDeleted
FROM rizon_db.user
WHERE userId = $1 AND status = $2;

