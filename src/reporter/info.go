package reporter

import (
	"fmt"
	"net/http"
)

func (r *Reporter) GetInfo(w http.ResponseWriter, req *http.Request) {
	isDbWorks, err := r.repo.CheckDb(req.Context())
	var response = "Null"
	if err != nil {
		response = fmt.Sprintf("Failed connect to DB:  %s", err)
	} else if isDbWorks {
		response = "Everything is fine"
	} else {
		response = "Something went wrong"
	}
	_, err = fmt.Fprintf(w, response)
	if err != nil {
		panic(err)
	}
}
