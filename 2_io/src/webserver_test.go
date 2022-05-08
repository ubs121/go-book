package io

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Сайн уу, "+req.URL.Path[1:])
}

func TestWebserver(t *testing.T) {

	// start the web server
	go func() {
		http.HandleFunc("/", HelloServer)
		err := http.ListenAndServe("localhost:8080", nil)

		if err != nil {
			log.Fatal("ListenAndServe: ", err.Error())
		}
	}()

	resp, err := http.Get("http://localhost:8080/Uuganaa")
	if err != nil {
		t.Error(err)
	}
	respBuf, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response: %s\n", string(respBuf))
}

func TestServeDir(t *testing.T) {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("/web")))
}
