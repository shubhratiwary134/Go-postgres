package main

import (
	"fmt"
	"fortune-teller-api/config"
	"fortune-teller-api/routes"
)

func main() {

	config.ConnectDB()
	fmt.Println("🎉 Fortune Teller API is running!")
	router := routes.SetupRouter()
	router.Run(":8080")
}
