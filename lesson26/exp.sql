with students_avg_grade as (
    select 
        s.id as student_id, sc.course_id as course_id, round(avg(g.grade)::numeric, 2) as avarage_grade
    from student as s
    right join 
        student_course as sc
    on 
        s.id = sc.student_id
    right join 
        grade as g
    on 
        sc.id = g.student_course_id
    group by 
        sc.course_id, s.id
)

select course_id, max(avarage_grade)
from students_avg_grade
group by course_id;
