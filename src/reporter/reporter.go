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

	var clickhouseSid *chschema.UUID = nil
	if len(request.SessionUuid) > 0 {
		id, err := uuid.Parse(request.SessionUuid)
		if err == nil {
			clickhouseSidOriginal := chschema.UUID(id)
			clickhouseSid = &clickhouseSidOriginal
		}
	}

	var appVersion *string = nil
	if len(request.Version) > 0 {
		appVersion = &request.Version
	}

	var platform models.Platform
	switch request.Platform {
	case pb.MetricReportRequest_ANDROID:
		platform = models.ANDROID
	case pb.MetricReportRequest_ANDROID_DEBUG:
		platform = models.ANDROID_DEBUG
	case pb.MetricReportRequest_IOS:
		platform = models.IOS
	case pb.MetricReportRequest_IOS_DEBUG:
		platform = models.IOS_DEBUG
	}

	for _, event := range request.Events {
		err = r.repo.SaveEvent(req.Context(),
			chschema.UUID(id),
			clickhouseSid,
			appVersion,
			platform,
			event)
		if err != nil {
			log.Println("Failed save event", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
