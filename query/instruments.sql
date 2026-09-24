-- name: ListInstruments :many
SELECT *
FROM instruments;

-- name: GetInstrument :one
SELECT *
FROM instruments
WHERE ticker = $1;