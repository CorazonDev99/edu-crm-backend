package user

const (
	CreateUserQuery = `INSERT INTO crm_user (full_name,birthday_date,
	added_date,phone_number,role_id,password,extra_data,photo) VALUES ($1,$2,$3,$4,
	$5,$6,$7,$8) RETURNING id`
	UpdateUserQuery = `UPDATE crm_user SET full_name=$1,
	birthday_date=$2 ,added_date=$3,phone_number=$4,role_id=$5,extra_data=$6,photo=$7,
	updated_at = NOW() WHERE id=$8`
	UpdateUserPasswordQuery = `UPDATE crm_user SET password=$1,updated_at = NOW() WHERE id=$2 AND deleted_at IS NULL`
	DeleteUserQuery         = `UPDATE crm_user SET deleted_at = NOW() ,updated_at = NOW() WHERE id=$1 AND deleted_at IS NULL`
	GetUserByIDQuery        = `SELECT id,full_name,phone_number,birthday_date,role_id,extra_data,added_date,photo FROM crm_user WHERE id=$1 AND deleted_at IS NULL`
	CheckUserByIDQuery      = `SELECT id FROM crm_user WHERE id=$1 AND deleted_at IS NULL`
	GetUserListByRoleQuery  = `SELECT id,full_name,phone_number,birthday_date,role_id,extra_data,
	added_date ,photo,COUNT(id) OVER() as total FROM crm_user WHERE deleted_at IS	NULL AND role_id=$1 LIMIT $2 OFFSET $3`
	GetUserListAllQuery = `SELECT id,full_name,phone_number,
	birthday_date,role_id,extra_data,photo,
	added_date,COUNT(id) OVER() as total FROM crm_user WHERE deleted_at IS NULL LIMIT $1
	OFFSET $2`
	SignInUserQuery = `SELECT id,
	role_id FROM crm_user WHERE  phone_number=$1  AND password=$2 AND
	deleted_at IS NULL`
	GetGroupStudentListQuery = `SELECT id,full_name,phone_number,birthday_date,role_id,extra_data,added_date,photo FROM crm_user WHERE id= ANY($1) AND deleted_at IS NULL `
)
