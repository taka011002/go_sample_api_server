package main

import (
	"fmt"
	"github.com/taka011002/go_sample_api_server/app"
	"github.com/taka011002/go_sample_api_server/app/infra"
)

func main() {
	// マイグレーション
	infra.Up()

	// サーバ起動
	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:8080")
	fmt.Println("========================")
	app.Run(":8080")
}