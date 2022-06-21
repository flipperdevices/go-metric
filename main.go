package main

import (
	"fmt"
	pb "github.com/flipperdevices/go-metric/internal/proto"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net/http"
)

func report(w http.ResponseWriter, req *http.Request) {
	rb, err := io.ReadAll(req.Body)
	if err != nil {
		log.Print("Failed read thread request body", err)
		return
	}

	var r pb.MetricReportRequest
	err = proto.Unmarshal(rb, &r)
	if err != nil {
		log.Print("Failed parse request body", err)
		return
	}
	fmt.Println(r.String())
	response, err := proto.Marshal(&pb.MetricReportResponse{
		Status: pb.MetricReportResponse_OK,
	})
	if err != nil {
		log.Print("Failed marshal response", err)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		log.Print("Failed write response", err)
		return
	}
}

func main() {
	http.HandleFunc("/report", report)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
