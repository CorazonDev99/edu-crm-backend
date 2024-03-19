package auth

const (
	//ROLE
	CreateRoleQuery string = `INSERT INTO auth_role (title, 
	description,document) VALUES ($1, $2,$3)`
	UpdateRoleQuery = `UPDATE auth_role SET title = $1, 
	description = $2 , document=$3 WHERE id = $4`
	GetRoleListByRoleQuery = `SELECT id, title, description,document , 	COUNT(id) OVER() as total 
	FROM auth_role WHERE  deleted_at IS NULL ORDER BY title ASC LIMIT $1 OFFSET $2`
	GetRoleByIDQuery = `SELECT id, title, description,document 
	FROM auth_role WHERE id = $1 AND deleted_at IS NULL`
	GetRoleTitleByIDQuery = `SELECT  title	FROM auth_role WHERE id = $1 AND deleted_at IS NULL`
	GetRoleIDbyTitleQuery = `SELECT  id	FROM auth_role WHERE title = $1 AND deleted_at IS NULL`
	DeleteRoleQuery       = `UPDATE auth_role SET deleted_at = NOW(),
	title =CONCAT( title::TEXT , NOW()::TEXT) 	WHERE id = $1`
	CheckRoleByIDQuery = `SELECT id FROM auth_role WHERE id = $1 AND
	deleted_at IS NULL`

	//AUTH_ACCOUNT
	CreateAuthAccountQuery  = `INSERT INTO auth_account (account_id,role_id) VALUES ($1, $2)`
	UpdateAuthAccountQuery  = `UPDATE auth_account SET refresh_token = $1,	access_token=$2 , role_id =$3 WHERE	account_id=$4 AND deleted_at IS NULL `
	GetAuthAccountByIDQuery = `SELECT id , account_id, role_id, refresh_token, access_token FROM auth_account WHERE account_id = $1 AND deleted_at IS NULL`
	GetAllAuthAccountQuery  = `SELECT id , account_id, role_id, refresh_token, access_token FROM auth_account WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`
	DeleteAuthAccountQuery  = `UPDATE auth_account SET deleted_at = NOW() WHERE id = $1`
	//UpdateAuthAccountRefreshTokenQuery = `UPDATE auth_account SET refresh_token = $1 WHERE id = $2`
	//CheckAuthAccountIDQuery = `SELECT id FROM auth_account WHERE account_id = $1 AND deleted_at IS NULL`
	//AUTH_ACCOUNT_ROLE_ENROLLMENT
	//GetAuthAccountRoleEnrollmentQuery = `SELECT id, auth_account_id, role_id FROM auth_account_role_enrollment WHERE id = $1 AND deleted_at IS NULL

	CreatePermissionQuery = `INSERT INTO permission (title, description,tag,url,method) VALUES ($1, $2, $3,$4,$5)`
	UpdatePermissionQuery = `UPDATE permission SET title = $1, 
	description = $2, 	tag = $3 , url=$4,method=$5 WHERE id = $6`
	DeletePermissionQuery  = `UPDATE permission SET deleted_at = NOW() WHERE id = $1 and deleted_at IS NULL`
	GetPermissionListQuery = `SELECT id, title, description, tag,url,method ,COUNT(id) OVER() as total
	FROM permission  WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`
	GetPermissionByIDQuery = `SELECT id, title, description,tag,url,method,COUNT(id) OVER() as total FROM permission  WHERE id = $1 AND deleted_at IS NULL`
)
