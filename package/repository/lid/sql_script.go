package lid

const (
	// Board
	CreateBoardQuery  = `INSERT INTO board (title) VALUES ($1)`
	UpdateBoardQuery  = `UPDATE board SET title = $1 WHERE id = $2`
	DeleteBoardQuery  = `UPDATE board SET deleted_at = NOW() WHERE id = $1`
	GetBoardListQuery = `SELECT id, title,	COUNT(id) OVER() as total	FROM board WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`
	GetBoardByIDQuery = `SELECT id, title FROM board WHERE id = $1 AND deleted_at IS NULL`
	CheckBoardByID    = `SELECT id FROM board WHERE id = $1 AND deleted_at IS NULL`
	// List
	CreateListQuery       = `INSERT INTO list (board_id, title) VALUES ($1, $2)`
	UpdateListQuery       = `UPDATE list SET board_id = $1, title = $2 WHERE id = $3`
	DeleteListQuery       = `UPDATE list SET deleted_at = NOW() WHERE id = $1`
	GetListListQuery      = `SELECT id, board_id,	title ,COUNT(id) OVER() as total FROM list WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`
	GetListByIDQuery      = `SELECT id, board_id, title FROM list WHERE id = $1 AND deleted_at IS NULL`
	CheckListByID         = `SELECT id FROM list WHERE id = $1 AND deleted_at IS NULL`
	GetBoardListListQuery = `SELECT id, board_id, 	title FROM list WHERE board_id = $1 AND deleted_at IS NULL`
	ListMoveQuery         = `	UPDATE list SET board_id=$3 where id=$1 AND	board_id=$2	AND deleted_at IS NULL`
	// Lid
	CreateLidQuery      = `INSERT INTO lid (list_id, full_name, phone_number, location, comment) VALUES ($1, $2, $3, $4, $5)`
	UpdateLidQuery      = `UPDATE lid SET list_id = $1, full_name = $2, phone_number = $3, location = $4, comment = $5 WHERE id = $6`
	DeleteLidQuery      = `UPDATE lid SET deleted_at = NOW() WHERE id = $1`
	GetLidListQuery     = `SELECT id, list_id, full_name, phone_number, location,created_at,comment,COUNT(id) OVER() as total FROM lid WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`
	GetLidByIDQuery     = `SELECT id, list_id, full_name, phone_number, location, comment,created_at FROM lid WHERE id = $1 AND deleted_at IS NULL`
	CheckLidByID        = `SELECT id FROM lid WHERE id = $1 AND deleted_at IS NULL`
	LidMoveQuery        = `	UPDATE lid SET list_id=$3 where id=$1 AND list_id=$2	AND deleted_at IS NULL`
	LidReplaceQuery     = `	UPDATE lid SET order=$3 where id=$1 AND order=$2 AND	deleted_at IS NULL`
	GetLidBYListIDQuery = `SELECT id, list_id, full_name, phone_number, location, comment,created_at FROM lid WHERE list_id = $1 AND deleted_at IS NULL`
)
