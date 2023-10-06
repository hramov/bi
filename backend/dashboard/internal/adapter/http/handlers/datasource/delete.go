package datasource_handler

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/errors"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	rawId := chi.URLParam(r, "id")
	if rawId == "" {
		utils.SendError(http.StatusBadRequest, app_errors.ErrNoId, w)
		return
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		utils.SendError(http.StatusBadRequest, app_errors.ErrWrongIdFormat, w)
		return
	}

	deletedId, err := h.service.Delete(ctx, id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, app_errors.ErrInternal, w)
		return
	}

	utils.SendResponse(http.StatusCreated, deletedId, w)
}
