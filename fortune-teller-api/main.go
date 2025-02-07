package main

import (
	"fmt"
	"fortune-teller-api/config"
)

func main() {
	config.ConnectDB()
	fmt.Println("🎉 Fortune Teller API is running!")
}
