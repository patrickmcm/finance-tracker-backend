CREATE TABLE instrument_prices (
    ticker varchar(12) REFERENCES instruments(ticker),
    timestamp timestamp with time zone not null,
    open numeric(14,6),
    close numeric(14,6),
    high numeric(14,6),
    low numeric(14,6),
    volume int,
    PRIMARY KEY (ticker,timestamp)
);