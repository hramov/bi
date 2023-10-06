package dashboard_handler

import (
	"context"
	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"
	app_errors "github.com/hramov/gvc-bi/backend/dashboard/internal/errors"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	dash, err := utils.GetBody[dashboards_dto_in.Dashboard](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, app_errors.ErrBadRequest, w)
		return
	}

	id, err := h.service.Create(ctx, dash)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, app_errors.ErrInternal, w)
		return
	}

	utils.SendResponse(http.StatusCreated, id, w)
}
