package main

import "github.com/taka011002/go_sample_api_server/app/infra"

func main() {
	db := infra.DB
	infra.Up(db)
	defer db.Close()
}
