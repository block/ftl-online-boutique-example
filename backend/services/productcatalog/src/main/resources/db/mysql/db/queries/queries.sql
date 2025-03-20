-- name: QueryProduct :one
SELECT id, name, description, picture, price_usd, categories FROM product WHERE id = ?;

-- name: QueryProducts :many
SELECT id, name, description, picture, price_usd, categories FROM product;

-- name: InsertProduct :exec
INSERT INTO product (id, name, description, picture, price_usd, categories) VALUES (?, ?, ?, ?, ?, ?);