package dashboard_handler

import (
	"context"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/errors"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) getAvailableTypes(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	data, err := h.service.GetAvailableTypes(ctx)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, app_errors.ErrInternal, w)
		return
	}
	utils.SendResponse(http.StatusOK, data, w)
}
