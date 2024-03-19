package admin

const (
	GetSettingsQuery = `SELECT company_title,company_logo,
	system_enter_logo,open_date,company_phone,site_color,instruction_file FROM settings WHERE deleted_at IS NULL`
	UpsertSettingsQuery = `SELECT upsert_settings($1,$2,$3,$4,$5,$6,$7)`
	CreateRoomQuery     = `INSERT INTO rooms (title,description,room_number,
	open_time,	close_time) VALUES ($1,$2,$3,$4,$5)`
	UpdateRoomQuery = `UPDATE rooms SET description=$1,title=$2,
	room_number=$3,open_time=$4,close_time=$5 WHERE id=$6`
	DeleteRoomQuery = `UPDATE rooms SET deleted_at = NOW() ,
	updated_at = NOW() WHERE id=$1`
	GetRoomByIDQuery = `SELECT id,description,title,room_number,open_time,
	close_time FROM rooms WHERE id=$1 AND deleted_at IS NULL`
	GetRoomListQuery = `SELECT id,description,title,room_number,open_time,
	close_time,COUNT(id) OVER() as total  FROM rooms WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`
	CheckRoomByIDQuery        = `SELECT id FROM rooms WHERE id=$1 AND deleted_at IS NULL`
	GetGroupListByRoomIDQuery = `SELECT edu_group.id,edu_group.title,edu_group.course_id,
	edu_group.teacher_id,edu_group.edu_days,edu_group.room_id,edu_group.price, edu_group.lesson_start_time,edu_group.status,edu_group.start_date,
	edu_group.end_date,edu_group.comment , COUNT(edu_group.id) OVER() as total 
	FROM edu_group INNER JOIN rooms ON edu_group.room_id = rooms.id WHERE rooms.deleted_at IS NULL AND edu_group.deleted_at IS NULL AND rooms.id=$1 LIMIT $2 OFFSET $3`
)
