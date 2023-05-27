package main

import (
	"fmt"
	"net/http"
)

type CounterHandler struct {
	data string
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ct.data = "{\"source\":\"server2\",\"data\":[{\"ID\":1,\"name\":\"ali\",\"age\":18},{\"ID\":2,\"name\":\"ali1\",\"age\":18}],\"message\":\"SUCCESS\"}"
	fmt.Println(ct.data)
	fmt.Fprintln(w, ct.data)
}

func main() {
	fmt.Println("Listening on Port 8012")

	th := &CounterHandler{}
	http.Handle("/posts", th)
	http.ListenAndServe(":8012", nil)
}

//https://zetcode.com/golang/http-server/
