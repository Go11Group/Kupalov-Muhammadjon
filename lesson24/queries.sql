create table author (
    id serial primary key,
    name varchar not null
);

create table book
(
    id          serial primary key,
    name        varchar not null,
    page        int,
    author_id   int references author(id)
);

-- reference ozi Sql tillari asosi yani referenceni relation desak ham bo'ladi tablelarni bog'lashda ishlatamiz
-- shuning uchun ham relation databse tili hisoblanadi. Bu yerda ham bookni authorga bog'layapmiz.

insert into author (name) values('Alisher Navoiy')
insert into author (name) values ('G''afur G''ulom');

-- bu yerda 1 va 2 chi id li authorlar quyidagi booklarga ulanadi.
insert into book(name, page, author_id) values('Hamsa', 1234, 1)
insert into book(name, page, author_id) values('Shum bola', 327, 2)

-- references yozildimi u yerdagi bog'liqlikni buzib bo'lmaydi bu yerda mavjud bo'lmagan author_id bersak
-- error beradi chunchi 3 idli author mavjud emas hali 
book_store=# insert into book(name, page, author_id) values('Shum bola', 327, 3);
ERROR:  insert or update on table "book" violates foreign key constraint "book_author_id_fkey"
DETAIL:  Key (author_id)=(3) is not present in table "author".
book_store=# 

-- Agar author tabledagi biror malumot bookga ulansa uniyam o'chirib bo'lmaydi chunki bu relationni buzadi.
book_store=# delete from author where id = 1;
ERROR:  update or delete on table "author" violates foreign key constraint "book_author_id_fkey" on table "book"
DETAIL:  Key (id)=(1) is still referenced from table "book".


