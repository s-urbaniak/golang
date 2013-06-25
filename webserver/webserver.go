package main

import (
	"fmt"
	"net/http"
)

func main() {

	var (
		dir  string
		port string
	)

	fmt.Printf("Please input which directory\nwhat you want to share, start with \"/\":\n")
	fmt.Scanf("%s", &dir)
	h := http.FileServer(http.Dir(dir))
	fmt.Printf("Please input port Number: \n")
	fmt.Scanf("%s", &port)
	http.ListenAndServe(":"+port, h)
}
