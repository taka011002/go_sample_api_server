package main

import (
	"fmt"
	"github.com/taka011002/go_sample_api_server/app"
	"os"
)

func main() {
	// サーバ起動
	fmt.Println("========================")
	fmt.Println(fmt.Sprintf("Server Start >> http://localhost:%s", os.Getenv("PORT")))
	fmt.Println("========================")
	app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}