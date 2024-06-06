select c.name
from
    student_course as sc 
join 
    course as c
on
    sc.course_id = c.id and $1 = sc.student_id 


select c.name, g.grade
from
    student_course as sc 
join 
    course as c
on
    sc.course_id = c.id and $1 = sc.student_id 
join 
    grade as g
on
    g.student_course_id = sc.id