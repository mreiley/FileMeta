package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Project json format

type jerr struct {
	Msg string `json:"msj"`
}

// {"name":"Mario20Dollares.pdf","type":"application/pdf","size":7156}:q
type meta struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}

// just begin the services
func HandlerSendFile(w http.ResponseWriter, req *http.Request) {
	fileBytes, err := os.ReadFile("./views/index.html")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)

	return
}

func fileanalyse(w http.ResponseWriter, req *http.Request) {

	_, fh, e := req.FormFile("upfile")

	if e != nil {
		log.Println(e)
	} else {

		res, _ := json.Marshal(meta{Name: fh.Filename, Type: fh.Header.Get("Content-Type"), Size: fh.Size})
		w.Write(res)

	}

}

func main() {
	http.HandleFunc("/", HandlerSendFile)
	http.HandleFunc("/api/fileanalyse", fileanalyse)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}
