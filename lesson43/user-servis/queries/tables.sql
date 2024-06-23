create table users(
    id    uuid primary key default gen_random_uuid(),
    name  varchar,
    phone varchar,
    age   int
);
migrate -database 'postgres://postgres:root@localhost:5432/atto_users?sslmode=disable' -path migrations up
migrate create -ext sql -dir migrations -seq create_user_table