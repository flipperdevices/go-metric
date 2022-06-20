package main

import (
	"net/http"
)

func report(w http.ResponseWriter, req *http.Request) {

}

func main() {
	http.HandleFunc("/report", report)
}
