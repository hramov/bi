package dashboard_handler

import (
	"github.com/go-chi/chi/v5"
	dashboards_dto_out "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/out"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
)

func (h *Handler) getByDashId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendError(http.StatusInternalServerError, "no id found", w)
		return
	}

	data, err := h.service.GetByDashId(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	var items []*dashboards_dto_out.Item

	for _, v := range data.Items {
		items = append(items, &dashboards_dto_out.Item{
			Id:          v.Id,
			DashId:      v.DashId,
			ItemType:    v.ItemType,
			Position:    v.Position,
			Title:       v.Title,
			Description: v.Description,
			DataQueries: v.DataQueries,
			Options:     v.Options,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			DeletedAt:   v.DeletedAt,
		})
	}

	result := &dashboards_dto_out.Dashboard{
		Id:          data.Id,
		DashId:      data.DashId,
		Title:       data.Title,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		DeletedAt:   data.DeletedAt,
		Items:       items,
	}

	utils.SendResponse(http.StatusOK, result, w)
}
