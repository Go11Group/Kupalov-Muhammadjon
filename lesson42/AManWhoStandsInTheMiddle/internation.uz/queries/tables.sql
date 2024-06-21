create table users(
    user_id uuid primary key default gen_random_uuid(),
    name varchar(100) not null,
    email varchar(100) unique not null,
    birthday timestamp not null,
    password varchar(100) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table courses(
    course_id uuid primary key default gen_random_uuid(),
    title varchar(100) not null,
    description text not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table lessons(
    lesson_id uuid primary key default gen_random_uuid(),
    course_id uuid references courses(course_id) not null,
    title varchar(100) not null,
    content text not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

create table enrollments(
    enrollment_id uuid primary key default gen_random_uuid(),
    user_id uuid references users(user_id) not null,
    course_id uuid references courses(course_id) not null,
    enrollment_date timestamp not null,
    created_at timestamp default current_timestamp,  
    updated_at timestamp,
    deleted_at timestamp
);


select 
        c.course_id, c.title, count(e.course_id) as enrollments_count
    from 
        courses as c
    join
        enrollments as e
    on
        c.course_id = e.course_id and 
        e.enrollment_date >= '2024-06-12' and
        e.enrollment_date <= '2025-06-16'
    group by 
        c.course_id, c.title
    order by 
        enrollments_count desc;