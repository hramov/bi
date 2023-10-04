package handler

import (
	"context"
	"fmt"
	"github.com/hramov/gvc-bi/backend/datastorage/internal/domain/data_source/connections"
	"github.com/hramov/gvc-bi/backend/datastorage/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) performQuery(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	query, err := utils.GetBody[QueryOptions](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot parse body: %v", err.Error()), w)
		return
	}

	source, err := h.repo.GetByCode(ctx, query.Source)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot find data source: %v", err.Error()), w)
		return
	}

	storage, err := connections.Get(source.Id)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("cannot check connection: %v", err.Error()), w)
		return
	}

	rows, err := storage.QueryContext(ctx, query.Query, query.Params...)

	if err != nil {
		utils.SendResponse(http.StatusInternalServerError, fmt.Sprintf("cannot perform query: %v", err.Error()), w)
		return
	}

	h.logger.Info(query.Query)

	utils.SendResponse(http.StatusOK, utils.Jsonify(rows), w)
}
