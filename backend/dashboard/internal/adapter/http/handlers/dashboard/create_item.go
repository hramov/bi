package dashboard_handler

import (
	"fmt"
	dashboards_dto_in "github.com/hramov/gvc-bi/backend/dashboard/internal/domain/dashboard/dto/in"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	item, err := utils.GetBody[dashboards_dto_in.Item](r)
	if err != nil {
		utils.SendError(http.StatusBadRequest, fmt.Sprintf("wrong body format: %v", err.Error()), w)
		return
	}

	id, err := h.service.CreateItem(item)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot save data to database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, id, w)
}
