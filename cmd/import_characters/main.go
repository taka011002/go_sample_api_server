package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/taka011002/go_sample_api_server/app/domain/entity"
	"github.com/taka011002/go_sample_api_server/app/domain/service"
	"github.com/taka011002/go_sample_api_server/app/infra"
	"github.com/taka011002/go_sample_api_server/app/infra/persistence"
	"log"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	file, err := os.Open(flag.Arg(0))
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	_, err = reader.Read() // ヘッダー処理
	if err != nil {
		log.Fatal(err)
	}

	p := persistence.NewCharacterPersistence(infra.DB)
	s := service.NewCharacterService(p)

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		characterRarityId, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		characterPower, err := strconv.Atoi(line[2])
		if err != nil {
			log.Fatal(err)
		}

		c := entity.Character{Name: line[0], CharacterRarityId: characterRarityId, Power: characterPower}
		if err := s.CreateOrUpdate(&c); err != nil {
			fmt.Println(err)
			fmt.Println("failed insert", c.Name)
		} else {
			fmt.Println("successed insert", c.Name)
		}

	}

}
