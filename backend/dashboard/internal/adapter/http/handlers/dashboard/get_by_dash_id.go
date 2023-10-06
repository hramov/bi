package dashboard_handler

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	dashboards_dto_out "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/out"
	app_errors "github.com/hramov/gvc-bi/backend/dashboard/internal/errors"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) getByDashId(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	if id == "" {
		utils.SendError(http.StatusInternalServerError, app_errors.ErrNoId, w)
		return
	}

	data, err := h.service.GetByDashId(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.SendCustomError(ctx, http.StatusNotFound, app_errors.New(err, app_errors.ErrNotFound, nil), w)
			return
		}
		utils.SendError(http.StatusInternalServerError, app_errors.ErrInternal, w)
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
