package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var flagvar int
	flag.IntVar(&flagvar, "port", 80, "mention port numbe here")
	flag.Parse()
	fmt.Println("port configured is", flagvar)

	fmt.Println("starting hello world")
	http.HandleFunc("/", hello)
	http.HandleFunc("/demox", demo)
	http.ListenAndServe(":"+strconv.Itoa(flagvar), nil)
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

func demo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("received request from", req.RemoteAddr)
	f, _ := os.Open("/tmp/soumya/podi_undakkanni.txt")
	f.WriteString("buhaaa")

	w.Write([]byte("success"))

}
