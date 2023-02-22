package main

import (
	"MyBlog/utils"
	"log"
)

func main() {
	_, err := utils.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}
}
