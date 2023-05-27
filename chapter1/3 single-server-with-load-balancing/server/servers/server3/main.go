package main

import (
	"fmt"
	"net/http"
)

type CounterHandler struct {
	data string
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ct.data = "{\"source\":\"server3\",\"data\":[{\"ID\":1,\"name\":\"ali\",\"age\":18},{\"ID\":2,\"name\":\"ali1\",\"age\":18}],\"message\":\"SUCCESS\"}"
	fmt.Println(ct.data)
	fmt.Fprintln(w, ct.data)
}

func main() {
	fmt.Println("Listening on Port 8013")

	th := &CounterHandler{}
	http.Handle("/posts", th)
	http.ListenAndServe(":8013", nil)
}

//https://zetcode.com/golang/http-server/
