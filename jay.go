package main

import (
	"fmt"
	"io"
	"net/http" //the http package also provides the APIs to create HTTPS secure servers
)

type httphandler struct{}

func (h httphandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello by")
	fmt.Fprintf(res, "Jay ")
	data := []byte("❤️")
	res.Write(data)
}

func main() {
	handler := httphandler{}
	//fmt.Println("HANDLER = ", handler)
	http.ListenAndServe(":9000", handler)
}
