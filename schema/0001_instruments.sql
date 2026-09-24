CREATE TYPE instrument_type AS ENUM ('ETF', 'STOCK');

CREATE TABLE instruments (
    ticker varchar(12) PRIMARY KEY,
    name varchar(255) NOT NULL,
    currency varchar(3) DEFAULT 'GBX',
    isin varchar(12) UNIQUE NOT NULL,
    type instrument_type NOT NULL
);