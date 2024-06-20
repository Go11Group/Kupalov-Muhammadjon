create table users(
    id uuid default gen_random_uuid(),
    name varchar,
    age int
    )