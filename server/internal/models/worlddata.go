package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type WorldData struct{
	Id int `db:"id"`
	Name string `db:"name"`
	OwnerUserId int `db:"owner_user_id"`
	ModifierData string `db:"modifier_data"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted int `db:"deleted"`
}

func CreateWorld(db *sqlx.DB, uid int, worldName string) (*WorldData, error) {
	// start transaction
	tx, err := db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}; defer tx.Rollback()

	w := &WorldData{}

	err = db.Get(w, `
		INSERT INTO worlddata (owner_user_id, name)
		VALUES (?, ?)
		RETURNING *
	`, uid, worldName)
	if err != nil {
		return nil, fmt.Errorf("failed to create world: %v\n", err)
	}

	_, err = db.Exec(`
		INSERT INTO users_worlddata(uid, wid)
		VALUES (?, ?)
		RETURNING *
	`, w.Id, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to create resolver entry: %v\n", err)
	}

	// commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return w, nil 
}

func GetWorldData(db *sqlx.DB, id int) (*WorldData, error){
	w := &WorldData{}

	err:= db.Unsafe().Get(w, `
		SELECT * FROM worlddata 
		WHERE id = ?
		AND deleted == 0
		LIMIT 1;
	`, id)

	if err != nil {
		return nil, fmt.Errorf("failed to get world: %v", err)
	}

	return w, nil
}
