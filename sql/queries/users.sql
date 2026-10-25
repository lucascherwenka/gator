-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT
*
FROM
users
WHERE
name = $1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT
*
FROM
users;

-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT
feeds.name AS feed_name,
feeds.url,
users.name AS user_name
FROM
feeds
INNER JOIN
users ON feeds.user_id = users.id;

-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)

SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedByUrl :one
SELECT
id,
name,
user_id,
url
FROM
feeds
WHERE
url = $1;

-- name: GetFeedFollowsForUser :many
SELECT
feed_follows.*,
feeds.name AS feed_name,
users.name AS user_name
FROM
feed_follows
INNER JOIN users ON feed_follows.user_id = users.id
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE
feed_follows.user_id = $1;

-- name: DeleteFollowForUser :exec
DELETE FROM feed_follows
WHERE
user_id = $1 AND feed_id = $2;

-- name: MarkFeedFetched :one
UPDATE 
feeds 
SET 
last_fetched_at = NOW(),
updated_at = NOW()
WHERE
feeds.id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT
*
FROM
feeds
ORDER BY
last_fetched_at NULLS FIRST
LIMIT
1;

-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT
posts.*,
feeds.name AS feed_name
FROM
posts
INNER JOIN
feed_follows ON posts.feed_id = feed_follows.feed_id
INNER JOIN feeds ON feeds.id = posts.feed_id
WHERE
feed_follows.user_id = $1
ORDER BY
published_at DESC
LIMIT
$2;