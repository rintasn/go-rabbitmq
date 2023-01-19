package main

import (
	"fmt"
	"log"
	"main/router"
	"net/http"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Server dijalankan pada port 81...")

	log.Fatal(http.ListenAndServe(":81", r))
}
