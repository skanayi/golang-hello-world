package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("starting hello world")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("received request from", req.RemoteAddr)
	hs, err := os.Hostname()

	if err != nil {

		fmt.Println("error occured while getting hostname")
		hs = ""

	}

	hs = "Hello world ,hostname is: " + hs

	w.Write([]byte(hs))

}
