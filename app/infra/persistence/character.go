package persistence

import (
	"database/sql"
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type characterPersistence struct {
	DB *sql.DB
}

func NewCharacterPersistence(DB *sql.DB) repository.CharacterRepository {
	return &characterPersistence{DB: DB}
}

func (cp characterPersistence) Create(name string) error {
	stmt, err := cp.DB.Prepare("INSERT INTO characters(id, name) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(0, name)
	return err
}

func (cp characterPersistence) Update(id int, name string) error {
	stmt, err := cp.DB.Prepare("UPDATE characters SET name = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, id)
	return err
}

func (cp characterPersistence) Delete(id int) error {
	stmt, err := cp.DB.Prepare("DELETE FROM characters WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}

func (cp characterPersistence) GetByName(name string) (*entity.Character, error) {
	row := cp.DB.QueryRow("SELECT * FROM characters WHERE username = ?", name)
	//row型をgolangで利用できる形にキャストする。
	return convertToCharacter(row)
}

func convertToCharacter(row *sql.Row) (*entity.Character, error) {
	c := entity.Character{}
	err := row.Scan(&c.Id, &c.Name)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
