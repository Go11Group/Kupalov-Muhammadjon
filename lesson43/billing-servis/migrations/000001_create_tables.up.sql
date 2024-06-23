create type transaction_type as enum('credit', 'debit');

create table cards(
    id      uuid primary key default gen_random_uuid(),
    number  varchar,
    user_id uuid not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table transactions (
    id          uuid primary key default gen_random_uuid(),
    card_id     uuid references card(id) not null,
    amount      int,
    terminal_id uuid default null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table stations(
    id   uuid primary key default gen_random_uuid(),
    name varchar,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table terminals(
    id          uuid primary key default gen_random_uuid(),
    station_id  uuid references station(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);