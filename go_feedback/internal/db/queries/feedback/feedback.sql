-- name: CreateFeedback :one
INSERT INTO rizon_db.feedback (userId,comments, isLeaveReview, isEnjoying)
VALUES ($1, $2,$3, $4) RETURNING *;


-- name: GetFeedbackByUserID :one
SELECT id,userId,comments, isLeaveReview, isEnjoying, createdAt
FROM rizon_db.feedback
WHERE userId = $1;

-- name: GetFeedbackByID :one
SELECT id,userId,comments, isLeaveReview, isEnjoying, createdAt
FROM rizon_db.feedback
WHERE id = $1;


