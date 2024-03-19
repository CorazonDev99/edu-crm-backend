package course

const (
	CreateCourseQuery = `INSERT INTO course (title,description,photo,duration,
	status,price,lesson_duration) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	UpdateCourseQuery  = `UPDATE course SET title=$1,description=$2,duration=$3,status=$4,price=$5,lesson_duration=$6,	photo= $7,updated_at = NOW() WHERE id=$8 AND deleted_at IS NULL`
	DeleteCourseQuery  = `UPDATE course SET deleted_at = NOW() ,updated_at = NOW() WHERE id=$1 AND deleted_at IS NULL`
	GetCourseByIDQuery = `SELECT id,title,description,duration,status,price,lesson_duration,photo FROM course WHERE id=$1 AND deleted_at IS NULL`
	CheckCourseByID    = `SELECT id FROM course WHERE id=$1 AND deleted_at IS NULL`
	GetCourseListQuery = `SELECT id , title,description,duration,status,photo,price,lesson_duration,COUNT(id) OVER(
	) as total  FROM course WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`
	CreateTeacherCourseEnrollmentQuery = `INSERT INTO course_teacher_enrollment (teacher_id,course_id) VALUES ($1,$2)`
	DeleteTeacherCourseEnrollmentQuery = `UPDATE course_teacher_enrollment SET deleted_at = NOW() ,updated_at = NOW() WHERE teacher_id=$1 AND course_id=$2`
	UpdateTeacherCourseEnrollmentQuery = `UPDATE course_teacher_enrollment SET
	teacher_id=$1, updated_at = NOW() WHERE teacher_id=$2 AND course_id=$3 `
	GetTeacherCourseListQuery = `SELECT id , title,description,duration,status,price,lesson_duration FROM course INNER JOIN course_teacher_enrollment ON course_teacher_enrollment.course_id = course.id WHERE course_teacher_enrollment.teacher_id=$1 AND  WHERE deleted_at IS NULL LIMIT $2 OFFSET $3`
	GetStudentCourseListQuery = `SELECT id , title,description,duration,status,price,lesson_duration FROM course INNER JOIN edu_group ON edu_group.course_id = course.id INNER JOIN edu_group_leaner_enrollment ON edu_group.id=edu_group_leaner_enrollment.learner_id WHERE edu_group_leaner_enrollment.learner_id=$1  AND edu_group_leaner_enrollment.deleted_at IS NULL LIMIT $2 OFFSET $3`
)
