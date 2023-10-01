package handler_ds

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/dashboard/internal"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/connections"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/repository"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/database"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"time"
)

type QueryOptions struct {
	Source string `json:"source"`
	Query  string `json:"query"`
	Params []any  `json:"params"`
}

type CheckResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Handler struct {
	repo   repository.DatasourceRepository
	logger internal.Logger
}

func New(repo repository.DatasourceRepository, logger internal.Logger) *Handler {
	return &Handler{
		repo:   repo,
		logger: logger,
	}
}

func (h *Handler) Register(r chi.Router) {
	r.Post("/check", h.checkConnection)
	r.Post("/perform", h.performQuery)
}

func (h *Handler) checkConnection(w http.ResponseWriter, r *http.Request) {

	dataSource, err := utils.GetBody[database.DataStorageOptions](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot parse body: %v", err.Error()), w)
		return
	}

	storage, err := database.NewStorage(dataSource)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot check connection: %v", err.Error()), w)
		return
	}
	defer storage.Close()

	result := CheckResult{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = storage.PingContext(ctx)

	if err != nil {
		result.Status = "declined"
		result.Message = err.Error()
		utils.SendResponse(http.StatusOK, result, w)
		return
	}

	result.Status = "accepted"
	utils.SendResponse(http.StatusOK, result, w)
}

func (h *Handler) performQuery(w http.ResponseWriter, r *http.Request) {
	query, err := utils.GetBody[QueryOptions](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot parse body: %v", err.Error()), w)
		return
	}

	source, err := h.repo.GetByCode(query.Source)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot find data source: %v", err.Error()), w)
		return
	}

	storage, err := connections.Get(source.Id)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot check connection: %v", err.Error()), w)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := storage.QueryContext(ctx, query.Query, query.Params...)

	if err != nil {
		utils.SendResponse(http.StatusInternalServerError, fmt.Sprintf("cannot perform query: %v", err.Error()), w)
		return
	}

	h.logger.Info(query.Query)

	utils.SendResponse(http.StatusOK, utils.Jsonify(rows), w)
}

func (h *Handler) recallDataSources(w http.ResponseWriter, r *http.Request) {
	ds, err := h.repo.Get()
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	var rc []connections.RawConnection

	for _, v := range ds {
		rc = append(rc, connections.RawConnection{
			SourceId: v.Id,
			DriverId: v.DriverId,
			Dsn:      v.Dsn,
		})
	}

	errs := connections.Connect(rc)

	errStr := ""

	for _, e := range errs {
		errStr += e.Error() + ":"
	}

	if len(errs) > 0 {
		utils.SendError(http.StatusInternalServerError, errStr, w)
		return
	}

	utils.SendResponse(http.StatusOK, "OK", w)
}
