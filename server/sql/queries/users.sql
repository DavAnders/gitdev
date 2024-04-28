-- name: CreateUser :one
INSERT INTO users (
        github_created_at,
        access_token,
        name,
        username,
        github_id,
        email,
        followers,
        following,
        bio,
        avatar_url,
        location
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11
    )
RETURNING github_id;
-- name: UpdateUserInfo :one
UPDATE users
SET access_token = $1,
    name = $2,
    email = $3,
    bio = $4,
    title = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE github_id = $6
RETURNING *;
-- name: GetAllUsers :many
SELECT *
FROM users
LIMIT 20;
-- name: GetUserByGithubID :one
SELECT users.*,
    COUNT(repos.*) AS num_repos
FROM users
    LEFT JOIN repos ON users.github_id = repos.user_github_id
WHERE github_id = $1
GROUP BY users.id;