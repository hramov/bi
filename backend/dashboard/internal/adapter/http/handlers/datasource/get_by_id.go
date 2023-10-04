package datasource_handler

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	data_source_dto_out "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/out"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) getById(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	rawId := chi.URLParam(r, "id")
	if rawId == "" {
		utils.SendError(http.StatusBadRequest, "need to pass id parameter", w)
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong id format: %v", err.Error()), w)
		return
	}

	source, err := h.service.GetById(ctx, id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}
	utils.SendResponse(http.StatusOK, &data_source_dto_out.Datasource{
		Id:          source.Id,
		Driver:      source.Driver,
		DriverId:    source.DriverId,
		Title:       source.Title,
		Dsn:         source.Dsn,
		Checked:     source.Checked,
		DateCreated: source.DateCreated,
	}, w)
}
