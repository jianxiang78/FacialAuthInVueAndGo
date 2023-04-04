package main

import (
	"./routes"
	"fmt"
)


func main() {
	fmt.Println("APP server start running...")
	routes.ApiRouter()
}

