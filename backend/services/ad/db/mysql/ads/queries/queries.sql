-- name: GetAd :one
SELECT name, text,url FROM ads WHERE name = ?;

-- name: GetAds :many
SELECT name, text,url FROM ads ORDER BY rand() LIMIT 2;
