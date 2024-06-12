create type gender as enum(
    "male",
    "female"
)

create table users(
    id uuid primary key default gen_random_uuid(),
    first_name varchar not null,
    last_name varchar not null,
    age int not null,
    gender gender not null,
    nation varchar not null,
    feild varchar not null,
    parent_name varchar not null,
    city varchar not null
);