package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type token struct {
	Token string `json:"token"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	base := "http://localhost"
	base += fmt.Sprintf(":%s/", os.Getenv("PORT"))

	for i := 0; i < 20; i++ {
		n := strconv.Itoa(i+1)
		jsonStr := `{"username": "`+ randString(10)  + `", "password": "password"}`

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

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}