package datasource_handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hramov/gvc-bi/backend/dashboard/pkg/utils"
	"net/http"
	"strconv"
)

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
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

	deletedId, err := h.service.Delete(id)
	if err != nil {
		utils.SendError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch data from database: %v", err.Error()), w)
		return
	}

	utils.SendResponse(http.StatusCreated, deletedId, w)
}
