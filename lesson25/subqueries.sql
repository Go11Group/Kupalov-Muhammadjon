-- Subquerylarni qisqa qilib aytganda query ichida query yozish huddi nested loopga o'xshash
-- har xil turlari bor bular:

-- Scalar Subquery
-- ushbu turda subquery tabledan 1 ta ma'lumot qaytaradi qachon sharta bajarilsa
 SELECT column_name
    FROM   table
    WHERE  name = 'smth'
           (SELECT id
            FROM ids 
            WHERE shart);

-- Row subquery
    -- ushbu turda birnecha qator ma'lumot qaytariladi va outer queryda filterlanadi
     SELECT column_name
    FROM   table
    WHERE  name = 'smth'
           in (SELECT id
            FROM ids 
            WHERE shart);

-- Column subquery
    -- bu turda bir nechta ustun ma'lumot qaytaradi

-- Table subquery
-- bu query bir nechta table ma'lumotlari bilan birgalikda ishlayotganimizda asqotadi ya'ni 
-- table qaytariladi subqueryda
 SELECT column_name 
    FROM
        (SELECT column_name, ...
         FROM   table1 )
    WHERE  condition;
