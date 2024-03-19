package group

const (
	CreateGroupQuery = `INSERT INTO edu_group (title,
	course_id,teacher_id,edu_days,room_id,price,lesson_start_time,status,
	start_date,end_date,comment) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	UpdateGroupQuery = `UPDATE edu_group SET title=$1,
	course_id=$2,teacher_id=$3,edu_days=$4,room_id=$5,price=$6,
	lesson_start_time=$7,status=$8,status_date = $9,end_date=$10,comment=$11,updated_at = NOW() WHERE id=$8`
	DeleteGroupQuery  = `UPDATE edu_group SET deleted_at = NOW() ,updated_at = NOW() WHERE id=$1 `
	GetGroupByIDQuery = `SELECT id,title,course_id,
	teacher_id,edu_days,room_id,price,lesson_start_time,status,start_date,
	end_date,comment FROM edu_group WHERE id=$1 AND deleted_at IS NULL`
	GetGroupListQuery = `SELECT id,title,course_id,	teacher_id,edu_days,room_id,price, lesson_start_time,status,start_date,
	end_date,comment , COUNT(id) OVER() as total FROM edu_group WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`
	GetTeacherGroupListQuery   = `SELECT id,title,course_id,teacher_id,edu_days,room_id,price, lesson_start_time,status,start_date,	end_date,comment,COUNT(id) OVER() as total FROM edu_group WHERE teacher_id = $1 AND deleted_at IS NULL LIMIT $2 OFFSET $3`
	GetCourseGroupListQuery    = `SELECT id,title,course_id,teacher_id,edu_days,room_id,price, lesson_start_time,status,start_date,	end_date,comment,COUNT(id) OVER() as total FROM edu_group WHERE course_id = $1 AND deleted_at IS NULL LIMIT $2 OFFSET $3`
	CreateGroupEnrollmentQuery = `INSERT INTO	edu_group_learner_enrollment (	group_id,learner_id) VALUES ($1,$2)`
	DeleteGroupEnrollmentQuery = `UPDATE edu_group_learner_enrollment SET
	deleted_at = NOW() ,updated_at = NOW() WHERE group_id=$1 AND deleted_at IS NULL`
	CheckGroupByIDQuery               = `SELECT id FROM edu_group WHERE id=$1 AND deleted_at IS NULL`
	DeleteUserAllGroupEnrollmentQuery = `UPDATE edu_group_learner_enrollment SET
	deleted_at = NOW() ,updated_at = NOW() WHERE learner_id=$1`
	DeleteUserFromGroupQuery = `UPDATE edu_group_learner_enrollment SET
	deleted_at = NOW() ,updated_at = NOW() WHERE learner_id=$1 AND group_id=$2`
	GetStudentGroupListQuery = `SELECT edu_group_learner_enrollment.learner_id, COUNT(id) OVER() as total FROM edu_group INNER JOIN edu_group_learner_enrollment ON edu_group.id=edu_group_learner_enrollment.group_id WHERE edu_group_learner_enrollment.learner_id = $1 AND edu_group.deleted_at IS NULL `
)
