--- CARS AND USERS degan joyiga etibor bermabman course bilan student ga qildim toliq ulangan many2many

create table 
    courses(
        id serial primary key,
        name varchar,
        feild varchar
    );

create table 
    students(
        id serial primary key,
        name varchar,
        surname varchar,
        age int
    );


create table 
    courses_students(
        id serial primary key,
        student_id int references students(id),
        course_id int references courses(id)
    );

insert into
    courses(name, feild)
    values(
        'Golang',
        'Programming'
    );

insert into
    courses(name, feild)
    values(
        'English for IT',
        'Language'
    );

insert into
    courses(name, feild)
    values(
        'Flutter',
        'Programming'
    );


insert into
    students(name, surname, age)
    values(
        'Saidakbar',
        'Pardaboyev',
        19
    );

insert into
    students(name, surname, age)
    values(
        'Muhammadjon',
        'Ko''palov',
        19
    );

insert into
    students(name, surname, age)
    values(
        'Iskandar',
        'Iskandarov',
        22
    );

insert into
    students(name, surname, age)
    values(
        'Kamronbek',
        'Qashqarov',
        19
    );


insert into
    courses_students(course_id, student_id)
    values(
        1,
        1
    );

insert into
    courses_students(course_id, student_id)
    values(
        2,
        1
    );

insert into
    courses_students(course_id, student_id)
    values(
        1,
        2
    );

insert into
    courses_students(course_id, student_id)
    values(
        3,
        3
    );

insert into
    courses_students(course_id, student_id)
    values(
        2,
        3
    );


 select s.id, s.name, s.surname, array_agg(c.name) as courses
    from 
        students as s
    right join 
        courses_students as cs 
    on
        cs.student_id = s.id
    right join 
        courses as c
    on 
        c.id = cs.course_id
    group by s.id;


 select c.id, c.name, array_agg((s.name, s.surname)) as students
    from 
        courses as c
    right join 
        courses_students as cs 
    on
        cs.course_id = c.id
    right join 
        students as s
    on 
        s.id = cs.student_id
    group by c.id;

-- 