package main

import (
	gd "chapter1-single-server/util"

	"fmt"
	"net/http"
)

func main() {
	fmt.Println("1Listening on Port 8001")
	gd.Serve()
	http.ListenAndServe(":8001", nil)
}

//https://github.com/bmf-san/godon
