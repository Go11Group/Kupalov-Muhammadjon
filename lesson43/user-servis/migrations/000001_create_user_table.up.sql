create table users(
    id    uuid primary key default gen_random_uuid(),
    name  varchar,
    phone varchar,
    age   int,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);