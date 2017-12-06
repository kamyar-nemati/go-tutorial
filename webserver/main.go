package main

import (
	"fmt"
	"net"
	"net/http"
)

func checkError(err *error) {
	if *err != nil {
		fmt.Println("Error:", (*err).Error())
	}
}

func httpHendler(w http.ResponseWriter, r *http.Request) {
	//log the request
	fmt.Printf("Requested from: %s | requested URI: %s\n",
		r.RemoteAddr, r.RequestURI)

	//prepare some response
	resp := []byte("Hello, World!")

	//send back the response
	_, err := w.Write(resp)

	checkError(&err)
}

func main() {
	fmt.Println("Server is running")

	//listen to some port on localhost
	ln, err := net.Listen("tcp", "127.0.0.1:9050")

	checkError(&err)

	//assign the client handler function
	http.HandleFunc("/", httpHendler)

	//start serving clients
	err = http.Serve(ln, nil)

	checkError(&err)
}
