package main

import (
	"fmt"
	"log"
	"task_manager/data"
	"task_manager/router"
)

func init() {
	fmt.Println("init runs")
	err := data.Dbconnect()
	if err != nil {
		// panic(err)
		fmt.Println("connection error")
		log.Fatal(err)
		return
	}
}

func main() {
	router.Router()
}
