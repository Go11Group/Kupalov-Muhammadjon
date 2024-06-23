CREATE TYPE transaction_type AS ENUM ('credit', 'debit');

CREATE TABLE cards (
    id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    number      varchar,
    user_id     uuid NOT NULL,
    created_at  timestamp DEFAULT current_timestamp,
    updated_at  timestamp,
    deleted_at  timestamp
);

CREATE TABLE transactions (
    id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    card_id         uuid REFERENCES cards(id) NOT NULL,
    amount          int,
    transaction_type transaction_type,
    terminal_id     uuid DEFAULT NULL,
    created_at      timestamp DEFAULT current_timestamp,
    updated_at      timestamp,
    deleted_at      timestamp
);

CREATE TABLE stations (
    id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name        varchar,
    created_at  timestamp DEFAULT current_timestamp,
    updated_at  timestamp,
    deleted_at  timestamp
);

CREATE TABLE terminals (
    id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    station_id  uuid REFERENCES stations(id),
    created_at  timestamp DEFAULT current_timestamp,
    updated_at  timestamp,
    deleted_at  timestamp
);
