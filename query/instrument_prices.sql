-- name: GetInstrumentPrices :many
SELECT *
FROM instrument_prices
WHERE ticker=$1 AND (timestamp between @datefrom AND @dateTo);