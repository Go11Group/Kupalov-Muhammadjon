create type transaction_type as enum('credit', 'debit');

create table card(
    id      uuid primary key default gen_random_uuid(),
    number  varchar,
    user_id uuid not null
);

create table transaction (
    id          uuid primary key default gen_random_uuid(),
    card_id     uuid references card(id) not null,
    amount      int,
    terminal_id uuid default null
);

create table station(
    id   uuid primary key default gen_random_uuid(),
    name varchar
);

create table terminal(
    id          uuid primary key default gen_random_uuid(),
    station_id  uuid references station(id)
);

migrate -database 'postgres://postgres:root@localhost:5432/atto_billing?sslmode=disable' -path migrations down4
migrate create -ext sql -dir migrations -seq create_tables