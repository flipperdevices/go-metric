package reporter

import (
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/uptrace/go-clickhouse/ch/chschema"
	"google.golang.org/protobuf/proto"

	pb "github.com/flipperdevices/go-metric/internal/proto"
	"github.com/flipperdevices/go-metric/src/models"
	"github.com/flipperdevices/go-metric/src/repository"
)

func New(r *repository.Repository) *Reporter {
	return &Reporter{repo: r}
}

type Reporter struct {
	repo *repository.Repository
}

func (r *Reporter) Report(w http.ResponseWriter, req *http.Request) {
	rb, err := io.ReadAll(req.Body)
	if err != nil {
		log.Print("Failed read thread request body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var request pb.MetricReportRequest
	if err = proto.Unmarshal(rb, &request); err != nil {
		log.Print("Failed parse request body", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	id, err := uuid.Parse(request.Uuid)
	if err != nil {
		log.Print("Failed parse uuid", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var platform models.Platform
	switch request.Platform {
	case pb.MetricReportRequest_ANDROID:
		platform = models.ANDROID
	case pb.MetricReportRequest_IOS:
		platform = models.IOS
	}

	for _, event := range request.Events {
		err = r.repo.SaveEvent(req.Context(), chschema.UUID(id), platform, event)
		if err != nil {
			log.Println("Failed save event", err)
			w.WriteHeader(http.StatusInsufficientStorage)
			return
		}
	}

	log.Print(request.String())
	sendResponse(w)
}

func sendResponse(w http.ResponseWriter) {
	response, err := proto.Marshal(&pb.MetricReportResponse{
		Status: pb.MetricReportResponse_OK,
	})
	if err != nil {
		log.Print("Failed marshal response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		// Unable to set status after attempt to write body, just log.
		log.Print("Failed write response", err)
		return
	}
}
