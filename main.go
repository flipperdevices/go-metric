package main

import (
	"context"
	"fmt"
	pb "github.com/flipperdevices/go-metric/internal/proto"
	"github.com/google/uuid"
	"github.com/uptrace/go-clickhouse/ch/chschema"
	"github.com/uptrace/go-clickhouse/chdebug"
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

	id, err := uuid.Parse(r.Uuid)
	if err != nil {
		log.Print("Failed parse uuid", err)
		return
	}
	var platform Platform
	switch r.Platform {
	case pb.MetricReportRequest_ANDROID:
		platform = android
	case pb.MetricReportRequest_IOS:
		platform = ios
	}

	for _, event := range r.Events {
		err = saveEvent(req.Context(), chschema.UUID(id), platform, event)
		if err != nil {
			log.Println("Failed save event", err)
			return
		}
	}

	fmt.Println(r.String())
	sendResponse(w)
}

func sendResponse(w http.ResponseWriter) {
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
	ctx := context.Background()
	db.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))
	err := db.ResetModel(ctx, (*AppOpen)(nil))
	if err != nil {
		panic(err)
	}

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	http.HandleFunc("/report", report)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
