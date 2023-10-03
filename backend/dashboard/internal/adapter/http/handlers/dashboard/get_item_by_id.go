package dashboard_handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"strconv"
)

func (h *Handler) getItemById(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.service.GetItemById(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, err.Error(), w)
		return
	}

	utils.SendResponse(http.StatusOK, data, w)
}
