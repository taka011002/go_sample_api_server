package main

import "github.com/taka011002/go_sample_api_server/app/infra"

func main() {
	infra.Init()
	defer infra.Close()
	infra.Up()
}