package handler_ds

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/internal/repository"
	"github.com/hramov/gvc-bi/backend/pkg/database"
	"github.com/hramov/gvc-bi/backend/pkg/utils"
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
	Repository repository.DatasourceRepository
}

func New(repository repository.DatasourceRepository) *Handler {
	return &Handler{
		Repository: repository,
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

	source, err := h.Repository.GetByCode(query.Source)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot find data source: %v", err.Error()), w)
		return
	}

	storage, err := database.NewStorageForQuery(source.Driver, source.Dsn)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot check connection: %v", err.Error()), w)
		return
	}
	defer storage.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := storage.QueryContext(ctx, query.Query, query.Params...)

	if err != nil {
		utils.SendResponse(http.StatusInternalServerError, fmt.Sprintf("cannot perform query: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusOK, utils.Jsonify(rows), w)
}
