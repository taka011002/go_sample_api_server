package persistence

import (
	"database/sql"
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/repository"
)

type characterRarityPersistence struct {
	DB *sql.DB
}

func NewCharacterRarityPersistence(DB *sql.DB) repository.CharacterRarityRepository {
	return &characterRarityPersistence{DB: DB}
}

func (cp characterRarityPersistence) Create(name string, rarity int) error {
	stmt, err := cp.DB.Prepare("INSERT INTO character_rarities(id, name, rarity) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(0, name, rarity)
	return err
}

func (cp characterRarityPersistence) Update(id int, name string, rarity int) error {
	stmt, err := cp.DB.Prepare("UPDATE character_rarities SET name = ?, rarity = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, rarity,id)
	return err
}

func (cp characterRarityPersistence) Delete(id int) error {
	stmt, err := cp.DB.Prepare("DELETE FROM character_rarities WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}

func (cp characterRarityPersistence) GetByName(name string) (*entity.CharacterRarity, error) {
	row := cp.DB.QueryRow("SELECT * FROM character_rarities WHERE name = ?", name)
	//row型をgolangで利用できる形にキャストする。
	return convertToCharacterRarity(row)
}

func (cp characterRarityPersistence) GetAll() (*entity.CharacterRarities, error) {
	rows, err := cp.DB.Query("SELECT * FROM character_rarities")
	if err != nil {
		return nil, err
	}

	return convertToCharacterRarities(rows)
}

func convertToCharacterRarity(row *sql.Row) (*entity.CharacterRarity, error) {
	c := entity.CharacterRarity{}
	err := row.Scan(&c.Id, &c.Name, &c.Rarity)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func convertToCharacterRarities(rows *sql.Rows) (*entity.CharacterRarities, error) {
	var res entity.CharacterRarities
	for rows.Next() {
		c := entity.CharacterRarity{}
		err := rows.Scan(&c.Id, &c.Name, &c.Rarity)

		if err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return &res, nil
}