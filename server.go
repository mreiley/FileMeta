/*
AUTOR: Mario Reiley backend Developer
NOTE: it is my port to Golang 1.19

* You can submit a form that includes a file upload.
* The form file input field has the name attribute set to upfile.
* When you submit a file, you receive the file name, type, and size in bytes within the JSON response.
*/
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type jerr struct {
	Msg string `json:"msj"`
}

// Create mehtos for this struct
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

// I did test with postman and Thunder client vscode extension and work fime
func fileanalyse(w http.ResponseWriter, req *http.Request) {

	_, fh, e := req.FormFile("upfile")

	if e != nil {
		log.Println(e) // Just log the error and continue
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
