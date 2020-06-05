package main

import (
	"fmt"
	"net/http"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func ServerHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola golang desde Kubernetes")
}

func main() {
	http.HandleFunc("/", ServerHttp)
	http.ListenAndServe(":9090", nil)
}
