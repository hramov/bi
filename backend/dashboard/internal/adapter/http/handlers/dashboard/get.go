package dashboard_handler

import (
	dashboards_dto_out "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/out"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
)

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.Get()
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
