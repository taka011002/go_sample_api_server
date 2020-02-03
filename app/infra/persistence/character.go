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

func (cp characterPersistence) Create(name string, characterRarityId int, power int) error {
	stmt, err := cp.DB.Prepare("INSERT INTO characters(id, name, character_rarity_id, power) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(0, name, characterRarityId, power)
	return err
}

func (cp characterPersistence) Update(id int, name string, characterRarityId int, power int) error {
	stmt, err := cp.DB.Prepare("UPDATE characters SET name = ?, character_rarity_id = ?, power = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, characterRarityId, power, id)
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
	row := cp.DB.QueryRow("SELECT * FROM characters WHERE name = ?", name)
	//row型をgolangで利用できる形にキャストする。
	return convertToCharacter(row)
}

func (cp characterPersistence) GetRand(characterRarityId int) (*entity.Character, error) {
	q := "SELECT tbl.* FROM characters AS tbl,( SELECT id FROM characters WHERE character_rarity_id = ? ORDER BY RAND() LIMIT 1) AS randam WHERE tbl.id = randam.id LIMIT 1"
	row := cp.DB.QueryRow(q, characterRarityId)
	return convertToCharacter(row)
}

func convertToCharacter(row *sql.Row) (*entity.Character, error) {
	c := entity.Character{}
	err := row.Scan(&c.Id, &c.Name, &c.CharacterRarityId, &c.Power)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
