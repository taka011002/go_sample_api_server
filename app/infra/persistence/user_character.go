package persistence

import (
	"database/sql"
	"github.com/taka011002/go_sample_api_server/app/domain/entity"

	"github.com/taka011002/go_sample_api_server/app/domain/repository"
	"strconv"
	"strings"
)

type userCharacterPersistence struct {
	DB *sql.DB
}

func NewUserCharacterPersistence(DB *sql.DB) repository.UserCharacterRepository {
	return &userCharacterPersistence{DB: DB}
}

func (up userCharacterPersistence) Create(userId int, characterId int) error {
	stmt, err := up.DB.Prepare("INSERT INTO user_character(id, user_id, character_id) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(0, userId, characterId)
	return err
}

//BULK INSERT
func (up userCharacterPersistence) Creates(userId int, characterId []int) error {
	q := "INSERT INTO user_character(id, user_id, character_id) VALUES"
	add := "(?, ?, ?),"
	var exec []interface{}
	for c := range characterId {
		q += add
		exec = append(exec, strconv.Itoa(0), strconv.Itoa(userId), strconv.Itoa(characterId[c]))
	}

	q = strings.TrimRight(q, ",") + ";"

	stmt, err := up.DB.Prepare(q)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(exec...)

	return err
}

func (up userCharacterPersistence) WhereByUserId(userId int) (*entity.UserCharacters, error) {
	q := "SELECT * FROM user_character AS uc INNER JOIN characters AS ch ON ch.id = uc.character_id WHERE uc.`user_id` = ?;"

	rows, err := up.DB.Query(q, userId)
	if err != nil {
		return nil, err
	}

	return convertToUserCharacters(rows)
}

func convertToUserCharacters(rows *sql.Rows) (*entity.UserCharacters, error) {
	var res entity.UserCharacters
	for rows.Next() {
		uc := entity.UserCharacter{}
		c := entity.Character{}
		err := rows.Scan(&uc.Id, &uc.UserId, &uc.CharacterId, &c.Id, &c.Name, &c.CharacterRarityId)

		if err != nil {
			return nil, err
		}

		uc.Character = c

		res = append(res, uc)
	}

	return &res, nil
}