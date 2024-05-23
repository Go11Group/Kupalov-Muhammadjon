create table rasta(id serial, name varchar, price real);

insert into rasta(name, price) values('Olma', 18000);
insert into rasta(name, price) values('Banan', 17000);
insert into rasta(name, price) values('Gilos', 20000);
insert into rasta(name, price) values('Qulupnay', 15000);
insert into rasta(name, price) values('Nok', 15000);
insert into rasta(name, price) values('Pomidor', 10000);
insert into rasta(name, price) values('Bodring', 10000);
insert into rasta(name, price) values('Apelsin', 15000);
insert into rasta(name, price) values('Limon', 20000);
insert into rasta(name, price) values('Kartoshka', 6000);
insert into rasta(name, price) values('Piyoz', 8000);

select * from rasta 

-- Order by sortlab beradi desak ham bo'ladi unga nima bo'yicha saralashini bersak bo'ldi bu keysda price 
-- bo'yicha saralab beradi minimum dan boshlab pastga ketadi kattalashib
select * from rasta order by price

-- bu holatda saralash maximum dan minmumga qarab ketadi kamayish tartibida
select * from rasta order by price desc

-- update yangilaydi yani shart qanoatlantirilsa set qismi ishga tushadi bu holatda 
-- table dan name qatori Bananga teng joyni topa olsa price ni 14000 tenglashtirib qo'yadi
update rasta set price = 14000 where name = 'Banan';
select * from rasta;

-- shart to'g'ri bo'lgan holatda o'sha qatorni o'chiradi 
delete from rasta where name = 'Banan';

-- group qilish uchun yangi column  qo'shib to'ldirib qo'ydim update bilan
alter table rasta add tur varchar;

update rasta set tur = 'meva' where id = 12;

-- Group by ozi agregat funksiyalar bilan ishlatiladi bular count(), max(), min(), sum(), avg() ya'ni
-- bir guruh malumotlar ustida amallar bajarsak ishlatishimiz mumkin.
select tur, COUNT(*) from rasta group by tur;
select tur, avg(price) from rasta group by tur;

-- Joinlar ozi nimaga kerak. Bitta dastur databazasida ancha ko'p tablelar bo'lishi mumkin va userga malumotlar
-- turli tablellardan keladi bu holatda biz bir nrchta tablelar bilan ishlaymiz va to'g'ri malumotlarni
-- olish uchun tablellar orasidagi bog'liqlik orqali ularni birlashtirib kerak malumotni olishimish kerak 
-- bo'ladi shu holatlarda ishlatamiz(va hokazo).
-- Inner Join ikkala tableda bor bo'lgan valuelarnigina qaytaradi.
create table dokon(id serial, name varchar, rasta_id int)
insert into dokon(name, rasta_id) values('Ali dokini', 1)
insert into dokon(name, rasta_id) values('Ali dokini', 12);
-- on qismida shart beriladi va osha shart bajarilgandagina ikkala table qo'shiladi 
select * from rasta as r inner join dokon as d  on d.rasta_id = r.id;
-- agar right join qilsak o'ng tarafdagi tabledan hamma ma'lumot olinadi chap tarafdan ma'lumot olinadi
-- shart bajarilsagina bo'lmasa hech qanday ma'lumot kelmaydi chapdan. Left join esa xuddi shu narsani teskarisi
select * from rasta as r right join dokon as d  on d.rasta_id = r.id;
select * from rasta as r left join dokon as d  on d.rasta_id = r.id;

-- full join qilsa hammasini qo'shadi shart to'g'ri kelsa ham kelmasa ham shart to'g'ri kelsa match bo'lgan pairlarni
-- qo'shadi shart to'g'ri kelmasa null bo'lib qoluradi osha columnlar
select * from rasta as r full join dokon as d  on d.rasta_id = r.id;




