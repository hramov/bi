package datasource_handler

import (
	"context"
	data_source_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/data_source/dto/in"
	"github.com/hramov/gvc-bi/backend/dashboard/internal/errors"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"time"
)

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()

	source, err := utils.GetBody[data_source_dto_in.Datasource](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, app_errors.ErrBadRequest, w)
		return
	}

	id, err := h.service.Create(ctx, source)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, app_errors.ErrInternal, w)
		return
	}

	utils.SendResponse(http.StatusCreated, id, w)
}
