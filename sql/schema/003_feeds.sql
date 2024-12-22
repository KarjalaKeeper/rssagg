-- +goose Up
CREATE TABLE feeds (
 id UUID PRIMARY KEY,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL,
 name TEXT NOT NULL,
 url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE

);
-- +goose Down
DROP TABLE feeds;

 -- name: GetNextFeedToFetch :one
 SELECT *  from feeds
 ORDER BY last_fetched_at ASC NULLS FIRST
 LIMIT 1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_ad = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;