-- name: CreateFeed :one
INSERT INTO feeds (id, name, url,user_id,created_at, updated_at)
VALUES ($1, $2, $3, $4,$5,$6)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE id = $1;

-- name: GetUserFeeds :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetNextFeedsToFetched :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedASFetched :one
Update feeds SET last_fetched_at =NOW(),updated_at =NOW() WHERE id=$1
RETURNING *;