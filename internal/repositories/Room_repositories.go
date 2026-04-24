package repositories

import (
	"Web-Chat/internal/domain/repository"
	"database/sql"
	"time"
)

type RoomRepo struct {
	db *sql.DB
}

func NewRoomRepo(db *sql.DB) repository.Room {
	return &RoomRepo{db: db}
}
func (RR *RoomRepo) DeleteRoom(RoomId int) error {
	SqlStatemnt := (`DELETE FROM private_chats WHERE id=$1`)
	_, err := RR.db.Exec(SqlStatemnt, RoomId)
	if err != nil {
		return err
	}
	return nil
}
func (RR *RoomRepo) FindRoomIdByUsersId(UserId1 int, UserId2 int) (int, error) {
	var RoomId int
	SqlStatemnt := (`SELECT id FROM private_chats WHERE (user1_id=$1 AND user2_id = $2) OR (user2_id = $1 AND user1_id = $2)`)
	err := RR.db.QueryRow(SqlStatemnt, UserId1, UserId2).Scan(&RoomId)
	if err != nil {
		return 0, err
	}
	return RoomId, nil
}
func (RR *RoomRepo) FindUsersByRoomId(RoomId int) (int, int, error) {
	var UserId1 int
	var UserId2 int
	SqlStatemnt := (`SELECT user1_id, user2_id FROM private_chats WHERE id=$1`)
	err := RR.db.QueryRow(SqlStatemnt, RoomId).Scan(&UserId1, &UserId2)
	if err != nil {
		return 0, 0, err
	}
	return UserId1, UserId2, nil
}
func (RR *RoomRepo) CreateRoom(UserId1, UserId2 int, ca time.Time) (int, error) {
	var Id int
	SqlStatement := (`INSERT INTO private_chats
	 (user1_id,user2_id,created_at) 
	 VALUES ($1,$2,$3)
	 ON CONFLICT (user1_id,user2_id) DO NOTHING 
	  RETURNING id`)
	SqlStatementifhave := (`SELECT id FROM private_chats WHERE (user1_id=$1 AND user2_id = $2) OR (user2_id = $1 AND user1_id = $2)`)
	err := RR.db.QueryRow(SqlStatement, UserId1, UserId2, ca).Scan(&Id)
	if err != nil {
		err := RR.db.QueryRow(SqlStatementifhave, UserId1, UserId2).Scan(&Id)
		if err != nil {
			return 0, err
		}
		return Id, err
	}
	return Id, nil
}
