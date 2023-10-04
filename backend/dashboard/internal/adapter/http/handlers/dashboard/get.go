package dashboard_handler

import (
	"context"
	dashboards_dto_out "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/out"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	data, err := h.service.Get(ctx)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	var response []*dashboards_dto_out.Dashboard

	for _, d := range data {
		response = append(response, &dashboards_dto_out.Dashboard{
			Id:          d.Id,
			DashId:      d.DashId,
			Title:       d.Title,
			Description: d.Description,
			CreatedAt:   d.CreatedAt,
			UpdatedAt:   d.UpdatedAt,
			DeletedAt:   d.DeletedAt,
		})
	}

	utils.SendResponse(http.StatusOK, response, w)
}
