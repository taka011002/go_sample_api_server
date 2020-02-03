package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type token struct {
	Token string `json:"token"`
}

func main() {
	base := "http://localhost"
	base += fmt.Sprintf(":%s/", os.Getenv("PORT"))

	for i := 0; i < 20; i++ {
		n := strconv.Itoa(i+1)
		jsonStr := `{"username": "taka` + n + `", "first_name": "string","last_name": "string","email": "taka` + n + `", "password": "password","phone": "string","user_status": 1}`

		req, err := http.NewRequest(http.MethodPost, base + "user", bytes.NewBuffer([]byte(jsonStr)))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		r, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		token := token{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&token); err != nil {
			log.Fatal(err)
		}
		_ = r.Body.Close()

		jsonStr = `{"times":`+ n + `}`
		req, err = http.NewRequest(http.MethodPost, base + "gacha/draw", bytes.NewBuffer([]byte(jsonStr)))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.Token))
		client = &http.Client{}
		r, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

	}
}